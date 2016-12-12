# Kubernetes Hierarchy Exporter

Kuberentes resource hierarchy to Prometheus bridge.

A Collector that can list and watch Kubernetes resources(mainly pod), and expose resource hierarchy. For example, a pod belongs to which deployment or daemonset.

# Building and Running

## Build
```
make
```
## Running
running outside Kuberentes(It will search for kubeconfig in ~/.kube)

```
./hierarchy_exporter --running-in-cluster=false
```

running in Kubernetes(It will use Kubernetes serviceaccount)

```
./hierarchy_exporter
```

## Collector Flags

currently no collector flags need to be configured

## General Flags

Name | Description
--- | ---
running-in-cluster | Optional, if this controller is running in a kubernetes cluster, use the pod secrets for creating a Kubernetes client. (default true)
log.level | Only log messages with the given severity or above. Valid levels: [debug, info, warn, error, fatal]. (default info)
version | Print version information

#Using Docker

You can deploy this exporter using the `cargo.caicloud.io/sysinfra/hierarchy-exporter` Docker image.

For example:

```
docker pull cargo.caicloud.io/sysinfra/hierarchy-exporter

docker run -d -p 9102:9102 -v ~/.kube/config:/root/.kube/config cargo.caicloud.io/sysinfra/hierarchy-exporter --running-in-cluster=false
```

then make requests:

```
$ curl localhost:9102/metrics
```

example response:

