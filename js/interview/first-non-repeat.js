/*
Find the first non-repeating letter in a string
*/

function firstNonRepeat(str) {
  let letterCount = new Map(str);
  for (let letter of str) {
    letterCount[letter] = (letterCount[letter] || 0) + 1;
  }
  for (let letter of str) {
    if (letterCount[letter] == 1) {
      return letter;
    }
  }
  throw new Error('no non repeats');
}
