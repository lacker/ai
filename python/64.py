#!/usr/bin/env python

from math import sqrt

def gcf(a, b):
  if a < 0:
    a = -a
  if b < 0:
    b = -b
  if a < b:
    a, b = b, a
  if b == 0:
    return a
  return gcf(b, a % b)
  

'''
A number of the form (a + b sqrt(c)) / d where those are ints
'''
class Num:
  def __init__(self, a, b, c, d):
    self.a = a
    self.b = b
    self.c = c
    self.d = d
    
    if self.d < 0:
      self.a = -self.a
      self.b = -self.b
      self.d = -self.d

    factor = gcf(gcf(self.a, self.b), self.d)
    self.a /= factor
    self.b /= factor
    self.d /= factor

  def key(self):
    return (self.a, self.b, self.c, self.d)

  def is_positive(self):
    if self.a >= 0 and self.b >= 0:
      return True
    if self.a <= 0 and self.b <= 0:
      return False
    if self.a * self.a > self.b * self.b * self.c:
      return self.a > 0
    return self.a < 0
    
  def floor(self):
    guess = 0
    while True:
      new_guess = guess + 1
      if not self.addint(-new_guess).is_positive():
        return guess
      guess = new_guess

  def addint(self, i):
    return Num(self.a + i * self.d, self.b, self.c, self.d)

  '''
  d / (a + b root c) = d * (a - b root c) / (a^2 - cb^2)
  = (ad - bd root c) / (a^2 - cb^2)
  '''
  def invert(self):
    return Num(self.a * self.d, -self.b * self.d, self.c,
               self.a * self.a - self.c * self.b * self.b)

'''
Yield the (term, thing) tuples of
x1 + 1 / (x2 + 1 / (x3 + ....
for a Num
term is each term like xi
"thing" is the value of the cf starting at that term
'''
def cf_terms(x):
  thing = x
  while True:
    term = thing.floor()
    yield (term, thing)
    rest = thing.addint(-term)
    thing = rest.invert()

'''
Find the period of the square root of x
'''
def sqrt_period(x):
  s = Num(0, 1, x, 1)
  i = -1
  seen = {}
  for term, thing in cf_terms(s):
    i += 1
    key = thing.key()
    if key in seen:
      return i - seen[key]
    seen[key] = i

def is_square(x):
  s = sqrt(x)
  return int(s) == s

answer = 0
for x in range(10001):
  if not is_square(x):
    p = sqrt_period(x)
    print x, 'has period', p
    if p % 2 == 1:
      answer += 1
print 'total:', answer
