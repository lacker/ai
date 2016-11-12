#!/usr/bin/env python

from math import floor, sqrt

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
  if len(values) == 1:
    return (values[0], 1)

  rest_numer, rest_denom = make_cf(values[1:])

  numer = values[0] * rest_numer + rest_denom
  denom = rest_numer
  return (numer, denom)
    
'''
Given d, find the (x, y) solution to
x^2 - dy^2 = 1
x^2 = 1 + dy^2
with minimum x.
It's kind of x/y = sqrt(d)-rounded-down.
'''
def solve(d):
  s = sqrt(d)
  if int(s) == s:
    return None, None
  terms = []
  for term in cf_terms(sqrt(d)):
    terms.append(term)
    try:
      x, y = make_cf(terms)
      if x * x - d * y * y == 1:
        return x, y
    except:
      print d, 'is bugged'
      return None, None

biggest_x = 0
for d in range(1001):
  x, y = solve(d)
  if x is not None and x >= biggest_x:
    biggest_x = x
    print 'x:', x, 'y:', y, 'd:', d
