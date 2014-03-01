#!/usr/bin/python

import cPickle
import gzip
import numpy
import os
import theano
import urllib

# Contains pickled Theano data for digit recognition
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

"""
Turns array data into a Theano shared variable so that Theano can
control its storage location and put it on the GPU.
Don't mutate the return value because its dependency guarantees are
unclear.
"""
def make_shared(data):
  assert data.dtype == theano.config.floatX
  return theano.shared(numpy.asarray(data), borrow=True)

"""
Turns int array data into a variable achieved by a cast from an
underlying float array so that the underlying data can be stored on
the GPU.
Don't mutate the return value.
"""
def make_int_shared(data):
  assert data.dtype == "int32"
  float_data = numpy.asarray(data, dtype=theano.config.floatX)
  shared = theano.shared(float_data, borrow=True)
  return T.cast(shared, "int32")

"""
Returns the MNIST image processing data in normal arrays.
"""
def unshared_mnist():
  f = gzip.open(data_path(MNIST), 'rb')
  answer = cPickle.load(f)
  f.close()
  return answer
  
"""
Returns the MNIST image processing data in shared memory.
"""
def mnist():
  ((train_input, train_output),
   (check_input, check_output),
   (test_input, test_output)) = unshared_mnist()

  print "mapping training data"
  s_train_input = make_shared(train_input)
  s_train_output = make_int_shared(train_output)
  print "mapping validation data"
  s_check_input = make_shared(check_input)
  s_check_output = make_int_shared(check_output)
  print "mapping testing data"
  s_test_input = make_shared(test_input)
  s_test_output = make_int_shared(test_output)
