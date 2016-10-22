#!/usr/bin/env python

'Return the first n values in the partial fraction thingy for e'
def e_values(n):
  answer = [2, 1, 2]
  next_big_val = 4
  while True:
    answer.append(1)
    answer.append(1)
    answer.append(next_big_val)
    next_big_val += 2
    if len(answer) >= n:
      return answer[:n]

'Returns a (numerator, denominator) tuple'
def make_continued_fraction(values):
  if len(values) == 1:
    return (values[0], 1)

  rest_numer, rest_denom = make_continued_fraction(values[1:])

  numer = values[0] * rest_numer + rest_denom
  denom = rest_numer
  return (numer, denom)

def solve(n):
  print 'solving', n
  numer, denom = make_continued_fraction(e_values(n))
  print numer, '/', denom
  print 'numer digit sum:', sum(map(int, str(numer)))

solve(10)
solve(100)

