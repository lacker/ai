#!/usr/bin/env python

nth = 0
numer = 1
denom = 1
more_digs = 0

while True:
  nth += 1
  if nth > 1000:
    break

  prev_numer, prev_denom = numer, denom
  
  numer = prev_numer + 2 * prev_denom
  denom = prev_numer + prev_denom

  print '#%d: %d/%d' % (nth, numer, denom)
  if len(str(numer)) > len(str(denom)):
    more_digs += 1

print more_digs, 'have more digs'
