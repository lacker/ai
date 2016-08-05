// Count the number of set bits in a number

function simple(number) {
  if (number < 0) {
    throw new Error('number shouldnt be negative');
  }
  let answer = 0;
  while (number > 0) {
    if (number & 1 == 1) {
      answer += 1;
    }
    number >>= 1;
  }
  return number;
}
