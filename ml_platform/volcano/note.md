## Scheduler

### Intro
- responsible for pod scheduling
- consist of a series of actions and plugins
  - action: action that should be executed in each step
  - plugin: provide the action algorithm details in difference scenarios


### Action
- enqueue
  - filter out tasks that meet the scheduling requirements bases a a series of filter algorithm
  - send them to to-be-schedule queue -> task status changes from pending to inqueue
- allocate
  - selection the most suitable node based on a series of predictions and optimization algorithm
- preempt
  - preemptive scheduling of high priority tasks in the same queue according to priority rules. 
- reclaim
  - reclaiming the resources allocated to the clsuuter based on the queue weight when a new tasks enters the queue
  - the cluster resource cannot meet the need of the queue
- backfill
  - backfill tasks in the pending state into cluster node to maximize the resource utilization of the node

### plugins
- gang
  - considers task not in READY state (including Binding, Bound, Rnning, ALlocated, Succced and Pipelined) have a higher priority
  - checks whether the resources allocated to the queue can meet the resources required by the task to run mim-available pods after trying to evicts soem pods and reclain resource
- conformance
  - considers that the tasks in namespace kube-system have a higher priority. These tasks will not be preempted.

- DRF
  - considers that task with fewer resource have a high priority

- nodeorder
  - score all nodes for a task by using a series of scoring algorithm