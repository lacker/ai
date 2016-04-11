#!/usr/bin/env python3
'''
Generates some data that is a subset of python for just typing in numbers.

So for example:

>>> 123
123
>>> 456
456
'''

import os
import random

MIN_NUMBER = 1
MAX_NUMBER = 127
BYTES = 10000000


'''A string for a random number.'''
def number():
  answer = random.randrange(MIN_NUMBER,
                            MAX_NUMBER + 1)
  return str(answer)

  
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

'''Multiplies some numbers.'''
def multiply_numbers():
  a = number()
  b = number()
  c = str(int(a) * int(b))
  return '>>> ' + a + ' * ' + b + '\n' + c + '\n'

def binary_multiply_numbers():
  a = int(number())
  b = int(number())
  c = a * b
  return '>{0:b}*{1:b}\n{2:b}\n'.format(a, b, c)
  
def binary_add_numbers():
  a = int(number())
  b = int(number())
  c = a + b
  return '>{0:b}+{1:b}\n{2:b}\n'.format(a, b, c)
  
'''Subtracts some numbers.'''
def subtract_numbers():
  a = number()
  b = number()
  c = str(int(a) - int(b))
  return '>>> ' + a + ' - ' + b + '\n' + c + '\n'
  
'''Mods some numbers.'''
def mod_numbers():
  a = number()
  b = number()
  c = str(int(a) % int(b))
  return '>>> ' + a + ' % ' + b + '\n' + c + '\n'

def less_than():
  a = number()
  b = number()
  c = str(int(a) < int(b))
  return '>>> ' + a + ' < ' + b + '\n' + c + '\n'
  
def greater_than():
  a = number()
  b = number()
  c = str(int(a) > int(b))
  return '>>> ' + a + ' > ' + b + '\n' + c + '\n'
  
def one_op():
  return random.choice([add_numbers, multiply_numbers, mod_numbers,
                        echo_number, subtract_numbers])()

def ineq():
  return random.choice([less_than, greater_than])()

'''Something like 5 * (3 + 2)'''
def two_ops():
  a, b, c = number(), number(), number()
  ops = ['*', random.choice('+-')]
  random.shuffle(ops)
  op1, op2 = ops
  template = random.choice([
    '%s %s (%s %s %s)',
    '(%s %s %s) %s %s'])
  problem = template % (a, op1, b, op2, c)
  d = str(eval(problem))
  return '>>> ' + problem + '\n' + d + '\n'
  
def main():
  dirname = 'data/numbers'
  for f in os.listdir(dirname):
    print('removing', dirname + '/' + f)
    os.remove(dirname + '/' + f)
  with open(dirname + '/input.txt', 'w') as f:
    written = 0
    while written < BYTES:

      text = binary_add_numbers()

      f.write(text)
      written += len(text)

      
if __name__ == '__main__':
  main()
