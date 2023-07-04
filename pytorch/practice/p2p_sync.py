import os
import torch
import torch.distributed as dist
import torch.multiprocessing as mp

def run(rank_id, size):
    tensor = torch.zeros(1)
    if rank_id == 0:
        tensor += 1
        # Send the tensor to process 1
        dist.send(tensor=tensor, dst=1) # distination rand id
        print('after send, Rank ', rank_id, ' has data ', tensor[0])
        dist.recv(tensor=tensor, src=1)
        print('after recv, Rank ', rank_id, ' has data ', tensor[0])
    else:
        # Receive tensor from process 0
        dist.recv(tensor=tensor, src=0)
        print('after recv, Rank ', rank_id, ' has data ', tensor[0])
        tensor += 1
        dist.send(tensor=tensor, dst=0)
        print('after send, Rank ', rank_id, ' has data ', tensor[0])


def init_process(rank_id, size, fn, backend='gloo'): # 这里的backend是指通信模块的类型，gloo是pytorch自己实现的，nccl是nvidia自己实现的
    """ Initialize the distributed environment. """
    os.environ['MASTER_ADDR'] = '127.0.0.1' # 设置主进程的ip地址
    os.environ['MASTER_PORT'] = '29500' # 设置主进程的端口号
    dist.init_process_group(backend, rank=rank_id, world_size=size) # 初始化进程组，执行网络通信模块的初始化
    # rank：为当前rank的index，用于标记当前是第几个rank，取值为0到work_size - 1之间的值；
    # world_size: 有多少个进程参与到分布式训练中;
    fn(rank_id, size)


if __name__ == "__main__":
    size = 2
    processes = []
    mp.set_start_method("spawn") # 用于指定Child process的启动方式，child process仅会继承parent process的必要resource，file descriptor和handle均不会继承。
    for rank in range(size):
        p = mp.Process(target=init_process, args=(rank, size, run))
        p.start()
        processes.append(p)

    for p in processes:
        p.join()