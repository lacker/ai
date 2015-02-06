#!/usr/bin/env luajit

require "torch"

train = torch.load("mnist.t7/train_32x32.t7", "ascii")
test = torch.load("mnist.t7/test_32x32.t7", "ascii")

-- Usually for normalization, AI stuff uses zero mean and unit norm
-- TODO: write helpers to do that.


-- Finds the global mean for all pixels
function globalMean(tensor)
  return tensor:sum() / tensor:nElement()
end

print(train.data:sum())
print(train.data:nElement())

print(globalMean(train.data))
