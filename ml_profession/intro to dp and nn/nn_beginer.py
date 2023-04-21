
import torch.nn as nn

## 20 is the hidden dimension. arbitary choice
model = nn.Sequential(
        nn.Linear(3,20), # 3 for the input features x1,x2,x3
        nn.ReLU(),
        nn.Linear(20,2)) # 2 for the classes

print(model)