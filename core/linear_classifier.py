#!/usr/bin/python

# TODO: figure out path relativity
import datasets

import numpy
import theano
import theano.tensor as T

"""
The formula to categorize an input vector 'x' is

y = xW + b

with the largest element of y indicating which category the input
vector is.

We use logistic regression treating the classification step as a
softmax to train this classifier.
"""
class LinearClassifier(object):
  def __init__(self, x, input_dimension, num_categories):
    # W and b are the parameters we need to learn.
    # We just initialize them with zeros; we don't need to
    # break symmetry because there are no hidden units.
    init_W = numpy.zeros((input_dimension, num_categories),
                         dtype=theano.config.floatX)
    self.W = theano.shared(value=init_W, name="W")
    init_b = numpy.zeros((num_categories,),
                         dtype=theano.config.floatX)
    self.b = theano.shared(value=init_b, name="b")

    # Tensor variable for input
    self.x = x

    # Predict categories with a linear transform plus max. The softmax
    # is just for the purposes of gradient descent.
    # y is a matrix with shape: batch size * num categories
    self.y_calculated = T.dot(self.x, self.W) + self.b

    # y_prob is the probability predicted for each category
    # y_prob is a matrix with shape: batch size * num categories
    self.y_prob = T.nnet.softmax(self.y_calculated)

    # predictions is which category is the most predicted.
    # predictions is a vector with shape: batch size
    self.predictions = T.argmax(self.y_prob, axis=1)

  """
  A formula for the loss function which we are trying to minimize,
  given the target correct classification.
  The target should be an array of length batch size, since each
  member of the batch has one correct classification.
  Uses the mean instead of sum n.l.l. to be more consistent for
  different batch sizes.
  """
  def negative_log_likelihood(self, target):
    log_probs = T.log(self.y_prob)
    # Select out the probabilities that correspond to the target
    # categories
    target_log_probs = log_probs[T.arange(target.shape[0]), target]
    return -T.mean(target_log_probs)
    
  """
  A formula for the error rate in classification, given the target
  correct classification.
  The target should be an array of length batch size, since each
  member of the batch has one correct classification.
  """
  def error_rate(self, target):
    assert target.ndim == self.predictions.ndim
    assert target.dtype.startswith("int")
    return T.mean(T.neq(target, self.predictions))
    

"""
A helper to break matrices of shared data into batches.
Slices along the first dimension.
This does not bounds-check because it is surprisingly difficult to get
the size of tensor variables in general.
"""
class Batcher(object):
  def __init__(self, batch_size, data):
    self.batch_size = batch_size
    self.data = data

  """
  Returns a tensor variable containing the indexth batch of data.
  """
  def batch(self, index):
    return self.data[index * self.batch_size:
                     (index + 1) * self.batch_size]

    
if __name__ == "__main__":
  # Run logistic regression on MNIST images
  # Hyperparameters
  batch_size = 500
  learning_rate = 0.13 

  ((train_input, train_output),
   (valid_input, valid_output),
   (test_input, test_output)) = datasets.mnist()

  # We need some symbolic values for the algorithm.
  # It seems like ideally these would be part of the LinearClassifier
  # and would not be exposed.
  index = T.lscalar()
  x = T.matrix("x")
  y = T.ivector("y")

  classifier = LinearClassifier(x, 28 * 28, 10)
  
  # Minimize this function during training
  cost = classifier.negative_log_likelihood(y)

  # Create symbolic methods to calculate error rate on test and
  # validation data
  test_input_batcher = Batcher(batch_size, test_input)
  test_output_batcher = Batcher(batch_size, test_output)
  test_error_rate = theano.function(
    inputs=[index],
    outputs=classifier.error_rate(y),
    givens={
      x: test_input_batcher.batch(index),
      y: test_output_batcher.batch(index)})
  valid_input_batcher = Batcher(batch_size, valid_input)
  valid_output_batcher = Batcher(batch_size, valid_output)
  valid_error_rate = theano.function(
    inputs=[index],
    outputs=classifier.error_rate(y),
    givens={
      x: valid_input_batcher.batch(index),
      y: valid_output_batcher.batch(index)})
  
  
