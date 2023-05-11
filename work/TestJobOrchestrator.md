# Test

### Prepare MySQL
Insert a job record:
```sql
INSERT INTO job(job_id, job_name, job_description, job_status, dataset_name, dataset_id, user_id, project_id, project_name, launch_command, node_group_id, node_group_name, instance_type, instance_info, number_of_instance, model_id, pod_names, image, ssd) VALUES (30, "job_name", "job_description", "ResourceCreationRequestSucceed", "dataset_name", 7, 10, 1, "project_name", "python3 /opt/pytorch-mnist/mnist.py --epochs=1", 0, "ng-test","t1_micro", "{\"Description\":\"t2.micro\",\"MemoryInGB\":1,\"NumberOfGpu\":1,\"NumberOfCpu\":1,\"StorageType\":\"\",\"StorageInGB\":0,\"NumberOfNetworkInterface\":2,\"OnDemandLinuxPrice\":0.0116}",1, "b5515783-3d32-4df7-adb9-5c4a42e931f2", "pod1 pod2", "docker.io/kubeflowkatib/pytorch-mnist:v1beta1-45c5727", 1); 
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
curl -X POST -i "http://0.0.0.0:8087/api/job/info" -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODM3MDI3MTEsImlhdCI6MTY4MzYxNjMxMSwidWlkIjoxfQ.gZZeucr-KRzavwVv_pYpLboCIOadrcHmn7EhaO4jbu4" -d '{"jobId": "32"}'
```

**GetJobList**
```bash
curl -X GET -i "http://59.108.228.3:9308/api/job/list" -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Nzk3MzMwODUsImlhdCI6MTY3OTY0NjY4NSwidWlkIjoxfQ.reDSDopTbdKg5IlOK-3aD4gf0atScYAo5Be7YmdKeeE"
```

**DeleteJob**
```bash
curl -X POST -i "http://59.108.228.3:9308/api/job/archive" -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Nzk1NjAyODEsImlhdCI6MTY3MDkyMDI4MSwidWlkIjoxMH0._KAelHYt5WMr1pAR8K3at27bLZVoyGwLXfNBLFYIXvA" -d '{"jobId": "3"}'
```

**StopJob**
```bash
curl -X POST -i "http://59.108.228.3:9308/api/job/stop" -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Nzk3MzMwODUsImlhdCI6MTY3OTY0NjY4NSwidWlkIjoxfQ.reDSDopTbdKg5IlOK-3aD4gf0atScYAo5Be7YmdKeeE" -d '{"jobId": "1"}'
```