```
# HELP go_gc_duration_seconds A summary of the GC invocation durations.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 0.000113942
go_gc_duration_seconds{quantile="0.25"} 0.000177335
go_gc_duration_seconds{quantile="0.5"} 0.00032271000000000003
go_gc_duration_seconds{quantile="0.75"} 0.000441611
go_gc_duration_seconds{quantile="1"} 0.000594463
go_gc_duration_seconds_sum 0.002290938
go_gc_duration_seconds_count 8
# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 32
# HELP http_request_duration_microseconds The HTTP request latencies in microseconds.
# TYPE http_request_duration_microseconds summary
http_request_duration_microseconds{handler="prometheus",quantile="0.5"} 2989.553
http_request_duration_microseconds{handler="prometheus",quantile="0.9"} 7402.042
http_request_duration_microseconds{handler="prometheus",quantile="0.99"} 7402.042
http_request_duration_microseconds_sum{handler="prometheus"} 23487.557999999997
http_request_duration_microseconds_count{handler="prometheus"} 5
# HELP http_request_size_bytes The HTTP request sizes in bytes.
# TYPE http_request_size_bytes summary
http_request_size_bytes{handler="prometheus",quantile="0.5"} 533
http_request_size_bytes{handler="prometheus",quantile="0.9"} 533
http_request_size_bytes{handler="prometheus",quantile="0.99"} 533
http_request_size_bytes_sum{handler="prometheus"} 2665
http_request_size_bytes_count{handler="prometheus"} 5
# HELP http_requests_total Total number of HTTP requests made.
# TYPE http_requests_total counter
http_requests_total{code="200",handler="prometheus",method="get"} 5
# HELP http_response_size_bytes The HTTP response sizes in bytes.
# TYPE http_response_size_bytes summary
http_response_size_bytes{handler="prometheus",quantile="0.5"} 854
http_response_size_bytes{handler="prometheus",quantile="0.9"} 2978
http_response_size_bytes{handler="prometheus",quantile="0.99"} 2978
http_response_size_bytes_sum{handler="prometheus"} 10570
http_response_size_bytes_count{handler="prometheus"} 5
# HELP kubernetes_build_info A metric with a constant '1' value labeled by major, minor, git version, git commit, git tree state, build date, Go version, and compiler from which Kubernetes was built, and platform on which it is running.
# TYPE kubernetes_build_info gauge
kubernetes_build_info{buildDate="1970-01-01T00:00:00Z",compiler="gc",gitCommit="$Format:%H$",gitTreeState="not a git tree",gitVersion="v1.3.2+$Format:%h$",goVersion="go1.6.2",major="1",minor="3",platform="darwin/amd64"} 1
# HELP kubernetes_resource_hierarchy Resource hierarchy of kubernetes
# TYPE kubernetes_resource_hierarchy gauge
kubernetes_resource_hierarchy{namespace="default",pod_name="addon-manager-master-v0.0.1-j45q0",pod_uid="5dfda34f-bf5c-11e6-a2ba-0800274a55b6",rc_name="addon-manager-master-v0.0.1"} 1
kubernetes_resource_hierarchy{namespace="default",pod_name="addon-manager-mongo-v0.0.1-c8i7j",pod_uid="55585ea0-bf5c-11e6-a2ba-0800274a55b6",rc_name="addon-manager-mongo-v0.0.1"} 1
kubernetes_resource_hierarchy{namespace="default",pod_name="auth-mongo-v3.0.5-662iy",pod_uid="ef9242c9-bf5d-11e6-a2ba-0800274a55b6",rc_name="auth-mongo-v3.0.5"} 1
kubernetes_resource_hierarchy{namespace="default",pod_name="auth-redis-slave-v0.0.1-7mr0h",pod_uid="efb5ea00-bf5d-11e6-a2ba-0800274a55b6",rc_name="auth-redis-slave-v0.0.1"} 1
kubernetes_resource_hierarchy{namespace="default",pod_name="auth-redis-slave-v0.0.1-mqm9d",pod_uid="efb59301-bf5d-11e6-a2ba-0800274a55b6",rc_name="auth-redis-slave-v0.0.1"} 1
kubernetes_resource_hierarchy{namespace="default",pod_name="auth-redis-slave-v0.0.1-mzo71",pod_uid="efb67418-bf5d-11e6-a2ba-0800274a55b6",rc_name="auth-redis-slave-v0.0.1"} 1
kubernetes_resource_hierarchy{namespace="default",pod_name="auth-redis-v0.1.2-6lgqg",pod_uid="efa55f07-bf5d-11e6-a2ba-0800274a55b6",rc_name="auth-redis-v0.1.2"} 1
kubernetes_resource_hierarchy{namespace="default",pod_name="auth-server-v0.2.5-uonf2",pod_uid="efe90888-bf5d-11e6-a2ba-0800274a55b6",rc_name="auth-server-v0.2.5"} 1
kubernetes_resource_hierarchy{namespace="default",pod_name="cds-mongo-v3.0.5-v91yw",pod_uid="f79fb31f-bf5d-11e6-a2ba-0800274a55b6",rc_name="cds-mongo-v3.0.5"} 1
kubernetes_resource_hierarchy{namespace="default",pod_name="clever-admin-4gkoy",pod_uid="fc22b0d3-c036-11e6-b0d6-0800274a55b6",rc_name="clever-admin"} 1
kubernetes_resource_hierarchy{namespace="default",pod_name="clever-mongo-vl8a5",pod_uid="f019b8c4-c036-11e6-b0d6-0800274a55b6",rc_name="clever-mongo"} 1
kubernetes_resource_hierarchy{namespace="default",pod_name="cluster-admin-server-v0.1.4-yvbvv",pod_uid="f78a604e-bf5d-11e6-a2ba-0800274a55b6",rc_name="cluster-admin-server-v0.1.4"} 1
kubernetes_resource_hierarchy{namespace="default",pod_name="cluster-admin-v0.1.4-hzfmt",pod_uid="f7b89944-bf5d-11e6-a2ba-0800274a55b6",rc_name="cluster-admin-v0.1.4"} 1
kubernetes_resource_hierarchy{namespace="default",pod_name="console-web-mongo-v3.0.5-dfkaa",pod_uid="fe682439-bf5d-11e6-a2ba-0800274a55b6",rc_name="console-web-mongo-v3.0.5"} 1
kubernetes_resource_hierarchy{namespace="default",pod_name="console-web-redis-v3.0.3-yyg16",pod_uid="fe79cedf-bf5d-11e6-a2ba-0800274a55b6",rc_name="console-web-redis-v3.0.3"} 1
kubernetes_resource_hierarchy{namespace="default",pod_name="console-web-v1.4.0-dcvyy",pod_uid="00d370ca-c018-11e6-b0d6-0800274a55b6",rc_name="console-web-v1.4.0"} 1
kubernetes_resource_hierarchy{namespace="default",pod_name="hadoop-datanode-1-33qi4",pod_uid="07300435-c046-11e6-b0d6-0800274a55b6",rc_name="hadoop-datanode-1"} 1
kubernetes_resource_hierarchy{namespace="default",pod_name="hadoop-datanode-2-786no",pod_uid="0735663a-c046-11e6-b0d6-0800274a55b6",rc_name="hadoop-datanode-2"} 1
kubernetes_resource_hierarchy{namespace="default",pod_name="hadoop-datanode-3-eeztc",pod_uid="073d2d09-c046-11e6-b0d6-0800274a55b6",rc_name="hadoop-datanode-3"} 1
kubernetes_resource_hierarchy{namespace="default",pod_name="hadoop-datanode-4-ax1u1",pod_uid="07408dda-c046-11e6-b0d6-0800274a55b6",rc_name="hadoop-datanode-4"} 1
kubernetes_resource_hierarchy{namespace="default",pod_name="hadoop-datanode-controller-6611h",pod_uid="b0b62d4c-c045-11e6-b0d6-0800274a55b6",rc_name="hadoop-datanode-controller"} 1
kubernetes_resource_hierarchy{namespace="default",pod_name="hadoop-datanode-controller-8yjb1",pod_uid="b0b60e7d-c045-11e6-b0d6-0800274a55b6",rc_name="hadoop-datanode-controller"} 1
kubernetes_resource_hierarchy{namespace="default",pod_name="hadoop-datanode-controller-9zl0g",pod_uid="b0b62c52-c045-11e6-b0d6-0800274a55b6",rc_name="hadoop-datanode-controller"} 1
kubernetes_resource_hierarchy{namespace="default",pod_name="hadoop-datanode-controller-hq6o7",pod_uid="b0b60d7a-c045-11e6-b0d6-0800274a55b6",rc_name="hadoop-datanode-controller"} 1
kubernetes_resource_hierarchy{namespace="default",pod_name="hadoop-namenode-controller-ijtok",pod_uid="43b4b6af-c044-11e6-b0d6-0800274a55b6",rc_name="hadoop-namenode-controller"} 1
kubernetes_resource_hierarchy{namespace="default",pod_name="kubernetes-admin-mongo-v3.0.5-ikb8j",pod_uid="0fc05dfb-bf5e-11e6-a2ba-0800274a55b6",rc_name="kubernetes-admin-mongo-v3.0.5"} 1
kubernetes_resource_hierarchy{namespace="default",pod_name="kubernetes-admin-v0.0.10-nfucw",pod_uid="0fa85c07-bf5e-11e6-a2ba-0800274a55b6",rc_name="kubernetes-admin-v0.0.10"} 1
kubernetes_resource_hierarchy{namespace="default",pod_name="open-api-v0.0.1-f5pce",pod_uid="120758b6-bf5e-11e6-a2ba-0800274a55b6",rc_name="open-api-v0.0.1"} 1
kubernetes_resource_hierarchy{namespace="default",pod_name="paging-v0.0.5-t4imr",pod_uid="1498299d-bf5e-11e6-a2ba-0800274a55b6",rc_name="paging-v0.0.5"} 1
kubernetes_resource_hierarchy{namespace="default",pod_name="storage-mongo-7xi5d",pod_uid="05aad5df-c036-11e6-b0d6-0800274a55b6",rc_name="storage-mongo"} 1
kubernetes_resource_hierarchy{namespace="default",pod_name="storage-nuyho",pod_uid="1c78d968-c036-11e6-b0d6-0800274a55b6",rc_name="storage"} 1
kubernetes_resource_hierarchy{ds_name="kube-push-v0.0.1",namespace="kube-system",pod_name="kube-push-v0.0.1-172t4",pod_uid="c3959510-bf5c-11e6-a2ba-0800274a55b6"} 1
kubernetes_resource_hierarchy{ds_name="kube-push-v0.0.1",namespace="kube-system",pod_name="kube-push-v0.0.1-3jd61",pod_uid="c396372d-bf5c-11e6-a2ba-0800274a55b6"} 1
kubernetes_resource_hierarchy{ds_name="kube-push-v0.0.1",namespace="kube-system",pod_name="kube-push-v0.0.1-3pvcr",pod_uid="c394f789-bf5c-11e6-a2ba-0800274a55b6"} 1
kubernetes_resource_hierarchy{ds_name="kube-push-v0.0.1",namespace="kube-system",pod_name="kube-push-v0.0.1-5nwth",pod_uid="c39518b1-bf5c-11e6-a2ba-0800274a55b6"} 1
kubernetes_resource_hierarchy{ds_name="kube-push-v0.0.1",namespace="kube-system",pod_name="kube-push-v0.0.1-70g9z",pod_uid="c394f32e-bf5c-11e6-a2ba-0800274a55b6"} 1
kubernetes_resource_hierarchy{ds_name="kube-push-v0.0.1",namespace="kube-system",pod_name="kube-push-v0.0.1-7j2je",pod_uid="c3956527-bf5c-11e6-a2ba-0800274a55b6"} 1
kubernetes_resource_hierarchy{ds_name="kube-push-v0.0.1",namespace="kube-system",pod_name="kube-push-v0.0.1-cmrc3",pod_uid="c395d051-bf5c-11e6-a2ba-0800274a55b6"} 1
kubernetes_resource_hierarchy{ds_name="kube-push-v0.0.1",namespace="kube-system",pod_name="kube-push-v0.0.1-kh9hn",pod_uid="c395c7c0-bf5c-11e6-a2ba-0800274a55b6"} 1
kubernetes_resource_hierarchy{ds_name="kube-push-v0.0.1",namespace="kube-system",pod_name="kube-push-v0.0.1-nkfl0",pod_uid="c3953eef-bf5c-11e6-a2ba-0800274a55b6"} 1
kubernetes_resource_hierarchy{ds_name="kube-push-v0.0.1",namespace="kube-system",pod_name="kube-push-v0.0.1-p6emo",pod_uid="c395c2f9-bf5c-11e6-a2ba-0800274a55b6"} 1
kubernetes_resource_hierarchy{ds_name="kube-push-v0.0.1",namespace="kube-system",pod_name="kube-push-v0.0.1-y507j",pod_uid="c3958e17-bf5c-11e6-a2ba-0800274a55b6"} 1
kubernetes_resource_hierarchy{ds_name="kube-push-v0.0.1",namespace="kube-system",pod_name="kube-push-v0.0.1-yv7wa",pod_uid="c395110d-bf5c-11e6-a2ba-0800274a55b6"} 1
kubernetes_resource_hierarchy{namespace="kube-system",pod_name="addon-manager-minion-caicloud-v0.0.3-galph",pod_uid="dceba814-bf5d-11e6-a2ba-0800274a55b6",rc_name="addon-manager-minion-caicloud-v0.0.3"} 1
kubernetes_resource_hierarchy{namespace="kube-system",pod_name="addon-manager-minion-v0.0.3-un1j6",pod_uid="9ad6a9b6-bf5c-11e6-a2ba-0800274a55b6",rc_name="addon-manager-minion-v0.0.3"} 1
kubernetes_resource_hierarchy{namespace="kube-system",pod_name="elasticsearch-logging-v1.2.0-n500m",pod_uid="b04c9340-bf5c-11e6-a2ba-0800274a55b6",rc_name="elasticsearch-logging-v1.2.0"} 1
kubernetes_resource_hierarchy{namespace="kube-system",pod_name="elasticsearch-logging-v1.2.0-ohc47",pod_uid="b04c91d7-bf5c-11e6-a2ba-0800274a55b6",rc_name="elasticsearch-logging-v1.2.0"} 1
kubernetes_resource_hierarchy{namespace="kube-system",pod_name="heapster-v1.2.0-1zfbp",pod_uid="bb14a442-bf5c-11e6-a2ba-0800274a55b6",rc_name="heapster-v1.2.0"} 1
kubernetes_resource_hierarchy{namespace="kube-system",pod_name="ingress-loadbalancer-controller-v0.0.1-jnajf",pod_uid="c1cbdb0a-bf5c-11e6-a2ba-0800274a55b6",rc_name="ingress-loadbalancer-controller-v0.0.1"} 1
kubernetes_resource_hierarchy{namespace="kube-system",pod_name="kibana-logging-v1.2.1-fkr5r",pod_uid="b05dd54e-bf5c-11e6-a2ba-0800274a55b6",rc_name="kibana-logging-v1.2.1"} 1
kubernetes_resource_hierarchy{namespace="kube-system",pod_name="monitoring-grafana-v3.1.0-efdjh",pod_uid="baf4b05b-bf5c-11e6-a2ba-0800274a55b6",rc_name="monitoring-grafana-v3.1.0"} 1
kubernetes_resource_hierarchy{namespace="kube-system",pod_name="monitoring-influxdb-v1.0.0-305e9",pod_uid="bb142d15-bf5c-11e6-a2ba-0800274a55b6",rc_name="monitoring-influxdb-v1.0.0"} 1
kubernetes_resource_hierarchy{namespace="kube-system",pod_name="monitoring-server-mongo-v0.0.3-cfukp",pod_uid="bb24ccda-bf5c-11e6-a2ba-0800274a55b6",rc_name="monitoring-server-mongo-v0.0.3"} 1
kubernetes_resource_hierarchy{namespace="kube-system",pod_name="monitoring-watcher-v1.0.0-yw48x",pod_uid="bafc9899-bf5c-11e6-a2ba-0800274a55b6",rc_name="monitoring-watcher-v1.0.0"} 1
kubernetes_resource_hierarchy{ds_name="monitoring-server-v1.0.0",namespace="kube-system",pod_name="monitoring-server-v1.0.0-2eqmz",pod_uid="ba4b0999-bf5c-11e6-a2ba-0800274a55b6"} 1
kubernetes_resource_hierarchy{ds_name="monitoring-server-v1.0.0",namespace="kube-system",pod_name="monitoring-server-v1.0.0-85h1m",pod_uid="ba4b5365-bf5c-11e6-a2ba-0800274a55b6"} 1
kubernetes_resource_hierarchy{ds_name="monitoring-server-v1.0.0",namespace="kube-system",pod_name="monitoring-server-v1.0.0-8ze8w",pod_uid="ba4be577-bf5c-11e6-a2ba-0800274a55b6"} 1
kubernetes_resource_hierarchy{ds_name="monitoring-server-v1.0.0",namespace="kube-system",pod_name="monitoring-server-v1.0.0-a74cv",pod_uid="ba4cc024-bf5c-11e6-a2ba-0800274a55b6"} 1
kubernetes_resource_hierarchy{ds_name="monitoring-server-v1.0.0",namespace="kube-system",pod_name="monitoring-server-v1.0.0-bq5wb",pod_uid="ba4ad8f6-bf5c-11e6-a2ba-0800274a55b6"} 1
kubernetes_resource_hierarchy{ds_name="monitoring-server-v1.0.0",namespace="kube-system",pod_name="monitoring-server-v1.0.0-douso",pod_uid="ba4c3f90-bf5c-11e6-a2ba-0800274a55b6"} 1
kubernetes_resource_hierarchy{ds_name="monitoring-server-v1.0.0",namespace="kube-system",pod_name="monitoring-server-v1.0.0-hoks6",pod_uid="ba4db60a-bf5c-11e6-a2ba-0800274a55b6"} 1
kubernetes_resource_hierarchy{ds_name="monitoring-server-v1.0.0",namespace="kube-system",pod_name="monitoring-server-v1.0.0-jll43",pod_uid="ba4be2e0-bf5c-11e6-a2ba-0800274a55b6"} 1
kubernetes_resource_hierarchy{ds_name="monitoring-server-v1.0.0",namespace="kube-system",pod_name="monitoring-server-v1.0.0-pvnrp",pod_uid="ba4d9662-bf5c-11e6-a2ba-0800274a55b6"} 1
kubernetes_resource_hierarchy{ds_name="monitoring-server-v1.0.0",namespace="kube-system",pod_name="monitoring-server-v1.0.0-tisg4",pod_uid="ba4b88f6-bf5c-11e6-a2ba-0800274a55b6"} 1
kubernetes_resource_hierarchy{ds_name="monitoring-server-v1.0.0",namespace="kube-system",pod_name="monitoring-server-v1.0.0-xtpck",pod_uid="ba4aed1c-bf5c-11e6-a2ba-0800274a55b6"} 1
kubernetes_resource_hierarchy{ds_name="monitoring-server-v1.0.0",namespace="kube-system",pod_name="monitoring-server-v1.0.0-z88ey",pod_uid="ba4af1e2-bf5c-11e6-a2ba-0800274a55b6"} 1
kubernetes_resource_hierarchy{dp_name="circle-clair-v0.0.1",namespace="cyclone",pod_name="circle-clair-v0.0.1-4144866192-7cbvs",pod_uid="09b3e719-bf5e-11e6-a2ba-0800274a55b6",rs_name="circle-clair-v0.0.1-4144866192"} 1
kubernetes_resource_hierarchy{dp_name="circle-etcd-server-v0.0.1",namespace="cyclone",pod_name="circle-etcd-server-v0.0.1-4219074860-66olz",pod_uid="09d2676a-bf5e-11e6-a2ba-0800274a55b6",rs_name="circle-etcd-server-v0.0.1-4219074860"} 1
kubernetes_resource_hierarchy{dp_name="circle-kafka-server-v0.0.1",namespace="cyclone",pod_name="circle-kafka-server-v0.0.1-1795679798-4uu8s",pod_uid="090d8e1b-bf5e-11e6-a2ba-0800274a55b6",rs_name="circle-kafka-server-v0.0.1-1795679798"} 1
kubernetes_resource_hierarchy{dp_name="circle-mongo-server-v0.0.1",namespace="cyclone",pod_name="circle-mongo-server-v0.0.1-1816487289-600cu",pod_uid="0aaba4b8-bf5e-11e6-a2ba-0800274a55b6",rs_name="circle-mongo-server-v0.0.1-1816487289"} 1
kubernetes_resource_hierarchy{dp_name="circle-postgres-server-v0.0.1",namespace="cyclone",pod_name="circle-postgres-server-v0.0.1-3210392097-shwww",pod_uid="0928c362-bf5e-11e6-a2ba-0800274a55b6",rs_name="circle-postgres-server-v0.0.1-3210392097"} 1
kubernetes_resource_hierarchy{dp_name="circle-server-v0.0.1",namespace="cyclone",pod_name="circle-server-v0.0.1-745400083-n7hil",pod_uid="0a1cee28-bf5e-11e6-a2ba-0800274a55b6",rs_name="circle-server-v0.0.1-745400083"} 1
kubernetes_resource_hierarchy{dp_name="circle-zookeeper-server-v0.0.1",namespace="cyclone",pod_name="circle-zookeeper-server-v0.0.1-1624779439-fb69o",pod_uid="09552577-bf5e-11e6-a2ba-0800274a55b6",rs_name="circle-zookeeper-server-v0.0.1-1624779439"} 1
kubernetes_resource_hierarchy{dp_name="kube-dns-autoscaler-v21",namespace="kube-system",pod_name="kube-dns-autoscaler-v21-1598293328-8kkev",pod_uid="bfd87274-bf5c-11e6-a2ba-0800274a55b6",rs_name="kube-dns-autoscaler-v21-1598293328"} 1
kubernetes_resource_hierarchy{dp_name="kube-dns-v21",namespace="kube-system",pod_name="kube-dns-v21-1006896387-lmjq2",pod_uid="bff4d255-bf5c-11e6-a2ba-0800274a55b6",rs_name="kube-dns-v21-1006896387"} 1
kubernetes_resource_hierarchy{dp_name="kube-dns-v21",namespace="kube-system",pod_name="kube-dns-v21-1006896387-nl6ea",pod_uid="4f32cc07-bf5d-11e6-a2ba-0800274a55b6",rs_name="kube-dns-v21-1006896387"} 1
# HELP last_scrape_duration_seconds Duration of the last scrape of metrics from event store.
# TYPE last_scrape_duration_seconds gauge
last_scrape_duration_seconds 0.002482021
# HELP last_scrape_error Whether the last scrape of metrics from event store resulted in an error (1 for error, 0 for success).
# TYPE last_scrape_error gauge
last_scrape_error 0
# HELP rest_client_request_latency_microseconds Request latency in microseconds. Broken down by verb and URL
# TYPE rest_client_request_latency_microseconds summary
rest_client_request_latency_microseconds{url="https://caicloud-stack.caicloudprivatetest.com/api/v1/namespaces/%7Bnamespace%7D/pods?labelSelector=%7Bvalue%7D",verb="GET",quantile="0.5"} 13682
rest_client_request_latency_microseconds{url="https://caicloud-stack.caicloudprivatetest.com/api/v1/namespaces/%7Bnamespace%7D/pods?labelSelector=%7Bvalue%7D",verb="GET",quantile="0.9"} 48682
rest_client_request_latency_microseconds{url="https://caicloud-stack.caicloudprivatetest.com/api/v1/namespaces/%7Bnamespace%7D/pods?labelSelector=%7Bvalue%7D",verb="GET",quantile="0.99"} 48682
rest_client_request_latency_microseconds_sum{url="https://caicloud-stack.caicloudprivatetest.com/api/v1/namespaces/%7Bnamespace%7D/pods?labelSelector=%7Bvalue%7D",verb="GET"} 212606
rest_client_request_latency_microseconds_count{url="https://caicloud-stack.caicloudprivatetest.com/api/v1/namespaces/%7Bnamespace%7D/pods?labelSelector=%7Bvalue%7D",verb="GET"} 9
rest_client_request_latency_microseconds{url="https://caicloud-stack.caicloudprivatetest.com/api/v1/pods?resourceVersion=%7Bvalue%7D",verb="GET",quantile="0.5"} 5.159784e+06
rest_client_request_latency_microseconds{url="https://caicloud-stack.caicloudprivatetest.com/api/v1/pods?resourceVersion=%7Bvalue%7D",verb="GET",quantile="0.9"} 5.159784e+06
rest_client_request_latency_microseconds{url="https://caicloud-stack.caicloudprivatetest.com/api/v1/pods?resourceVersion=%7Bvalue%7D",verb="GET",quantile="0.99"} 5.159784e+06
rest_client_request_latency_microseconds_sum{url="https://caicloud-stack.caicloudprivatetest.com/api/v1/pods?resourceVersion=%7Bvalue%7D",verb="GET"} 5.159784e+06
rest_client_request_latency_microseconds_count{url="https://caicloud-stack.caicloudprivatetest.com/api/v1/pods?resourceVersion=%7Bvalue%7D",verb="GET"} 1
rest_client_request_latency_microseconds{url="https://caicloud-stack.caicloudprivatetest.com/apis/extensions/v1beta1/deployments?resourceVersion=%7Bvalue%7D",verb="GET",quantile="0.5"} 5.075102e+06
rest_client_request_latency_microseconds{url="https://caicloud-stack.caicloudprivatetest.com/apis/extensions/v1beta1/deployments?resourceVersion=%7Bvalue%7D",verb="GET",quantile="0.9"} 5.075102e+06
rest_client_request_latency_microseconds{url="https://caicloud-stack.caicloudprivatetest.com/apis/extensions/v1beta1/deployments?resourceVersion=%7Bvalue%7D",verb="GET",quantile="0.99"} 5.075102e+06
rest_client_request_latency_microseconds_sum{url="https://caicloud-stack.caicloudprivatetest.com/apis/extensions/v1beta1/deployments?resourceVersion=%7Bvalue%7D",verb="GET"} 5.075102e+06
rest_client_request_latency_microseconds_count{url="https://caicloud-stack.caicloudprivatetest.com/apis/extensions/v1beta1/deployments?resourceVersion=%7Bvalue%7D",verb="GET"} 1
# HELP rest_client_request_status_codes Number of http requests, partitioned by metadata
# TYPE rest_client_request_status_codes counter
rest_client_request_status_codes{code="200",host="caicloud-stack.caicloudprivatetest.com",method="GET"} 13
# HELP scrapes_total Total number of times event store was scraped for metrics.
# TYPE scrapes_total counter
scrapes_total 7
```
# How to use with cAdvisor metrics

for example we want to get the cpu usage of a whole deployment, the original query will be `rate(container_cpu_usage_seconds_total[5m])`, it will list all the container's cpu usage. Because cAdvisor on expose `kubernetes_pod_name` label, so we could only aggregate the metric on pod: `sum(rate(container_cpu_usage_seconds_total{kubernetes_pod_name="abc"}[5m]))`.
By using hierarchy exporter, now we can aggregate the metric on deployment/replicaset/replicationcontroller/daemonset:
`rate(container_cpu_usage_seconds_total[5m]) * on (io_kubernetes_pod_uid) group_left (kubernetes_dp_name) kuberentes_resource_mapper {kubernetes_dp_name="abcabc"}` to get all the cpu usage of container in `abcabc` deployment, and aggregate using `sum(rate(container_cpu_usage_seconds_total[5m]) * on (io_kubernetes_pod_uid) group_left (kubernetes_dp_name) kuberentes_resource_mapper {kubernetes_dp_name="abcabc"})`
