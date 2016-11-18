#!/usr/bin/env python

from math import sqrt

FACTOR_CACHE = {}

def factorize(n):
  if n in FACTOR_CACHE:
    return FACTOR_CACHE[n]
  if n == 1:
    return []
  for k in [2] + range(3, int(sqrt(n)) + 1, 2):
    if n % k == 0:
      answer = [k] + factorize(n / k)
      break
  else:
    answer = [n]
  FACTOR_CACHE[n] = answer
  return answer

'''
Returns a map where multiplying key^value together gives you n
And each key is a prime
'''
def unique_factorize(n):
  answer = {}
  for key in factorize(n):
    if key in answer:
      answer[key] += 1
    else:
      answer[key] = 1
  return answer


def phi(n):
  answer = 1
  for key, value in unique_factorize(n).items():
    answer *= (key - 1) * (key ** (value - 1))
  return answer

  

def isperm(a, b):
  return sorted(list(str(a))) == sorted(list(str(b)))
  

best_ratio = 10
best_n = None
for n in range(2, 10000000):
  p = phi(n)
  if isperm(p, n):
    print 'phi(%d) = %d' % (n, p)
    ratio = float(n) / p
    if ratio < best_ratio:
      print 'ratio =', ratio
      best_ratio = ratio
      best_n = n
print best_n

  
