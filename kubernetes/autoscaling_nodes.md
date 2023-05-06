### Context
Usage of HorizontalPodAutoscaler (HPA) is one of the most critical aspects of making a resilient, fault-tolerant, and highly-available system. However, it is of no use if there are no nodes with available resources. When Kubernetes cannot schedule new Pods because there’s not enough available memory or CPU, new Pods will be unschedulable and in the pending status. If we do not increase the capacity of our cluster, pending Pods might stay in that state indefinitely. 

### Cluster Autoscaler
- purpose: to adjust size of the cluster by adding or removing worker nodes
  - add new nodes when pods cannot be scheduled due to insufficient resources
  - eliminate nodes when they are underuntilize for a period