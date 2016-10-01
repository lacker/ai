#!/usr/bin/env python

from math import sqrt

'''
Finds the square root but casts to an integer.
Works on big ints.
Could be off by 1 or 2
'''
def int_sqrt(n):
  guess = 1
  while True:
    new_guess = (guess + n / guess) / 2
    if abs(guess - new_guess) < 2:
      return max(guess, new_guess)
    guess = new_guess
  

'Sum of the first 100 post-decimal digits of the square root of n'
def sodos(n):
  return sum(map(int, str(int_sqrt(n * 10 ** 220))[:100]))

def is_square(n):
  x = int(sqrt(n))
  return x * x == n
  
print sodos(2)

answer = 0
for x in range(1, 101):
  if not is_square(x):
    answer += sodos(x)

print answer
  
