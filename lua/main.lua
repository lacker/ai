#!/usr/bin/env luajit

require "nn"
require "torch"

mnistTrain = torch.load("mnist.t7/train_32x32.t7", "ascii")
mnistTest = torch.load("mnist.t7/test_32x32.t7", "ascii")

-- Creates a Dataset from an mnist-format input that has "data" and
-- "labels".
-- TODO: make "Dataset" a real class
function makeTrainingDataset(abnormal)
  local normalized = torch.FloatTensor(abnormal.data:size())
  normalized:copy(abnormal.data)
  local mean = normalized:mean()
  local std = normalized:std()
  normalized:add(-mean)
  normalized:div(std)
  return {
    original=abnormal.data,
    labels=abnormal.labels,
    mean=mean,
    std=std,
    normalized=normalized,
  }
end

-- Makes a new dataset using the same transformation by which dataset
-- was originally created.
-- abnormal should have "data" and "labels".
function makeTestDataset(dataset, abnormal)
  local normalized = torch.FloatTensor(abnormal.data:size())
  normalized:copy(abnormal.data)
  normalized:add(-dataset.mean)
  normalized:div(dataset.std)
  return {
    original=abnormal.data,
    labels=abnormal.labels,
    normalized=normalized,
  }
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

train = makeTrainingDataset(mnistTrain)
test = makeTestDataset(train, mnistTest)
model = makeModel(train)

-- Ghetto testing
assert(string.format("%.4f", test.normalized[3][1][4][2]) == "-0.3635")
