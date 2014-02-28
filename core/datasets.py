#!/usr/bin/python

import os
import urllib

MNIST = "http://www.iro.umontreal.ca/~lisa/deep/data/mnist/mnist.pkl.gz"

"""
Gives us a local data path for the data available at the particular URL source.
"""
def data_path(source):
  _, fname = os.path.split(source)
  answer = os.path.abspath(os.path.expanduser("~/data/" + fname))
  if not os.path.isfile(answer):
    # We need to download it
    print "Downloading data from " + source
    urllib.urlretrieve(source, answer)
  return answer
  
