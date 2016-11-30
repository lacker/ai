#!/usr/bin/python

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

# Sum of proper divisors
def sopd(n):
  answer = 1
  for key, value in unique_factorize(n).items():
    answer *= (key ** (value + 1) - 1) / (key - 1)
  answer -= n
  return answer

best_length = 0
best_n = -1

for start in range(1, 1000001):
  length = 1
  n = start
  seen = set([n])
  while True:
    n = sopd(n)
    if n == start:
      print 'chain', n, 'has length', length
      if length > best_length:
        best_length = length
        best_n = n
        print '***', n
      break
    if n > 1000000:
      break
    if n in seen:
      # Loops without including start
      break
    if n < start:
      break
    seen.add(n)
    length += 1

print 'END', best_n, 'has length', best_length
