#!/usr/bin/env python3
'''
Generates some data that is a subset of python for just typing in numbers.

So for example:

>>> 123
123
>>> 456
456
'''

import random

MIN_NUMBER_LENGTH = 1
MAX_NUMBER_LENGTH = 4
BYTES = 10000000


'''A string for a random number.'''
def number():
  number = random.choice('123456789')
  number_length = random.randrange(MIN_NUMBER_LENGTH,
                                   MAX_NUMBER_LENGTH + 1)
  while len(number) < number_length:
    number += random.choice('0123456789')

  return number

  
'''Just a number that python echoes back.'''
def echo_number():
  n = number()

  return '>>> ' + n + '\n' + n + '\n'

'''Adds some numbers.'''
def add_numbers():
  a = number()
  b = number()
  c = str(int(a) + int(b))
  return '>>> ' + a + ' + ' + b + '\n' + c + '\n'

  
def main():
  with open('data/numbers/input.txt', 'w') as f:
    written = 0
    while written < BYTES:

      text = add_numbers()

      f.write(text)
      written += len(text)

      
if __name__ == '__main__':
  main()
