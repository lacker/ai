#!/usr/bin/env python3
'''
Scores an input that looks like

>>> # python statement
evaled result

So a success might be

>>> 3 + 5
8

and a failure might be

>>> 5 > 3
False

Then prints some stats.
'''

def binary_eval(s):
  part1, part2 = s.split('*')
  a = int(part1, 2)
  b = int(part2, 2)
  c = a * b
  return '{0:b}'.format(c)

# PREFIX = '>>> '
# eval_fn = eval
PREFIX = '>'
eval_fn = binary_eval

def main():
  with open('save/output.txt') as f:
    correct = 0
    incorrect = 0
    last_line = ''
    graded = []
    for line in f.readlines():
      line = line.strip()
      graded.append(line)
      
      if last_line.startswith(PREFIX):
        code = last_line[4:]

        response = ''
        try:
          response = repr(eval_fn(code))
        except:
          response = 'Error'
          
        if line == response:
          correct += 1
          graded.append('Correct')
        else:
          incorrect += 1
          graded.append('Incorrect')

      last_line = line

    print('%d / %d = %.2f%% correct\n' % (
      correct, correct + incorrect,
      100.0 * correct / (correct + incorrect)))

    for line in graded:
      print(line)
          

if __name__ == '__main__':
  main()
