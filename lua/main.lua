#!/usr/bin/env luajit

require "image"
require "nn"
require "torch"

-- Making this work might save memory
-- torch.setdefaulttensortype("torch.FloatTensor")

torch.manualSeed(1337)

mnistTrain = torch.load("mnist.t7/train_32x32.t7", "ascii")
mnistTrain.data:resize(60000, 32, 32)
mnistTest = torch.load("mnist.t7/test_32x32.t7", "ascii")
mnistTest.data:resize(10000, 32, 32)

-- itorch-displays a 2d tensor as an image.
function show(tensor)
  local image = torch.Tensor(tensor)
  image:resize(1, tensor:size()[1], tensor:size()[2])
  itorch.image(image)
end

-- Slices a 1d byte tensor along the first dimension
function sliceBytes(tensor, first, last)
  return torch.ByteTensor(tensor:storage(), first,
                          torch.LongStorage{last - first + 1})
end

-- Slices a 3d tensor along the first dimension
function slice3D(tensor, first, last)
  local outsize = torch.LongStorage{
    last - first + 1,
    tensor:size(2),
    tensor:size(3),
  }
  return torch.Tensor(tensor:storage(),
                      1 + (first - 1) * tensor:stride(1),
                      outsize)
end

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
  dataset.normalized = torch.Tensor(dataset.original:size())
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
  test.normalized = torch.Tensor(test.original:size())
  test.normalized:copy(test.original)
  test.normalized:add(-self.mean)
  test.normalized:div(self.std)
  return test
end

-- A Net is a neural net with helper functions
Net = {}
function Net:new(trainingDataset)
  local net = {
    data=trainingDataset,
  }
  setmetatable(net, {__index = Net})

  net:makeDeepModel()
  -- net:makeLinearModel()

  return net
end

-- Input size is not hard coded here
function Net:makeLinearModel()
  -- The model to train
  local ninputs = self.data.normalized:stride(1)
  self.model = nn.Sequential()
  self.model:add(nn.Reshape(ninputs))
  self.model:add(nn.Linear(ninputs, 10))
  self.model:add(nn.LogSoftMax())

  self.criterion = nn.ClassNLLCriterion()
end

-- This makes a deep model a la:
-- http://torch.cogbits.com/doc/tutorials_supervised/
-- The 32x32-ness is hard coded here.
function Net:makeDeepModel()
  -- Reshape to have an extra size-1 dimension because standard torch
  -- convolution stuff expects multiple feature planes, a la rgb.
  self.model = nn.Sequential()
  self.model:add(nn.Reshape(1, 32, 32))

  -- Layer 1, convolutional
  self.model:add(nn.SpatialConvolution(1, 16, 5, 5))
  self.model:add(nn.Tanh())
  self.model:add(nn.SpatialMaxPooling(2, 2, 2, 2))

  -- Layer 2, convolutional
  self.model:add(nn.SpatialConvolution(16, 128, 5, 5))
  self.model:add(nn.Tanh())
  self.model:add(nn.SpatialMaxPooling(2, 2, 2, 2))

  -- Layer 3, a normal classification net
  self.model:add(nn.Reshape(128 * 5 * 5))
  self.model:add(nn.Linear(128 * 5 * 5, 200))
  self.model:add(nn.Tanh())
  self.model:add(nn.Linear(200, 10))
  self.model:add(nn.LogSoftMax())

  self.criterion = nn.ClassNLLCriterion()

end

-- Trains on a single input-output pair, or a batch.
--
-- If there is a batch of input-outputs, the first dimension of input
-- and label should have the same size and each index should represent
-- another data point.
--
-- input should be a tensor with the input data.
-- label should just be a number with the digit+1 (stupid 1-indexing)
function Net:train(input, label)
  local predicted = self.model:forward(input)
  -- TODO: why do we throw this error away?
  local err = self.criterion:forward(predicted, label)
  self.model:zeroGradParameters()
  local t = self.criterion:backward(predicted, label)
  self.model:backward(input, t)
  self.model:updateParameters(0.01)
