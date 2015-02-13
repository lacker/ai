#!/usr/bin/env luajit

require "nn"
require "torch"

mnistTrain = torch.load("mnist.t7/train_32x32.t7", "ascii")
mnistTest = torch.load("mnist.t7/test_32x32.t7", "ascii")

-- A Dataset can be either training or testing.
Dataset = {}
function Dataset:new(data, labels)
  local dataset = {
    original=data,
    labels=labels
  }
  setmetatable(dataset, {__index = Dataset})
  return dataset
end

-- Creates a Dataset from an mnist-format input that has "data" and
-- "labels".
function Dataset.makeTraining(abnormal)
  local dataset = Dataset:new(abnormal.data, abnormal.labels)
  dataset.normalized = torch.FloatTensor(dataset.original:size())
  dataset.normalized:copy(dataset.original)
  dataset.mean = dataset.normalized:mean()
  dataset.std = dataset.normalized:std()
  dataset.normalized:add(-dataset.mean)
  dataset.normalized:div(dataset.std)
  return dataset
end

-- Makes a new dataset using the same transformation by which dataset
-- was originally created.
-- abnormal should have "data" and "labels".
function Dataset:makeTest(abnormal)
  local test = Dataset:new(abnormal.data, abnormal.labels)
  test.normalized = torch.FloatTensor(test.original:size())
  test.normalized:copy(test.original)
  test.normalized:add(-self.mean)
  test.normalized:div(self.std)
  return test
end

-- Create a linear regression model to train on the training data
function makeModel(dataset)
  local ninputs = dataset.normalized:stride(1)
  local m = nn.Sequential()
  m:add(nn.Reshape(ninputs))
  m:add(nn.Linear(ninputs, 10))
  m:add(nn.LogSoftMax())
  return m
end

train = Dataset.makeTraining(mnistTrain)
test = train:makeTest(mnistTest)
model = makeModel(train)

-- Ghetto testing
assert(string.format("%.4f", test.normalized[3][1][4][2]) == "-0.3635")
