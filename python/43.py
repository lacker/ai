#!/usr/bin/env python

from itertools import permutations

answer = 0
for perm in permutations('0123456789'):
  s = ''.join(perm)
  if s.startswith('0'):
    continue
  if int(s[1:4]) % 2 != 0: continue
  if int(s[2:5]) % 3 != 0: continue
  if int(s[3:6]) % 5 != 0: continue
  if int(s[4:7]) % 7 != 0: continue
  if int(s[5:8]) % 11 != 0: continue
  if int(s[6:9]) % 13 != 0: continue
  if int(s[7:10]) % 17 != 0: continue
  print int(s)
  answer += int(s)

print answer
