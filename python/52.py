#!/usr/bin/env python

n = 1
while True:
  n += 1
  digs = set(str(n))
  if (digs == set(str(2 * n)) and
      digs == set(str(3 * n)) and
      digs == set(str(4 * n)) and
      digs == set(str(5 * n)) and
      digs == set(str(6 * n))):
    print n
      
