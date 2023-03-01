# Test

### Prepare MySQL
Insert a job record:
```sql
INSERT INTO job(job_id, job_status, dataset_name, dataset_id, user_id, project_id, project_name, command, node_group_id, node_group_name, instance_type, size, model_id, image) VALUES (30, "ResourceCreationRequestSucceed", "dataset_name", 7, 10, 1, "project_name", "python3 /opt/pytorch-mnist/mnist.py --epochs=1", 0, "ng-test","t1_micro", 1, "b5515783-3d32-4df7-adb9-5c4a42e931f2", "docker.io/kubeflowkatib/pytorch-mnist:v1beta1-45c5727"); 
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

```bash
curl -X POST -i "http://59.108.228.3:9308/api/job/create" -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Nzk1NjAyODEsImlhdCI6MTY3MDkyMDI4MSwidWlkIjoxMH0._KAelHYt5WMr1pAR8K3at27bLZVoyGwLXfNBLFYIXvA" -d '{"image": "docker.io/kubeflowkatib/pytorch-mnist:v1beta1-45c5727", "command":["python3", "/opt/pytorch-mnist/mnist.py", "--epochs=1"], "dataset_name": "dataset_name","dataset_id": "7","project_name": "project_name","project_id": "1","instance_type": "t1_micro","size": 1}'
```


{"job_id":"9","job_status":"ResourceCreationProcessComplete"}

### dockers
- `docker logs dev_joborchestrator-api_1`
- `docker logs test_resourcemanager-rpc_1`


### kafka
- kafka-console-producer --bootstrap-server localhost:9092 --topic resource-creation --property value.serializer=custom.class.serialization.JsonSerializer`

- {"job_id":"30","job_status":"ResourceCreationProcessComplete"}