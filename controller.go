package main

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/apis/extensions"
	"k8s.io/kubernetes/pkg/client/cache"
	k8s_clientset "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
	k8s_client "k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/controller/framework"
	"k8s.io/kubernetes/pkg/runtime"
	deploymentutil "k8s.io/kubernetes/pkg/util/deployment"
	"k8s.io/kubernetes/pkg/watch"
)

const (
	// Resync period for the kube controller loop.
	resyncPeriod = 5 * time.Minute
)

var (
	keyFunc = cache.MetaNamespaceKeyFunc
)

// Controller listwatch events, pod, deployments and  handle them
type Controller struct {
	client        *k8s_client.Client
	clientset     *k8s_clientset.Clientset
	stopCh        chan struct{}
	stopLock      sync.Mutex
	shutdown      bool
	podController *framework.Controller
	podStore      cache.Store
	dpController  *framework.Controller
	dpStore       cache.Store
	dpMapper      map[string]*api.PodList
}

// NewController returns Controller instance or an error
func NewController(client *k8s_client.Client) (*Controller, error) {
	c := &Controller{
		client:   client,
		stopCh:   make(chan struct{}),
		dpMapper: map[string]*api.PodList{},
	}

	c.podStore, c.podController = framework.NewInformer(
		&cache.ListWatch{
			ListFunc:  podListFunc(c.client, api.NamespaceAll),
			WatchFunc: podWatchFunc(c.client, api.NamespaceAll),
		},
		&api.Pod{}, resyncPeriod, framework.ResourceEventHandlerFuncs{})

	c.dpStore, c.dpController = framework.NewInformer(
		&cache.ListWatch{
			ListFunc:  dpListFunc(c.client, api.NamespaceAll),
			WatchFunc: dpWatchFunc(c.client, api.NamespaceAll),
		},
		&extensions.Deployment{}, resyncPeriod, framework.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				c.updateDp(obj)
			},
			UpdateFunc: func(old, cur interface{}) {
				c.updateDp(cur)
			},
			DeleteFunc: func(obj interface{}) {
				c.deleteDp(obj)
			},
		})

	return c, nil
}

func (c *Controller) updateDp(obj interface{}) {
	dp := obj.(*extensions.Deployment)
	key, err := keyFunc(dp)
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	pods, err := deploymentutil.ListPods(dp, func(ns string, options api.ListOptions) (*api.PodList, error) {
		return c.client.Pods(ns).List(options)
	})
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	c.dpMapper[key] = pods
}

func (c *Controller) deleteDp(obj interface{}) {
	dp := obj.(*extensions.Deployment)
	key, err := keyFunc(dp)
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	delete(c.dpMapper, key)
}

func podListFunc(c *k8s_client.Client, ns string) func(api.ListOptions) (runtime.Object, error) {
	return func(options api.ListOptions) (runtime.Object, error) {
		return c.Pods(ns).List(options)
	}
}

func podWatchFunc(c *k8s_client.Client, ns string) func(api.ListOptions) (watch.Interface, error) {
	return func(options api.ListOptions) (watch.Interface, error) {
		return c.Pods(ns).Watch(options)
	}
}

func dpListFunc(c *k8s_client.Client, ns string) func(api.ListOptions) (runtime.Object, error) {
	return func(options api.ListOptions) (runtime.Object, error) {
		return c.Deployments(ns).List(options)
	}
}

func dpWatchFunc(c *k8s_client.Client, ns string) func(api.ListOptions) (watch.Interface, error) {
	return func(options api.ListOptions) (watch.Interface, error) {
		return c.Deployments(ns).Watch(options)
	}
}

// Run event store
func (c *Controller) Run() {
	log.Infoln("start event store...")
	go c.podController.Run(c.stopCh)
	go c.dpController.Run(c.stopCh)
	<-c.stopCh
}

