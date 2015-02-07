#!/usr/bin/env luajit

require "torch"

train = torch.load("mnist.t7/train_32x32.t7", "ascii")
test = torch.load("mnist.t7/test_32x32.t7", "ascii")

-- Finds the global mean for all pixels
function globalMean(tensor)
  return tensor:sum() / tensor:nElement()
end

-- Normalizes data with the provided mean to have a unit norm.
-- The first dimension of the tensor is assumed to index the different
-- inputs for the data.
function normalize(tensor, mean)
  -- Apply the negative mean as an offset
  normalized = torch.add(tensor, -mean)

  -- Square to start calculating norm
  squares = torch.pow(normalized, 2)

  -- Sum across the first index and square root to get norms
  norms = TODO

  -- Divide the original tensor elements by norms to normalize
  -- TODO
end


mean = globalMean(train.data)

print(mean)
