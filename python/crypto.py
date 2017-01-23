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

def modbuckets(text, modulus):
  '''
  Makes equivalence classes for the given modulus, each class is a list/bucket
  Returns a list of these buckets
  '''
  output = [[] for _ in range(modulus)]
  for i, ch in enumerate(text):
    index = i % modulus
    output[index].append(ch)
  return output

def indextest(text):
  '''
  Tries out the indices of coincidence for blocks of different lengths.
  '''
  for m in range(1, 11):
    buckets = modbuckets(text, m)
    indices = map(ioc, buckets)
    print 'block size %d:' % m
    print indices