// Stop event store
func (c *Controller) Stop() error {
	c.stopLock.Lock()
	defer c.stopLock.Unlock()

	if !c.shutdown {
		c.shutdown = true
		close(c.stopCh)

		return nil
	}

	return fmt.Errorf("shutdown already in progress")
}

// GetCreatedBy return the SerializedReference in created-by annotation
func GetCreatedBy(pod *api.Pod) (*api.SerializedReference, error) {
	raw, ok := pod.Annotations["kubernetes.io/created-by"]
	if !ok {
		return nil, fmt.Errorf("no created-by annotation")
	}
	obj, err := runtime.Decode(api.Codecs.UniversalDecoder(), []byte(raw))
	if err != nil {
		return nil, err
	}
	return obj.(*api.SerializedReference), nil
}

func mergePodLabelsAndAnnotationsToPromLabels(pod *api.Pod) (map[string]string, error) {
	labels := map[string]string{}
	labelPattern := regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)
	for k, v := range pod.Labels {
		k = strings.Replace(k, "-", "_", -1)
		k = strings.Replace(k, ".", "_", -1)
		if labelPattern.MatchString(k) {
			labels[k] = v
		}
	}
	for k, v := range pod.Annotations {
		k = strings.Replace(k, "-", "_", -1)
		k = strings.Replace(k, ".", "_", -1)
		if labelPattern.MatchString(k) {
			labels[k] = v
		}
	}
	if len(labels) == 0 {
		return nil, fmt.Errorf("no valid label or annotation")
	}
	labels["pod_uid"] = string(pod.GetUID())
	labels["pod_name"] = pod.Name
	labels["namespace"] = pod.Namespace
	return labels, nil
}

// Scrape pod to rc/rs/dp map
func (c *Controller) Scrape(ch chan<- prometheus.Metric) error {
	var err error
	err = nil
	for _, obj := range c.podStore.List() {
		pod := obj.(*api.Pod)
		createdBy, _ := GetCreatedBy(pod)
		hierarchyLabels := map[string]string{
			"pod_uid":   string(pod.GetUID()),
			"pod_name":  pod.Name,
			"namespace": pod.Namespace,
		}
		if createdBy != nil && createdBy.Reference.Kind == "ReplicationController" {
			hierarchyLabels["rc_name"] = createdBy.Reference.Name
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc("kubernetes_resource_hierarchy", "Resource hierarchy of kubernetes", []string{}, hierarchyLabels),
				prometheus.GaugeValue, 1)
		} else if createdBy != nil && createdBy.Reference.Kind == "DaemonSet" {
			hierarchyLabels["ds_name"] = createdBy.Reference.Name
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc("kubernetes_resource_hierarchy", "Resource hierarchy of kubernetes", []string{}, hierarchyLabels),
				prometheus.GaugeValue, 1)
		}
		labels, err := mergePodLabelsAndAnnotationsToPromLabels(pod)
		if err == nil {
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc("kubernetes_pod_label_mapper", "Label mapper for kubernetes pods", []string{}, labels),
				prometheus.GaugeValue, 1)
		}
	}
	for key, podList := range c.dpMapper {
		_, name, err := cache.SplitMetaNamespaceKey(key)
		if err != nil {
			continue
		}
		for _, pod := range podList.Items {
			createdBy, _ := GetCreatedBy(&pod)
			labels := map[string]string{
				"pod_uid":   string(pod.GetUID()),
				"pod_name":  pod.Name,
				"namespace": pod.Namespace,
			}
			if createdBy != nil && createdBy.Reference.Kind == "ReplicaSet" {
				labels["dp_name"] = name
				labels["rs_name"] = createdBy.Reference.Name
				ch <- prometheus.MustNewConstMetric(
					prometheus.NewDesc("kubernetes_resource_hierarchy", "Resource hierarchy of kubernetes", []string{}, labels),
					prometheus.GaugeValue, 1)
			}
		}
	}
	return err
}
