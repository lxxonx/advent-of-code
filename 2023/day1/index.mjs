import { readFileSync } from "fs";

const input = readFileSync("./input.txt", "utf8").split("\n");

const problem1 = (input) => {
  let sum = 0;
  for (const line of input) {
    const nums = line.replace(/[A-Za-z]/g, "");

    const first = parseInt(nums[0]);

    const last = parseInt(nums[nums.length - 1]);

    const num = +[first, last].join("");

    sum += num;
  }

  console.log(sum);
};

const problem2 = (input) => {
  const validNumbers = [
    "one",
    "two",
    "three",
    "four",
    "five",
    "six",
    "seven",
    "eight",
    "nine",
    "1",
    "2",
    "3",
    "4",
    "5",
    "6",
    "7",
    "8",
    "9",
  ];

  const numStringMap = {
    one: "1",
    two: "2",
    three: "3",
    four: "4",
    five: "5",
    six: "6",
    seven: "7",
    eight: "8",
    nine: "9",
  };

  let sum = 0;
  for (const line of input) {
    const firstIndexMap = {};
    const lastIndexMap = {};

    validNumbers.forEach((n) => {
      firstIndexMap[n] = Infinity;
    });
    validNumbers.forEach((n) => {
      lastIndexMap[n] = -1;
    });

    for (const key of validNumbers) {
      const firstIndex = line.indexOf(key);
      const lastIndex = line.lastIndexOf(key);

      if (firstIndex !== -1) {
        firstIndexMap[key] = firstIndex;
      }
      lastIndexMap[key] = lastIndex;
    }

    const firstNum = +Object.entries(firstIndexMap)
      .filter(([, v]) => v !== Infinity)
      .sort(([, a], [, b]) => a - b)
      .map(([k, v]) => numStringMap[k] || k)[0];

    const lastNum = +Object.entries(lastIndexMap)
      .filter(([, v]) => v !== -1)
      .sort(([, a], [, b]) => b - a)
      .map(([k, v]) => numStringMap[k] || k)[0];

    const num = +[firstNum, lastNum].join("");

    sum += num;
  }

  console.log(sum);
};

// problem1(input);
problem2(input);
