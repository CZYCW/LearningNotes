from torch.utils.tensorboard import SummaryWriter
import time

writer = SummaryWriter(log_dir="./", flush_secs=10)
for i in range(100):
    time.sleep(1)
    writer.add_scalar('Loss/train', i, i)
    writer.add_scalar('Loss/test', i + 1, i)

