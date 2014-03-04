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
  def __init__(self, input_dimension, num_categories):
    # W and b are the parameters we need to learn.
    # We just initialize them with zeros; we don't need to
    # break symmetry because there are no hidden units.
    init_W = numpy.zeros((input_dimension, num_categories),
                         dtype=theano.config.floatX)
    self.W = theano.shared(value=init_W, name="W")
    init_b = numpy.zeros((num_categories,),
                         dtype=theano.config.floatX)
    self.b = theano.shared(value=init_b, name="b")

    # x will be provided anew for each batch to train, so we can leave
    # it symbolic. It should have shape: batch size * input dimension
    self.x = T.matrix("x")

    # Predict categories with a linear transform plus max. The softmax
    # is just for the purposes of gradient descent.
    # y is a matrix with shape: batch size * num categories
    self.y = T.dot(self.x, self.W) + self.b

    # y_prob is the probability predicted for each category
    # y_prob is a matrix with shape: batch size * num categories
    self.y_prob = T.nnet.softmax(self.y)

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
  def errors(self, target):
    assert target.ndim == self.predictions.ndim
    assert target.dtype.startswith("int")
    return T.mean(T.neq(target, self.predictions))
    

"""
A helper to break matrices of shared data into batches.
If the data doesn't slice evenly into batches, this just drops the
data at the end that doesn't fit into a batch.
Slices along the first dimension.
"""
class Batcher(object):
  def __init__(self, batch_size, data):
    self.batch_size = batch_size
    self.data = data

    # Rounds down to not include non-full batches.
    self.num_batches = (self.data.get_value(borrow=True).shape[0] /
                        self.batch_size)

  """
  Returns a tensor variable containing the indexth batch of data.
  """
  def batch(self, index):
    assert index < self.num_batches
    return self.data[index * self.num_batches:
                     (index + 1) * self.num_batches]

    
if __name__ == "__main__":
  # Run logistic regression on MNIST images
  # Hyperparameters
  batch_size = 600
  learning_rate = 0.13 

  ((train_input, train_output),
   (valid_input, valid_output),
   (test_input, test_output)) = datasets.mnist()

  classifier = LinearClassifier(28 * 28, 10)
