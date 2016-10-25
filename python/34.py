#!/usr/bin/env python

def fac(n):
  if n == 0:
    return 1
  return n * fac(n - 1)
  
fmap = [fac(n) for n in range(10)]

for i in range(10, 10000000):
  z = sum(fmap[int(char)] for char in str(i))
  if z == i:
    print z
