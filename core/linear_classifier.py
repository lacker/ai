#!/usr/bin/python
"""
This has some linear-classifier-specific stuff and some
digit-image-specific stuff.
"""

# TODO: figure out path relativity
import core.datasets as datasets

import cPickle
import numpy
import os
import PIL.Image as Image
import theano
import theano.tensor as T

"""
Turns an array into a picture of a digit and display it.
"""
def show(arr):
  image = Image.new("RGB", (28, 28))
  pixels = image.load()
  matrix = arr.reshape(28, 28)
  for x in range(28):
    for y in range(28):
      val = int(matrix.T[x][y] * 256)
      pixels[x, y] = (val, val, val)
  image.show()

"""
Turns a picture of a digit into an array.
"""
def load_pic(name="test"):
  image = Image.open(os.path.expanduser("~/data/%s.png" % name))
  image = image.resize((28, 28), Image.ANTIALIAS)
  pixels = image.load()
  matrix = numpy.zeros((28, 28), dtype="float32")
  for x in range(28):
    for y in range(28):
      val = pixels[x, y][0] / 256.0
      matrix[x][y] = val
  return matrix.T.reshape(-1)

  
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
    # x needs dimension batch size * input dimension
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
  Compiles a function to predict the category of a single vector input.
  """
  def predictor(self):
    matrix_predictor = theano.function(
      inputs=[self.x],
      outputs=self.predictions)
    return lambda v: matrix_predictor(v.reshape(1, -1))[0]
    
    
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
  def batch_error_rate(self, target):
    assert target.ndim == self.predictions.ndim
    assert target.dtype.startswith("int")
    return T.mean(T.neq(target, self.predictions))

  """
  Creates a function to calculate error rate on a dataset for a
  particular batch index.
  """
  def batch_error_rate_function(self, dataset):
    index = T.lscalar()
    y = T.ivector("y")
    return theano.function(
      inputs=[index],
      outputs=self.batch_error_rate(y),
      givens={
        self.x: dataset.input_batch(index),
        y: dataset.output_batch(index)})

  """
  Compile a function to calculate the error rate over a whole dataset.
  This goes batch-by-batch to save GPU memory.
  """
  def error_rate_function(self, dataset):
    f = self.batch_error_rate_function(dataset)
    n = dataset.num_batches
    return lambda: numpy.mean(map(f, range(n)))

    
"""
A helper to analyze an input/output dataset pair.
"""
class Dataset(object):
  def __init__(self, input_set, output_set, batch_size):
    self.input_set = input_set
    self.output_set = output_set
    self.batch_size = batch_size
    self.num_batches = (input_set.get_value(borrow=True).shape[0] /
                        batch_size)

  def input_batch(self, index):
    # Can't run this because index will be a tensor variable
    # assert index < self.num_batches
    return self.input_set[index * self.batch_size:
                          (index + 1) * self.batch_size]
    
  def output_batch(self, index):
    # Can't run this because index will be a tensor variable
    # assert index < self.num_batches
    return self.output_set[index * self.batch_size:
                          (index + 1) * self.batch_size]
    

"""
Pickles an object into the data directory.
"""
def save(obj, name):
  path = os.path.abspath(os.path.expanduser("~/data/" + name + ".pkl"))
  f = open(path, "w")
  cPickle.dump(obj, f)
  f.close()

"""
Loads something saved with 'save'.
LinearClassifier might have to be imported in the current scope for
this to work for classifiers, depending on how it was saved. Sorry.
"""
def load(name):
  path = os.path.abspath(os.path.expanduser("~/data/" + name + ".pkl"))
  f = open(path)
  answer = cPickle.load(f)
  f.close()
  return answer
  

"""
Trains a classifier, saves it to "digits", and returns it.
"""
def do_training():
  # Run logistic regression on MNIST images
  # Hyperparameters
  batch_size = 500
  learning_rate = 0.13 

  ((train_input, train_output),
   (valid_input, valid_output),
   (test_input, test_output)) = datasets.mnist()
  training = Dataset(train_input, train_output, batch_size)
  validation = Dataset(valid_input, valid_output, batch_size)
  testing = Dataset(test_input, test_output, batch_size)
  
  # We need some symbolic values for the algorithm.
  # TODO: put a lot more of this logic into the LinearClassifier
  # itself. It doesn't really make sense to include arbitrary tensor
  # variables out here.
  x = T.matrix("x")
  y = T.ivector("y")

  classifier = LinearClassifier(x, 28 * 28, 10)
  
  # Minimize this function during training
  cost = classifier.negative_log_likelihood(y)

  # Find the gradients for cost relative to the shared parameters
  W_gradient = T.grad(cost=cost, wrt=classifier.W)
  b_gradient = T.grad(cost=cost, wrt=classifier.b)

  # Update with gradient descent
  updates = [(classifier.W, classifier.W - learning_rate * W_gradient),
             (classifier.b, classifier.b - learning_rate * b_gradient)]

  # Compile a method to train one step of training
  index = T.lscalar()
  train = theano.function(
    inputs=[index],
    outputs=cost,
    updates=updates,
    givens={
      x: training.input_batch(index),
      y: training.output_batch(index)})

  runs = 100
  validator = classifier.error_rate_function(validation)
  tester = classifier.error_rate_function(testing)
  for run in range(runs):
    print "training pass", run
    for batch_index in range(training.num_batches):
      c = train(batch_index)
      if batch_index == 0:
        print "training on batch", batch_index, "had cost", c
    if run % 10 == 9:
      print "validation error rate:", validator()
  print "testing error rate:", tester()
      
  save(classifier, "digits")
  return classifier
  

if __name__ == "__main__":
  # To train anew
  c = do_training()
  print "that should have gotten to ~ cost 0.23535"
  
  # To check against the old data
  # c = load("digits-old")

  pred = c.predictor()
  print "loading a four to test..."
  pic = load_pic("four")
  print "looks like a", pred(pic)
