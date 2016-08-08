/*
You are given an array with integers between 1 and 1,000,000. One integer is missing. How can you determine which one?
*/

function missingFrom(arr) {
  let n = 1000000;
  let target = n * (n + 1) / 2;

  let sum = 0;
  for (let element of arr) {
    sum += element;
  }

  return target - sum;
}
