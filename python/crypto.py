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

# Vigenere from 1.21 b
ciphertext = '''
KCCPKEGUFDPHQTYAVINRRTMVGRKDNBVFDETDGILTXRGUD
DKOTFMBPVGEGLTGCKQRACQCWDNAWCRXIZAKFTLEWRPTYC
QKYVXCHKFTPONCQQRHJVAJUWETMCMSPKQDYHJVDAHCTRL
SVSKCGCZQQDZXGSFRLSWCWSJTBHAFSIASPRJAHKJRJUMV
GKMITZHFPDISPZLVLGWTFPLKKEBDPGCEBSHCTJRWXBAFS
PEZQNRWXCVYCGAONWDDKACKAWBBIKFTIOVKCGGHJVLNHI
FFSQESVYCLACNVRWBBIREPBBVFEXOSCDYGZWPFDTKFQIY
CWHJVLNHIQIBTKHJVNPIST
'''.replace('\n', '')

if __name__ == '__main__':
  # Doesn't seem to yield much. Hrmph.
  indextest(ciphertext)

  # This makes me suspect that block size = 6
  blast(ciphertext)
