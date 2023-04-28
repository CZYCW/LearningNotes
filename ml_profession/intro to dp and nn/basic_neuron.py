import torch


def neuron(input):
  weights = torch.Tensor([0.5, 0.5, 0.5])
  b = torch.Tensor([0.5])
  return torch.add(torch.matmul(input, weights), b)

if __name__ == '__main__':
  input = torch.Tensor([1, 2, 3])
  print(neuron(input))