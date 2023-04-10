## A Starter
Horizonatal Pod Autoscaler
- automatically scale number of pods in a deployment or statefulset. 
- As a kubernetes API resource and a controller. 
- periodically adjusts the number of replicas in a statefulset or a deployment to match the observed CPU utilization to the target specified by user.

Metrics Server
- **Heapster**: a tool that enables container cluster monitoring and performance analysis.
  - collects and interprets various metrics like resource usage, events and so on.
  - it is considered deprecated in favor of metrics-server.
- What is metrics server
  - it collects information about used resources (memory, CPU) of nodes and pods. 
  - doesn't store metrics
  - provide API that can be used to retrieve current resource usage
  - collect cluster wide metrics and allow us to retrieve them through its api. 
- flow
  - periodically fetch metrics with CPU and memory usage from kubelets running on nodes.
  - other entities requrest data from metrics server through api server which has the master metrics api. 

Retrieve metrics on nodes:
- `kubectl top nodes`