end

function Net:trainIndex(i)
  self:train(self.data.normalized[i], self.data.labels[i])
end

function Net:trainRange(first, last)
  local dataBatch = slice3D(self.data.normalized, first, last)
  local labelBatch = sliceBytes(self.data.labels, first, last)
  self:train(dataBatch, labelBatch)
end

-- Trains against each data point in one big batch.
function Net:trainAllBatch()
  local start = os.time()
  self:trainRange(1, self.data.normalized:size(1))
  print(string.format("%d seconds elapsed", os.time() - start))
end

-- Trains against each data point one by one.
function Net:trainAllOneByOne()
  local start = os.time()
  for i = 1,self.data.normalized:size(1) do
    self:trainIndex(i)
  end
  print(string.format("%d seconds elapsed", os.time() - start))
end

-- Trains against each data point with minibatches.
function Net:trainAllMiniBatch()
  local start = os.time()
  local step = 100
  for i = 1,self.data.normalized:size(1),step do
    self:trainRange(i, i + step - 1)
  end
  print(string.format("%d seconds elapsed", os.time() - start))
end

-- Print performance on the provided dataset.
function Net:test(dataset)
  local right = 0
  local wrong = 0
  for i = 1,dataset.normalized:size(1) do
    local label = self:bestLabel(dataset.normalized[i])
    if dataset.labels[i] == label then
      right = right + 1
    else
      wrong = wrong + 1
    end
  end

  print("right:", right)
  print("wrong:", wrong)
end

-- Print performance on the provided dataset.
-- Uses a batch for efficiency.
function Net:testMiniBatch(dataset, first, last)
  local inputs = slice3D(dataset.normalized, first, last)
  local predictedLabels = self:bestLabels(inputs)
  local targetLabels = sliceBytes(dataset.labels, first, last)
  local size = predictedLabels:size()[1]
  if size ~= targetLabels:size()[1] then
    error()
  end
  local right = predictedLabels:eq(targetLabels):sum()
  -- TODO: print stuff
end

-- Returns the classification scores for labels
function Net:classify(input)
  return self.model:forward(input)
end

-- Returns the best label for a picture
function Net:bestLabel(input)
  local classes = self:classify(input)
  local m,i = classes:max(1)
  return i[1]
end

-- Returns the best labels for a set of pictures
function Net:bestLabels(inputs)
  local classified = self:classify(inputs)
  local vals,indexes = classified:max(2)
  indexes:resize(indexes:size(1))
  return indexes
end

-- Returns the best digit for a picture
function Net:bestDigit(input)
  return self:bestLabel(input) - 1
end

-- Shows an example of a particular class via random permutation
function Net:example(digit)
  local pic = torch.rand(32, 32):add(-0.5):mul(2)
  local label = digit + 1
  local picScore = self:classify(pic)[label]

  for i = 1,500 do
    local newPic = torch.rand(32, 32):add(-0.5):mul(0.1)
    newPic:add(pic)
    local newPicScore = self:classify(newPic)[label]
    if newPicScore > picScore then
      pic = newPic
      picScore = newPicScore
    end
  end

  return pic
end

-- Shows an average of examples of a particular class
function Net:averageExample(digit)
  local num = 100
  local sum = torch.Tensor(32, 32):zero()
  for i = 1,num do
    sum:add(self:example(digit))
  end
  return sum:div(num)
end

-- Shows the pixels that, when just this pixel is activated, we think
-- this is the right class
function Net:pixels(digit)
  local pic = torch.Tensor(32, 32):zero()

  for i = 1,32 do
    for j = 1,32 do
      local pixelPic = torch.Tensor(32, 32):zero()
      pixelPic[i][j] = 1
      if self:bestDigit(pixelPic) == digit then
        pic[i][j] = 1
      end
    end
  end

  return pic
end

train = Dataset.makeTraining(mnistTrain)
test = train:makeTest(mnistTest)
net = Net:new(train)

-- Ghetto testing
assert(string.format("%.4f", test.normalized[3][4][2]) == "-0.3635")
