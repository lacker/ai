#!/usr/bin/env python

CACHE = {}

def is_prime(n):
  if n < 2:
    return False
  if n == 2:
    return True
  if n % 2 == 0:
    return False
  if n in CACHE:
    return CACHE[n]
  d = 3
  while True:
    if d * d > n:
      CACHE[n] = True
      return True
    if n % d == 0:
      CACHE[n] = False
      return False
    d += 2

def left_truncate(n):
  trunk = str(n)[1:]
  if not trunk:
    return 0
  return int(trunk)

def right_truncate(n):
  trunk = str(n)[:-1]
  if not trunk:
    return 0
  return int(trunk)

def is_left_trunk_prime(n):
  left = left_truncate(n)
  if left == 0:
    return True
  if not is_prime(left):
    return False
  return is_left_trunk_prime(left_truncate(n))

def is_right_trunk_prime(n):
  right = right_truncate(n)
  if right == 0:
    return True
  if not is_prime(right):
    return False
  return is_right_trunk_prime(right_truncate(n))

def is_double_trunk_prime(n):
  return is_prime(n) and is_left_trunk_prime(n) and is_right_trunk_prime(n)

count = 0
addem = 0

n = 10
while True:
  if is_double_trunk_prime(n):
    count += 1
    addem += n
    print count, ':', n, ':', addem
  n += 1
  
