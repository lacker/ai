#!/usr/bin/env luajit

require "nn"
require "torch"

train = torch.load("mnist.t7/train_32x32.t7", "ascii")
test = torch.load("mnist.t7/test_32x32.t7", "ascii")

-- Finds the global mean for all pixels
function globalMean(tensor)
  return tensor:sum() / tensor:nElement()
end

-- Converts an input tensor to something with mean 0 and 1 standard
-- deviation.
-- Returns the initial mean and std used to normalize as well.
function normalize(inputTensor)
  output = torch.FloatTensor(inputTensor:size())
  output:copy(inputTensor)
  mean = output:mean()
  std = output:std()
  output:add(-mean)
  output:div(std)
  return {mean=mean, std=std, data=output}
end

-- Create a linear regression model to train on the training data
function makeModel(trainingData)
  ninputs = trainingData:stride(1)
  m = nn.Sequential()
  m:add(nn.Reshape(ninputs))
  m:add(nn.Linear(ninputs, 10))
  m:add(nn.LogSoftMax())
  return m
end

model = makeModel(train.data)