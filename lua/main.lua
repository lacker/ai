#!/usr/bin/env luajit

require "torch"

train = torch.load("mnist.t7/train_32x32.t7", "ascii")
test = torch.load("mnist.t7/test_32x32.t7", "ascii")

-- Usually for normalization, AI stuff uses zero mean and unit norm

-- TODO: tools for mean and norm
