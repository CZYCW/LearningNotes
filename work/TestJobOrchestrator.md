# Test

### Prepare MySQL
Insert a job record:
```sql
INSERT INTO job(job_id, job_name, job_description, job_status, dataset_name, dataset_id, user_id, project_id, project_name, launch_command, node_group_id, node_group_name, instance_type, number_of_instance, model_id, pod_names, image, ssd) VALUES (30, "job_name", "job_description", "ResourceCreationRequestSucceed", "dataset_name", 7, 10, 1, "project_name", "python3 /opt/pytorch-mnist/mnist.py --epochs=1", 0, "ng-test","t1_micro", 1, "b5515783-3d32-4df7-adb9-5c4a42e931f2", "pod1 pod2", "docker.io/kubeflowkatib/pytorch-mnist:v1beta1-45c5727", 1); 
```

### Add a column to mysql

```sql
ALTER TABLE node_group
ADD COLUMN job_id VARCHAR(225);
```


### Send request to job orchestrator 
```bash
curl -X GET -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzczMDc0ODIsImlhdCI6MTY3NzIyMTA4MiwidWlkIjoxfQ.1wFINymXRlsGvyyoBZFUgNuCpGlvdamYTTTH5fJi2sc" -i "http://127.0.0.1:9300/api/user/info" -H "Content-Type: application/json" -d '
{              
    "userId": 1
}'
```
**CreateJob**
```bash
curl -X POST -i "http://59.108.228.3:9408/api/job/create" -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Nzk1NjAyODEsImlhdCI6MTY3MDkyMDI4MSwidWlkIjoxMH0._KAelHYt5WMr1pAR8K3at27bLZVoyGwLXfNBLFYIXvA" -d '{"jobName": "job_name", "jobDescription": "description", "image": "docker.io/kubeflowkatib/pytorch-mnist:v1beta1-45c5727", "launchCommand":"python3 /opt/pytorch-mnist/mnist.py --epochs=1", "datasetName": "dataset_name","datasetId": "7","projectName": "project_name","projectId": "1","instanceType": "t1_micro","numberOfInstance": , "ssd": 1}'
```

{"job_id":"9","job_status":"ResourceCreationProcessComplete"}

**GetJobInfo**
```bash
curl -X POST -i "http://59.108.228.3:9408/api/job/info" -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Nzk1NjAyODEsImlhdCI6MTY3MDkyMDI4MSwidWlkIjoxMH0._KAelHYt5WMr1pAR8K3at27bLZVoyGwLXfNBLFYIXvA" -d '{"jobId": "10"}'
```

**GetJobList**
```bash
curl -X GET -i "http://59.108.228.3:9408/api/job/list" -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Nzk1NjAyODEsImlhdCI6MTY3MDkyMDI4MSwidWlkIjoxMH0._KAelHYt5WMr1pAR8K3at27bLZVoyGwLXfNBLFYIXvA"
```

### Send request to job manager

grpcurl -plaintext -d '{"node_group": "
t2-micro-2-cfvde8t6k51qis3db3n0", "image": "docker.io/kubeflowkatib/pytorch-mnist:v1beta1-45c5727", "command": "python3 /opt/pytorch-mnist/mnist.py --epochs=1", "dataset_id":"7", "project_id": "1","model_id": "7","user_id": "10","nodegroup_quantity": "2","version": "t1_micro","global_job_id": 3}' 59.108.228.3:9306 jobmanager.JobManager/create


### dockers
- `docker logs -f dev_joborchestrator-api_1`
- `docker logs test_resourcemanager-rpc_1`


### kafka
- kafka-console-producer --bootstrap-server localhost:9092 --topic resource-creation --property value.serializer=custom.class.serialization.JsonSerializer`

- {"job_id":"30","job_status":"ResourceCreationProcessComplete"}