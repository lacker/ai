#!/usr/bin/env python

from itertools import islice

def nth(iterable, n):
  return next(islice(iterable, n - 1, n))

def champers():
  n = 0
  while True:
    n += 1
    for ch in str(n):
      yield int(ch)

def foo(n):
  ans = nth(champers(), n)
  print n, '->', ans
  return ans

foo(11)
foo(12)
  
print reduce(lambda a, b: a * b, map(foo, [
  1,
  10,
  100,
  1000,
  10000,
  100000,
  1000000,
  ]))


