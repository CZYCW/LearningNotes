## torch.distributed

### Components
- DistributedDataParallel
  - model is replicated on every process
  - every model replica gets a chunk of the input data
  - DDP takes cases of gradient communication to keep model replicas synchronized and overlaps it with gradient computations to speed up training
- RPC based distributed training
  - support distributed pipeline parallelism, parameter server paradigm
  - helps manager remote object lifetime and extends the autograd engine beyond machine boundary

- Collective Communication (c10d)
  - support sending tensors across processes within a group.
  - offers both collective communication APIs (all_reduce and all_together) and P2P commnunication APIs(send and isend)

### backends
- use NCCL backend for GPU training
- use Gloo for CPU training

