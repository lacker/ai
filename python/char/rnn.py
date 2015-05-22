#!/usr/bin/python
"""
Character-level recurrent neural networks, using Keras.
See:
https://github.com/karpathy/char-rnn
for inspiration.
"""

import os

chatfile = os.path.realpath(__file__ + "/../data/chat.txt")
print "chatfile:", chatfile
