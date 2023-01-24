import fs from "fs";

const i = fs.readFileSync("input.txt", "utf8");

const compareArray = (a, b) => {
  for (let i = 0; i < Math.min(a.length, b.length); i++) {
    if (Number.isInteger(a[i]) && Number.isInteger(b[i])) {
      if (a[i] !== b[i]) {
        return a[i] - b[i];
      }
    } else {
      let aArr = Number.isInteger(a[i]) ? [a[i]] : a[i];
      let bArr = Number.isInteger(b[i]) ? [b[i]] : b[i];
      const cmpNext = compareArray(aArr, bArr);
      if (cmpNext !== 0) {
        return cmpNext;
      }
    }
  }

  return a.length - b.length;
};

const solution = (input) => {
  let count = 0;
  const sets = input.split("\n\n");

  let arr = [[[2]], [[6]]];

  for (const index in sets) {
    const pair = sets[index].split("\n");

    const left = JSON.parse(pair[0]);
    const right = JSON.parse(pair[1]);

    arr.push(left);
    arr.push(right);

    if (compareArray(left, right) < 0) {
      count += Number(index) + 1;
    }
  }

  // part 1
  console.log(count);

  // part 2
  arr = arr.sort(compareArray);

  const d1 =
    arr.findIndex((i) => JSON.stringify(i) === JSON.stringify([[2]])) + 1;
  const d2 =
    arr.findIndex((i) => JSON.stringify(i) === JSON.stringify([[6]])) + 1;

  console.log(d1 * d2);
};

solution(i);
