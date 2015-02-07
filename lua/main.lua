#!/usr/bin/env luajit

require "torch"

train = torch.load("mnist.t7/train_32x32.t7", "ascii")
test = torch.load("mnist.t7/test_32x32.t7", "ascii")

-- Finds the global mean for all pixels
function globalMean(tensor)
  return tensor:sum() / tensor:nElement()
end

-- Converts a byte tensor to a double format.
-- The range 0-255 gets mapped to 0-1.
-- This takes a slice of the mnist training or test data and converts
-- it to something readily displayable.
function normalize(byteTensor)
  output = torch.Tensor(byteTensor:size())
  output:copy(byteTensor)
  output:div(255)
  return output
end

