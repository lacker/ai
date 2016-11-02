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

def best_prime(length):
  window = []
  n = 1
  while True:
    n += 1
    if is_prime(n):
      window.append(n)
      if len(window) > length:
        window = window[1:]
        if length % 2 == 0 and window[0] != 2:
          break
      if len(window) == length:
        s = sum(window)
        if s >= 1000000:
          return
        if is_prime(s):
          print length, window, s
          return
        
length = 20
while True:
  length += 1
  best_prime(length)
