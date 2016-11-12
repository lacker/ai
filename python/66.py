#!/usr/bin/env python

from decimal import *
from math import floor, sqrt

getcontext().prec = 100

'''
Yield the terms of
x1 + 1 / (x2 + 1 / (x3 + ....
'''
def cf_terms(x):
  thing = x
  while True:
    term = int(floor(thing))
    yield term
    rest = thing - term
    if rest == 0:
      return
    thing = 1 / rest

'Returns a (numerator, denominator) tuple'
def make_cf(values):
  if not values:
    raise Exception('values is zeroy')
    
  i = len(values) - 1

  rest_numer, rest_denom = values[i], 1

  while i > 0:
    i -= 1
    numer = values[i] * rest_numer + rest_denom
    denom = rest_numer
    rest_numer = numer
    rest_denom = denom

  return rest_numer, rest_denom
    
'''
Given d, find the (x, y) solution to
x^2 - dy^2 = 1
x^2 = 1 + dy^2
with minimum x.
It's kind of x/y = sqrt(d)-rounded-down.
'''
def solve(d):
  s = Decimal(d).sqrt()
  if int(s) == s:
    return None, None
  terms = []
  for term in cf_terms(s):
    terms.append(term)
    x, y = make_cf(terms)
    if x * x - d * y * y == 1:
      return x, y

biggest_x = 0
for d in range(1001):
  x, y = solve(d)
  if (x is not None and x >= biggest_x) or d < 20:
    biggest_x = x
    print 'x:', x, 'y:', y, 'd:', d
