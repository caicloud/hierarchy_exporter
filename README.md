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
go_gc_duration_seconds{quantile="0"} 0.00010721600000000001
go_gc_duration_seconds{quantile="0.25"} 0.00013281500000000002
go_gc_duration_seconds{quantile="0.5"} 0.000165117
go_gc_duration_seconds{quantile="0.75"} 0.000246668
go_gc_duration_seconds{quantile="1"} 0.000306199
go_gc_duration_seconds_sum 0.0029041740000000003
go_gc_duration_seconds_count 16
# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 29
# HELP http_request_duration_microseconds The HTTP request latencies in microseconds.
# TYPE http_request_duration_microseconds summary
http_request_duration_microseconds{handler="prometheus",quantile="0.5"} 2527.476
http_request_duration_microseconds{handler="prometheus",quantile="0.9"} 4674.808
http_request_duration_microseconds{handler="prometheus",quantile="0.99"} 4893.489
http_request_duration_microseconds_sum{handler="prometheus"} 61762.674999999996
http_request_duration_microseconds_count{handler="prometheus"} 20
# HELP http_request_size_bytes The HTTP request sizes in bytes.
# TYPE http_request_size_bytes summary
http_request_size_bytes{handler="prometheus",quantile="0.5"} 656
http_request_size_bytes{handler="prometheus",quantile="0.9"} 656
http_request_size_bytes{handler="prometheus",quantile="0.99"} 656
http_request_size_bytes_sum{handler="prometheus"} 13120
http_request_size_bytes_count{handler="prometheus"} 20
# HELP http_requests_total Total number of HTTP requests made.
# TYPE http_requests_total counter
http_requests_total{code="200",handler="prometheus",method="get"} 20
# HELP http_response_size_bytes The HTTP response sizes in bytes.
# TYPE http_response_size_bytes summary
http_response_size_bytes{handler="prometheus",quantile="0.5"} 2398
http_response_size_bytes{handler="prometheus",quantile="0.9"} 2413
http_response_size_bytes{handler="prometheus",quantile="0.99"} 2414
http_response_size_bytes_sum{handler="prometheus"} 47924
http_response_size_bytes_count{handler="prometheus"} 20
# HELP kubernetes_build_info A metric with a constant '1' value labeled by major, minor, git version, git commit, git tree state, build date, Go version, and compiler from which Kubernetes was built, and platform on which it is running.
# TYPE kubernetes_build_info gauge
kubernetes_build_info{buildDate="1970-01-01T00:00:00Z",compiler="gc",gitCommit="$Format:%H$",gitTreeState="not a git tree",gitVersion="v1.3.2+$Format:%h$",goVersion="go1.6.2",major="1",minor="3",platform="darwin/amd64"} 1
# HELP kubernetes_resource_hierarchy Resource hierarchy of kubernetes
# TYPE kubernetes_resource_hierarchy gauge
kubernetes_resource_hierarchy{io_kubernetes_pod_uid="10f87c63-a6de-11e6-8cb1-0800274a55b6",kubernetes_namespace="kube-system",kubernetes_pod_name="kube-dns-v12-br7bs",kubernetes_rc_name="kube-dns-v12"} 1
kubernetes_resource_hierarchy{io_kubernetes_pod_uid="11052115-a6de-11e6-8cb1-0800274a55b6",kubernetes_namespace="kube-system",kubernetes_pod_name="monitoring-server-mongo-v0.0.3-zvmqy",kubernetes_rc_name="monitoring-server-mongo-v0.0.3"} 1
kubernetes_resource_hierarchy{io_kubernetes_pod_uid="16e6f86b-a6de-11e6-8cb1-0800274a55b6",kubernetes_namespace="kube-system",kubernetes_pod_name="monitoring-influxdb-v1.0.0-m05hk",kubernetes_rc_name="monitoring-influxdb-v1.0.0"} 1
kubernetes_resource_hierarchy{io_kubernetes_pod_uid="19b270bb-b13e-11e6-a34d-0800274a55b6",kubernetes_ds_name="sysdig-agent",kubernetes_namespace="kube-system",kubernetes_pod_name="sysdig-agent-x1rrw"} 1
kubernetes_resource_hierarchy{io_kubernetes_pod_uid="19b27210-b13e-11e6-a34d-0800274a55b6",kubernetes_ds_name="sysdig-agent",kubernetes_namespace="kube-system",kubernetes_pod_name="sysdig-agent-cq1t4"} 1
kubernetes_resource_hierarchy{io_kubernetes_pod_uid="19b277c2-b13e-11e6-a34d-0800274a55b6",kubernetes_ds_name="sysdig-agent",kubernetes_namespace="kube-system",kubernetes_pod_name="sysdig-agent-7o7wm"} 1
kubernetes_resource_hierarchy{io_kubernetes_pod_uid="19b27abd-b13e-11e6-a34d-0800274a55b6",kubernetes_ds_name="sysdig-agent",kubernetes_namespace="kube-system",kubernetes_pod_name="sysdig-agent-af0w7"} 1
kubernetes_resource_hierarchy{io_kubernetes_pod_uid="19b295ed-b13e-11e6-a34d-0800274a55b6",kubernetes_ds_name="sysdig-agent",kubernetes_namespace="kube-system",kubernetes_pod_name="sysdig-agent-v1rlz"} 1
kubernetes_resource_hierarchy{io_kubernetes_pod_uid="1ccd6369-a6de-11e6-8cb1-0800274a55b6",kubernetes_namespace="default",kubernetes_pod_name="admin-mongo-qidtl",kubernetes_rc_name="admin-mongo"} 1
kubernetes_resource_hierarchy{io_kubernetes_pod_uid="1ce0e7a5-a6de-11e6-8cb1-0800274a55b6",kubernetes_namespace="kube-system",kubernetes_pod_name="monitoring-grafana-v3.1.0-ntfah",kubernetes_rc_name="monitoring-grafana-v3.1.0"} 1
kubernetes_resource_hierarchy{io_kubernetes_pod_uid="654e6927-b537-11e6-97fc-0800274a55b6",kubernetes_ds_name="monitoring-server-v1.0.0",kubernetes_namespace="kube-system",kubernetes_pod_name="monitoring-server-v1.0.0-d11fe"} 1
kubernetes_resource_hierarchy{io_kubernetes_pod_uid="65773561-b537-11e6-97fc-0800274a55b6",kubernetes_ds_name="monitoring-server-v1.0.0",kubernetes_namespace="kube-system",kubernetes_pod_name="monitoring-server-v1.0.0-s77v3"} 1
kubernetes_resource_hierarchy{io_kubernetes_pod_uid="659a4081-b537-11e6-97fc-0800274a55b6",kubernetes_ds_name="monitoring-server-v1.0.0",kubernetes_namespace="kube-system",kubernetes_pod_name="monitoring-server-v1.0.0-9k3id"} 1
kubernetes_resource_hierarchy{io_kubernetes_pod_uid="65b64174-b537-11e6-97fc-0800274a55b6",kubernetes_ds_name="monitoring-server-v1.0.0",kubernetes_namespace="kube-system",kubernetes_pod_name="monitoring-server-v1.0.0-a77us"} 1
kubernetes_resource_hierarchy{io_kubernetes_pod_uid="65d8f3e4-b537-11e6-97fc-0800274a55b6",kubernetes_ds_name="monitoring-server-v1.0.0",kubernetes_namespace="kube-system",kubernetes_pod_name="monitoring-server-v1.0.0-ud9zq"} 1
kubernetes_resource_hierarchy{io_kubernetes_pod_uid="665019e4-b537-11e6-97fc-0800274a55b6",kubernetes_ds_name="monitoring-server-v1.0.0",kubernetes_namespace="kube-system",kubernetes_pod_name="monitoring-server-v1.0.0-fsf2s"} 1
kubernetes_resource_hierarchy{io_kubernetes_pod_uid="80f931f8-a6e0-11e6-81cd-0800274a55b6",kubernetes_namespace="kube-system",kubernetes_pod_name="heapster-v1.2.0-wy4uv",kubernetes_rc_name="heapster-v1.2.0"} 1
kubernetes_resource_hierarchy{io_kubernetes_pod_uid="8881c080-b51a-11e6-97fc-0800274a55b6",kubernetes_namespace="kube-system",kubernetes_pod_name="elasticsearch-logging-v1-gfpff",kubernetes_rc_name="elasticsearch-logging-v1"} 1
kubernetes_resource_hierarchy{io_kubernetes_pod_uid="88824214-b51a-11e6-97fc-0800274a55b6",kubernetes_namespace="kube-system",kubernetes_pod_name="elasticsearch-logging-v1-1ahcd",kubernetes_rc_name="elasticsearch-logging-v1"} 1
kubernetes_resource_hierarchy{io_kubernetes_pod_uid="88824598-b51a-11e6-97fc-0800274a55b6",kubernetes_namespace="kube-system",kubernetes_pod_name="elasticsearch-logging-v1-vmosc",kubernetes_rc_name="elasticsearch-logging-v1"} 1
kubernetes_resource_hierarchy{io_kubernetes_pod_uid="888285b8-b51a-11e6-97fc-0800274a55b6",kubernetes_namespace="kube-system",kubernetes_pod_name="elasticsearch-logging-v1-iuh13",kubernetes_rc_name="elasticsearch-logging-v1"} 1
kubernetes_resource_hierarchy{io_kubernetes_pod_uid="8884ded1-b51a-11e6-97fc-0800274a55b6",kubernetes_namespace="kube-system",kubernetes_pod_name="elasticsearch-logging-v1-00hwp",kubernetes_rc_name="elasticsearch-logging-v1"} 1
kubernetes_resource_hierarchy{io_kubernetes_pod_uid="8d688be3-b53b-11e6-97fc-0800274a55b6",kubernetes_ds_name="sysdig-agent",kubernetes_namespace="kube-system",kubernetes_pod_name="sysdig-agent-e5m5k"} 1
kubernetes_resource_hierarchy{io_kubernetes_pod_uid="f052344d-acb7-11e6-a986-0800274a55b6",kubernetes_namespace="kube-system",kubernetes_pod_name="kibana-logging-v1.2.1-4v1t9",kubernetes_rc_name="kibana-logging-v1.2.1"} 1
kubernetes_resource_hierarchy{io_kubernetes_pod_uid="0d28f073-a6de-11e6-8cb1-0800274a55b6",kubernetes_dp_name="allenyyy",kubernetes_namespace="allen",kubernetes_pod_name="allenyyy-150039532-ely15",kubernetes_rs_name="allenyyy-150039532"} 1
kubernetes_resource_hierarchy{io_kubernetes_pod_uid="0d5bfb4b-a6de-11e6-8cb1-0800274a55b6",kubernetes_dp_name="qme-test",kubernetes_namespace="test1",kubernetes_pod_name="qme-test-3739904865-ysuig",kubernetes_rs_name="qme-test-3739904865"} 1
kubernetes_resource_hierarchy{io_kubernetes_pod_uid="376e25e4-b073-11e6-a820-0800274a55b6",kubernetes_dp_name="nginx-pc",kubernetes_namespace="allen",kubernetes_pod_name="nginx-pc-535175895-sxhet",kubernetes_rs_name="nginx-pc-535175895"} 1
kubernetes_resource_hierarchy{io_kubernetes_pod_uid="6ba722da-b20d-11e6-bd6e-0800274a55b6",kubernetes_dp_name="mongo",kubernetes_namespace="allen",kubernetes_pod_name="mongo-1392715141-661eb",kubernetes_rs_name="mongo-1392715141"} 1
kubernetes_resource_hierarchy{io_kubernetes_pod_uid="7b449ec0-b207-11e6-bd6e-0800274a55b6",kubernetes_dp_name="nginx-pc",kubernetes_namespace="allen",kubernetes_pod_name="nginx-pc-535175895-4nv51",kubernetes_rs_name="nginx-pc-535175895"} 1
kubernetes_resource_hierarchy{io_kubernetes_pod_uid="8513472a-b32a-11e6-bd6e-0800274a55b6",kubernetes_dp_name="fileserver",kubernetes_namespace="storage-1480090199",kubernetes_pod_name="fileserver-3254767661-7wfz5",kubernetes_rs_name="fileserver-3254767661"} 1
kubernetes_resource_hierarchy{io_kubernetes_pod_uid="98ae7e09-a7eb-11e6-a986-0800274a55b6",kubernetes_dp_name="qme",kubernetes_namespace="test1",kubernetes_pod_name="qme-3429656455-4ket9",kubernetes_rs_name="qme-3429656455"} 1
kubernetes_resource_hierarchy{io_kubernetes_pod_uid="9b5a7b93-b085-11e6-a820-0800274a55b6",kubernetes_dp_name="test-btn",kubernetes_namespace="allen",kubernetes_pod_name="test-btn-4010109795-eg9n4",kubernetes_rs_name="test-btn-4010109795"} 1
kubernetes_resource_hierarchy{io_kubernetes_pod_uid="c48c4243-b206-11e6-bd6e-0800274a55b6",kubernetes_dp_name="test1001",kubernetes_namespace="test1",kubernetes_pod_name="test1001-4288028312-4todg",kubernetes_rs_name="test1001-4288028312"} 1
kubernetes_resource_hierarchy{io_kubernetes_pod_uid="d12cdfad-b50c-11e6-bd6e-0800274a55b6",kubernetes_dp_name="tf-ps0",kubernetes_namespace="distributed-tensorflow-1480297736",kubernetes_pod_name="tf-ps0-1120799526-n1f73",kubernetes_rs_name="tf-ps0-1120799526"} 1
kubernetes_resource_hierarchy{io_kubernetes_pod_uid="d1502975-b50c-11e6-bd6e-0800274a55b6",kubernetes_dp_name="tensorboard",kubernetes_namespace="distributed-tensorflow-1480297736",kubernetes_pod_name="tensorboard-1613150639-vggp6",kubernetes_rs_name="tensorboard-1613150639"} 1
kubernetes_resource_hierarchy{io_kubernetes_pod_uid="e15132fe-b50c-11e6-bd6e-0800274a55b6",kubernetes_dp_name="tf-ps0",kubernetes_namespace="distributed-tensorflow-1480297763",kubernetes_pod_name="tf-ps0-1121389350-v56nd",kubernetes_rs_name="tf-ps0-1121389350"} 1
kubernetes_resource_hierarchy{io_kubernetes_pod_uid="e18674ad-b50c-11e6-bd6e-0800274a55b6",kubernetes_dp_name="tensorboard",kubernetes_namespace="distributed-tensorflow-1480297763",kubernetes_pod_name="tensorboard-1613740463-f9c8e",kubernetes_rs_name="tensorboard-1613740463"} 1
# HELP last_scrape_duration_seconds Duration of the last scrape of metrics from event store.
# TYPE last_scrape_duration_seconds gauge
last_scrape_duration_seconds 0.0009874600000000001
# HELP last_scrape_error Whether the last scrape of metrics from event store resulted in an error (1 for error, 0 for success).
# TYPE last_scrape_error gauge
last_scrape_error 0
# HELP rest_client_request_latency_microseconds Request latency in microseconds. Broken down by verb and URL
# TYPE rest_client_request_latency_microseconds summary
rest_client_request_latency_microseconds{url="https://sysinfra.caicloudprivatetest.com/api/v1/namespaces/%7Bnamespace%7D/pods?labelSelector=%7Bvalue%7D",verb="GET",quantile="0.5"} 8244
rest_client_request_latency_microseconds{url="https://sysinfra.caicloudprivatetest.com/api/v1/namespaces/%7Bnamespace%7D/pods?labelSelector=%7Bvalue%7D",verb="GET",quantile="0.9"} 12217
rest_client_request_latency_microseconds{url="https://sysinfra.caicloudprivatetest.com/api/v1/namespaces/%7Bnamespace%7D/pods?labelSelector=%7Bvalue%7D",verb="GET",quantile="0.99"} 28810
rest_client_request_latency_microseconds_sum{url="https://sysinfra.caicloudprivatetest.com/api/v1/namespaces/%7Bnamespace%7D/pods?labelSelector=%7Bvalue%7D",verb="GET"} 300808
rest_client_request_latency_microseconds_count{url="https://sysinfra.caicloudprivatetest.com/api/v1/namespaces/%7Bnamespace%7D/pods?labelSelector=%7Bvalue%7D",verb="GET"} 30
rest_client_request_latency_microseconds{url="https://sysinfra.caicloudprivatetest.com/api/v1/pods?resourceVersion=%7Bvalue%7D",verb="GET",quantile="0.5"} 190149
rest_client_request_latency_microseconds{url="https://sysinfra.caicloudprivatetest.com/api/v1/pods?resourceVersion=%7Bvalue%7D",verb="GET",quantile="0.9"} 190149
rest_client_request_latency_microseconds{url="https://sysinfra.caicloudprivatetest.com/api/v1/pods?resourceVersion=%7Bvalue%7D",verb="GET",quantile="0.99"} 190149
rest_client_request_latency_microseconds_sum{url="https://sysinfra.caicloudprivatetest.com/api/v1/pods?resourceVersion=%7Bvalue%7D",verb="GET"} 190149
rest_client_request_latency_microseconds_count{url="https://sysinfra.caicloudprivatetest.com/api/v1/pods?resourceVersion=%7Bvalue%7D",verb="GET"} 1
rest_client_request_latency_microseconds{url="https://sysinfra.caicloudprivatetest.com/apis/extensions/v1beta1/deployments?resourceVersion=%7Bvalue%7D",verb="GET",quantile="0.5"} 124108
rest_client_request_latency_microseconds{url="https://sysinfra.caicloudprivatetest.com/apis/extensions/v1beta1/deployments?resourceVersion=%7Bvalue%7D",verb="GET",quantile="0.9"} 124108
rest_client_request_latency_microseconds{url="https://sysinfra.caicloudprivatetest.com/apis/extensions/v1beta1/deployments?resourceVersion=%7Bvalue%7D",verb="GET",quantile="0.99"} 124108
rest_client_request_latency_microseconds_sum{url="https://sysinfra.caicloudprivatetest.com/apis/extensions/v1beta1/deployments?resourceVersion=%7Bvalue%7D",verb="GET"} 124108
rest_client_request_latency_microseconds_count{url="https://sysinfra.caicloudprivatetest.com/apis/extensions/v1beta1/deployments?resourceVersion=%7Bvalue%7D",verb="GET"} 1
# HELP rest_client_request_status_codes Number of http requests, partitioned by metadata
# TYPE rest_client_request_status_codes counter
rest_client_request_status_codes{code="200",host="sysinfra.caicloudprivatetest.com",method="GET"} 47
# HELP scrapes_total Total number of times event store was scraped for metrics.
# TYPE scrapes_total counter
scrapes_total 22
```
# How to use with cAdvisor metrics

for example we want to get the cpu usage of a whole deployment, the original query will be `rate(container_cpu_usage_seconds_total[5m])`, it will list all the container's cpu usage. Because cAdvisor on expose `kubernetes_pod_name` label, so we could only aggregate the metric on pod: `sum(rate(container_cpu_usage_seconds_total{kubernetes_pod_name="abc"}[5m]))`.
By using hierarchy exporter, now we can aggregate the metric on deployment/replicaset/replicationcontroller/daemonset:
`rate(container_cpu_usage_seconds_total[5m]) * on (io_kubernetes_pod_uid) group_left (kubernetes_dp_name) kuberentes_resource_mapper {kubernetes_dp_name="abcabc"}` to get all the cpu usage of container in `abcabc` deployment, and aggregate using `sum(rate(container_cpu_usage_seconds_total[5m]) * on (io_kubernetes_pod_uid) group_left (kubernetes_dp_name) kuberentes_resource_mapper {kubernetes_dp_name="abcabc"})`
