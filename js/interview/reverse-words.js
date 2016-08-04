/*
Reverse words in a string in-place
*/

// Reverses a single word in a string, given the first and last indices
// to swap.
function reverseWord(str, first, last) {
  if (first >= last) {
    return;
  }
  let tmp = str[first];
  str[first] = str[last];
  str[last] = tmp;
  reverseWord(str, first + 1, last - 1);
}

// Reverse all words, splitting on ' '
function reverseAllWords(str) {
  let lastSpace = -1;
  let checkForSpace = 0;
  while (checkForSpace <= str.length) {
    // The index str.length counts as a space
    if (checkForSpace == str.length || str[checkForSpace] == ' ') {
      reverseWord(str, lastSpace + 1, checkForSpace - 1);
    }
  }
}
