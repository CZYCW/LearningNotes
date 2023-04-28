import torch
import torch.nn as nn

class Model(nn.Module): 
    def __init__(self): 
        super(Model, self).__init__() 
        self.linear1 = nn.Linear(2, 3) 
        self.linear2 = nn.Linear(3, 2) 
  
    def forward(self, x): 
        h = torch.sigmoid(self.linear1(x)) 
        o = torch.sigmoid(self.linear2(h)) 
        return o

model= Model() 
X = torch.randn((1, 2)) 
print(X)
Y = model(X) 
print(Y)