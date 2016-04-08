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

MIN_NUMBER_LENGTH = 15
MAX_NUMBER_LENGTH = 20
BYTES = 10000000


'''Just a number that python echoes back.'''
def number():
  number = random.choice('123456789')
  number_length = random.randrange(MIN_NUMBER_LENGTH,
                                   MAX_NUMBER_LENGTH + 1)
  while len(number) < number_length:
    number += random.choice('0123456789')

  return '>>> ' + number + '\n' + number + '\n'
  
  
def main():
  with open('data/numbers/input.txt', 'w') as f:
    written = 0
    while written < BYTES:

      text = number()

      f.write(text)
      written += len(text)

      
if __name__ == '__main__':
  main()
