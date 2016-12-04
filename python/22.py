#!/usr/bin/env python

names = [x.strip('"') for x in sorted(open('p022_names.txt').read().split(','))]

def letterscore(ch):
  return ord(ch) - ord('A') + 1

def wordscore(word):
  return sum(letterscore(ch) for ch in word)

answer = 0
for i, name in enumerate(names):
  score = (i + 1) * wordscore(name)
  answer += score
print answer
    

