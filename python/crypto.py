#!/usr/bin/env python

def ioc(text):
  'Finds the index of coincidence of a text.'
  total = 0
  counter = {}
  for ch in text:
    total += 1
    if ch in counter:
      counter[ch] += 1
    else:
      counter[ch] = 1

  if not total:
    raise Exception('cannot get ioc of empty list')
    
  answer = 0.0
  for ch, count in counter.items():
    frac = float(count) / total
    answer += frac * frac

  return answer
