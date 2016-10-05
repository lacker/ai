#!/usr/bin/env python

from math import sqrt

factor_map = {}

def factorize(n):
  if n in factor_map:
    return factor_map[n]
  if n == 1:
    return []
  for k in [2] + range(3, int(sqrt(n)) + 1, 2):
    if n % k == 0:
      answer = [k] + factorize(n / k)
      break
  else:
    answer = [n]
  factor_map[n] = answer
  return answer


fours = []
for n in range(2, 100000000):
  factors = factorize(n)
  print n, 'factorizes to', factors
  is_a_four = (len(set(factors)) == 4)
  fours.append(is_a_four)
  if fours[-4:] == [True, True, True, True]:
    break
else:
  print 'fail'
  
