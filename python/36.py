#!/usr/bin/env python

def is_pal(s):
  return s == ''.join(reversed(s))

answer = 0
for n in range(1000000):
  binary = '{0:b}'.format(n)
  if is_pal(str(n)) and is_pal(binary):
    print n, '=', binary
    answer += n
print
print answer
