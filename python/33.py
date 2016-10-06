#!/usr/bin/env python

def naive_cancel(a, b):
  to_cancel = set(str(a)).intersection(set(str(b)))
  if len(to_cancel) != 1:
    return 1, 0
  bad_digit = to_cancel.pop()
  if bad_digit == '0':
    return 1, 0
  c_string = ''.join(digit for digit in str(a) if digit != bad_digit)
  d_string = ''.join(digit for digit in str(b) if digit != bad_digit)
  if len(c_string) != 1 or len(d_string) != 1:
    return 1, 0
  return int(c_string), int(d_string)

# Test a / b
for b in range(11, 100):
  for a in range(10, b):
    c, d = naive_cancel(a, b)
    if b * c == a * d:
      print a, '/', b
