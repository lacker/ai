#!/usr/bin/python

# TODO: figure out path relativity
import datasets

import theano

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
    # it symbolic.
    # TODO: figure out why we do or do not need x's dimensions here
    self.x = T.matrix("x")

    # Predict categories with a linear transform plus max. The softmax
    # is just for the purposes of gradient descent.
    self.y = T.dot(self.x, self.W) + self.b

    # y_prob is the probability of each category
    self.y_prob = T.nnet.softmax(self.y)

    # y_pred is a 1 in the predicted category, 0 in others
    self.y_pred = T.argmax(self.y_prob, axis=1)
    
    
if __name__ == "__main__":
  # Run logistic regression on MNIST images
  batch_size = 600
  classifier = LinearClassifier(batch_size, 28 * 28, 10)
