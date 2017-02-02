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
    print 'average:', float(sum(indices)) / len(indices)
    print indices

def blast(text):
  tcount = {}
  for i in range(len(text)):
    triple = text[i:i+3]
    if len(triple) < 3:
      break
    if triple in tcount:
      print '%d = %d - %d' % (i - tcount[triple], i, tcount[triple])
    tcount[triple] = i

ENGLISH = {    
  'A': .082,
  'B': .015,
  'C': .028,
  'D': .043,
  'E': .127,
  'F': .022,
  'G': .020,
  'H': .061,
  'I': .070,
  'J': .002,
  'K': .008,
  'L': .040,
  'M': .024,
  'N': .067,
  'O': .075,
  'P': .019,
  'Q': .001,
  'R': .060,
  'S': .063,
  'T': .091,
  'U': .028,
  'V': .010,
  'W': .023,
  'X': .001,
  'Y': .020,
  'Z': .001,
}

def dot(a, b):
  '''
  dot-products two distribution-vectors encoded as dicts.
  '''
  answer = 0
  for key, value1 in a.items():
    value2 = b.get(key, 0)
    answer += value1 * value2
  return answer

def s2vec(s):
  answer = {}
  for ch in s:
    if ch in answer:
      answer[ch] += 1
    else:
      answer[ch] = 1
  return answer
  
def ch2num(ch):
  return ord(ch) - ord('A')

def num2ch(num):
  return chr(ord('A') + num)

def chrotate(ch, n):
  return num2ch((ch2num(ch) + n) % 26)

def srotate(s, n):
  return ''.join(chrotate(ch, n) for ch in s)
  
def rotate(vector, n):
  '''
  vector is a dict. n is like, if n = 2 then A -> C.
  '''
  answer = {}
  for key, value in vector.items():
    answer[chrotate(key, n)] = value
  return answer

def make_count_vector(text):
  answer = {}
  for ch in text:
    if ch in answer:
      answer[ch] += 1
    else:
      answer[ch] = 1
  return answer
  
def best_rot_score(vector):
  best_rot = -1
  best_score = 0
  for rot in range(26):
    score = dot(ENGLISH, rotate(vector, rot))
    if score > best_score:
      best_rot = rot
      best_score = score
  return best_rot, best_score

def vcrack(ciphertext, m):
  '''
  Prints out crack info
  m is the keyword length
  '''
  plainslices = []
  for i in range(m):
    modslice = ciphertext[i::m]
    cvector = make_count_vector(modslice)
    rot, score = best_rot_score(cvector)
    print 'rot:', rot, 'score:', score
    plainslice = srotate(modslice, rot)
    print 'ps:', plainslice
    plainslices.append(plainslice)
  for block in zip(*plainslices):
    print ''.join(block)

def chaffine(ch, a, b):
  '''
  Does an ax+b on a character
  '''
  return num2ch((ch2num(ch) * a + b) % 26)

def saffine(s, a, b):
  '''
  Does an ax+b on a string
  '''
  return ''.join(chaffine(ch, a, b) for ch in s)

def brutecrack(s1, s2):
  '''
  Prints all (a, b) where saffine takes s1 -> s2
  '''
  for a in range(1, 26):
    for b in range(0, 26):
      if saffine(s1, a, b) == s2:
        print (a, b)

def gcd(a, b):
  if b > a:
    a, b = b, a
  if b == 0:
    return a
  if b == 1:
    return 1
  if a == b:
    return a
  return gcd(b, a % b)
        
def best_affine(s):
  '''
  Finds the affine transformation of s that looks the most English.
  '''
  best_score = -1
  best_transform = None
  for a in range(1, 26):
    if gcd(a, 26) != 1:
      continue
    for b in range(0, 26):
      transform = saffine(s, a, b)
      score = dot(ENGLISH, make_count_vector(transform))
      if score > best_score:
        best_score = score
        best_transform = transform
  return best_transform
        
        
if __name__ == '__main__':
  print 'TODO: solving 1.21 c'
