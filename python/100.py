#!/usr/bin/env python

from fractions import Fraction

'''
15 blue, 6 red:
P(BB) = (15/21)*(14/20) = 1/2

b / (b + r) * (b - 1) / (b + r - 1) = 1 / 2

say b + r = t

2 * b * (b - 1) == t * (t - 1)

so b/t is approx sqrt(1/2)

found:
1 red, 1 total
3 red, 4 total
15 red, 21 total
85 red, 120 total
493 red, 697 total
2871 red, 4060 total
16731 red, 23661 total
97513 red, 137904 total
568345 red, 803761 total
3312555 red, 4684660 total

but needs more scaling
'''

'''
Continued fraction of sqrt(1/2) with n 2's
'''
def frac(n):
  if n == 0:
    return Fraction(1, 1)
  return 1 / (1 + (1 / (1 + (1 / frac(n - 1)))))

a, b = 1, 1
for x in range(40):
  f = frac(x)
  if f * f < 0.5:
    a += f.numerator
    b += f.denominator
    print a, b
    if b > 10 ** 12:
      break
