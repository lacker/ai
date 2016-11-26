#!/usr/bin/env python

from itertools import combinations, permutations

def pair_results(x, y):
  answer = [x + y,
            x - y,
            y - x,
            x * y]
  if y != 0:
    answer.append(float(x) / y)
  if x != 0:
    answer.append(float(y) / x)
  return answer

def perm_results(perm):
  assert len(perm) >= 2
  if len(perm) == 2:
    return pair_results(*perm)
  a, b, rest = perm[0], perm[1], perm[2:]
  answer = []
  for pair_result in pair_results(a, b):
    for result in perm_results([pair_result] + rest):
      answer.append(result)
  return answer

def score(digits):
  results = set()
  for perm in permutations(digits):
    for result in perm_results(list(perm)):
      int_result = int(result)
      if int_result != result:
        continue
      results.add(int_result)

  score = 0
  while score + 1 in results:
    score += 1
  return score

  
best_score = 0
for digits in combinations([0, 1, 2, 3, 4, 5, 6, 7, 8, 9], 4):
  s = score(digits)
  if s > best_score:
    best_score = s
    print s, digits
