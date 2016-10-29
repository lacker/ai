#!/usr/bin/env python

def find_pandig(base):
  pandig = ''
  for n in range(1, 10):
    pandig += str(base * n)
    if len(pandig) > 9:
      if n < 3:
        raise 'done'
      return None
    if len(pandig) == 9:
      if '0' not in pandig and len(set(pandig)) == 9:
        return pandig
      else:
        return None
  raise 'fail'


for base in range(1, 100000):
  pandig = find_pandig(base)
  if pandig:
    print base, '->', pandig
