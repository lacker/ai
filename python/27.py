#!/usr/bin/env python

def is_prime(n):
  if n < 2:
    return False
  if n == 2:
    return True
  if n % 2 == 0:
    return False
  d = 3
  while True:
    if d * d > n:
      return True
    if n % d == 0:
      return False
    d += 2

def quadratics(a, b):
  n = 0
  while True:
    yield n * n + a * n + b
    n += 1

def consecs(a, b):
  answer = 0
  for p in quadratics(a, b):
    if is_prime(p):
      answer += 1
    else:
      return answer

max_val = 0

for abs_b in range(1, 1001):
  if not is_prime(abs_b):
    continue
  for a in range(-1000, 1001):
    for b in (-abs_b, abs_b):
      val = consecs(a, b)
      if val > max_val:
        print a, b, '=>', val
        max_val = val
