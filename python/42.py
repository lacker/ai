#!/usr/bin/env python

def letterscore(letter):
  return ord(letter) - ord('A') + 1

def wordscore(word):
  return sum(map(letterscore, word))
  
words = map(lambda s: s.strip('"'),
            open('p042_words.txt').read().strip().split(','))

def triangles():
  for n in range(1, 1000):
    yield n * (n + 1) / 2

def is_triangle(x):
  for t in triangles():
    if t == x:
      return True
    if t > x:
      return False
  raise Exception('is_triangle sucks')

count = 0
for word in words:
  if is_triangle(wordscore(word)):
    count += 1
    print word
print count
