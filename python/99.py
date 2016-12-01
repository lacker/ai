#!/usr/bin/env python

from math import log

best_line_num = 0
best_val = 0

for i, line in enumerate(open('p099_base_exp.txt')):
  line_num = i + 1
  a, b = map(int, line.strip().split(','))
  val = b * log(a)
  if val > best_val:
    best_line_num, best_val = line_num, val
    print line_num, a, b