**DownloadModel**
```bash
curl -X POST -i "http://10.0.10.53:8087/api/job/modeldownload" -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODMzNjIxOTAsImlhdCI6MTY4MzI3NTc5MCwidWlkIjoxfQ.8PXr05acmMSeu7__vyigZ5eovzU9ajF7dOoMlcEsX-U" -d '{"jobId": "20"}'
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


### resource manager
```
grpcurl -plaintext -d '{"instance_type": 3, "size": 2, "job_id": "10"}' 59.108.228.3:9404 resourcemanager.ResourceManager/create
```


## JobOrchestrator

db.createUser({
  user: "root",
  pwd: "1cRfzdfi22",
  roles: [ { role: "root", db: "admin" } ]
})


### test against cluster
```bash
curl -X POST -i "http://0.0.0.0:8087/api/job/create" -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODA1ODk5ODYsImlhdCI6MTY4MDUwMzU4NiwidWlkIjoxfQ.Aow6YUBEJ7xQKWK8YvH9eM1xZvOIyTiGHUCFQCtjcPE" -d '{"jobName": "job_name", "jobDescription": "description", "image": "docker.io/kubeflowkatib/pytorch-mnist:v1beta1-45c5727", "launchCommand":"python3 /opt/pytorch-mnist/mnist.py --epochs=1", "datasetName": "dataset_name","datasetId": "1","projectName": "project_name","projectId": "1","instanceType": "t2_medium","numberOfInstance": 1, "ssd": 1}'
```

### Test with a t2_medium
```bash
curl -X POST -i "http://0.0.0.0:8087/api/job/create" -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODM3MDI3MTEsImlhdCI6MTY4MzYxNjMxMSwidWlkIjoxfQ.gZZeucr-KRzavwVv_pYpLboCIOadrcHmn7EhaO4jbu4" -d '{"jobName": "job_name", "jobDescription": "description", "image": "docker.io/kubeflowkatib/pytorch-mnist:v1beta1-45c5727", "launchCommand":"python3 /opt/pytorch-mnist/mnist.py --epochs=1", "datasetName": "dataset_name","datasetId": "1","projectId": "1","instanceType": "t2_medium","numberOfInstance": 1, "ssd": 1}'
```

### Test Colossal AI image
```bash
curl -X POST -i "http://59.108.228.3:9408/api/job/create" -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Nzk1NjAyODEsImlhdCI6MTY3MDkyMDI4MSwidWlkIjoxMH0._KAelHYt5WMr1pAR8K3at27bLZVoyGwLXfNBLFYIXvA" -d '{"jobName": "job_name", "jobDescription": "description", "image": "hpcaitech/colossalai:0.2.5", "launchCommand":"colossalai run --nproc_per_node 1 ../mnt/project/train.py --config ../mnt/project/config.py --optimizer lars --synthetic", "datasetName": "dataset_name","datasetId": "1","projectName": "project_name","projectId": "1","instanceType": "g5_xlarge","numberOfInstance": 1, "ssd": 1}'
```

```bash
curl -X POST -i "http://59.108.228.3:9408/api/job/create" -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Nzk1NjAyODEsImlhdCI6MTY3MDkyMDI4MSwidWlkIjoxMH0._KAelHYt5WMr1pAR8K3at27bLZVoyGwLXfNBLFYIXvA" -d '{"jobName": "job_name", "jobDescription": "description", "image": "hpcaitech/colossalai:0.2.5", "launchCommand":"colossalai run --nproc_per_node 1 ../mnt/project/train.py --config ../mnt/project/config.py --optimizer lars --path ../mnt/dataset/", "datasetName": "dataset_name","datasetId": "1","projectName": "project_name","projectId": "1","instanceType": "g5_xlarge","numberOfInstance": 1, "ssd": 1}'
```


### Remove a container
- joborch  `docker compose -p test --env-file ./deploy/docker-compose/.test.env rm -sv joborchestrator-api`
- rm `docker compose -p dev --env-file ./deploy/docker-compose/.dev.env rm -sv resourcemanager-rpc`

grpcurl -plaintext -d '{ 
    "node_group": "test-node-group-3",
    "image": "docker.io/kubeflowkatib/pytorch-mnist:v1beta1-45c5727",
    "command": "python3 /opt/pytorch-mnist/mnist.py --epochs=1",
    "dataset_id":"7",
    "project_id": "1",
    "model_id": "7",
    "user_id": "10",
    "nodegroup_quantity": "2",
    "version": "v1",
    "global_job_id": "8",
    "instance_info": {
        "NumberOfGpu": 0
    }
}' localhost:8080 jobmanager.JobManager/create



curl -X POST -i "http://0.0.0.0:8087/api/job/create" -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODIxNDU0MzYsImlhdCI6MTY4MjA1OTAzNiwidWlkIjoxfQ.hrfpvD4tDYCMofBEMvbuVqi95mpwp-1RaHvlTUgaRcA" -d '{"jobName": "job_name", "jobDescription": "description", "image": "docker.io/kubeflowkatib/pytorch-mnist:v1beta1-45c5727", "launchCommand":"python3 /opt/pytorch-mnist/mnist.py --epochs=1", "datasetName": "dataset_name","datasetId": "1","projectName": "project_name","projectId": "1","instanceType": "t2_medium","numberOfInstance": 1, "ssd": 1}'


grpcurl -plaintext -d '{                                          ✘ INT  base 15:01:07
    "node_group": "test-node-group-4",
    "image": "docker.io/kubeflowkatib/pytorch-mnist:v1beta1-45c5727",
    "command": "python3 /opt/pytorch-mnist/mnist.py --epochs=1",
    "dataset_id":"7",
    "project_id": "1",
    "model_id": "7",
    "user_id": "10",
    "nodegroup_quantity": "2",
    "version": "v1",
    "global_job_id": "9",
    "instance_info": {
        "NumberOfGpu": 0
    }
}' localhost:8089 jobmanager.JobManager/create


### E2E Test Colossal with model
curl -X POST -i "http://59.108.228.3:9308/api/job/create" -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Nzk1NjAyODEsImlhdCI6MTY3MDkyMDI4MSwidWlkIjoxMH0._KAelHYt5WMr1pAR8K3at27bLZVoyGwLXfNBLFYIXvA" -d '{"jobName": "job_name", "jobDescription": "description", "image": "hpcaitech/colossalai:0.2.5", "launchCommand":"colossalai run --nproc_per_node 1 ../mnt/project/train.py --config ../mnt/project/config.py --optimizer lars --path ../mnt/dataset/", "datasetName": "dataset_name","datasetId": "1","projectName": "project_name","projectId": "1","instanceType": "g5_xlarge","numberOfInstance": 1, "ssd": 1}'

```bash
grpcurl -plaintext -d '{ 
    "node_group": "test-colossal",
    "image": "hpcaitech/colossalai:0.2.5",
    "command": "cd / && colossalai run --nproc_per_node 1 /mnt/project/large_batch_optimizer/train.py --config /mnt/project/large_batch_optimizer/config.py --optimizer lars --dataset /mnt/dataset/ --output /output/model",
    "dataset_s3_path": "luchen-storage:10/dataset/1",
    "project_s3_path": "luchen-storage:10/project/2",
    "nodegroup_quantity": "1",
    "training_job_name": "job1",
    "global_job_id": "1",
    "global_job_id": "9",
    "input_model_s3_path": "luchen-storage:10/model/8",
    "instance_info": {
        "NumberOfGpu": 1
    }
}' localhost:8089 jobmanager.JobManager/create;
```

```
grpcurl -plaintext -d '{
    "model_name": "test-colossal",
    "model_s3_location": "bucket:1",
    "user_id": 1
}' 0.0.0.0:8080 modelmanager.ModelManager.register
```

grpcurl -plaintext -d '{                                                                                                                   base 
    "model_name": "test-colossal"
    "model_s3_location": "bucket:1"
    "user_id": 1
    
}' localhost:8080 modelmanager.ModelManager/create;