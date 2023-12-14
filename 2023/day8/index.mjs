import { readFileSync } from "fs";

let [instructions, input] = readFileSync("./input.txt", "utf-8").split("\n\n");
instructions = instructions.split("");
input = input.split("\n").reduce((acc, line) => {
  const removeParenthesis = line.replace("(", "").replace(")", "");
  let [key, value] = removeParenthesis.split(" = ");
  value = value.split(", ");
  if (!acc[key]) {
    acc[key] = value;
  } else {
    acc[key].push(value);
  }
  return acc;
}, {});

const problem1 = (instructions, input) => {
  const getNext = (instructions, current) => {
    for (const instruction of instructions) {
      if (instruction === "R") {
        current = input[current][1];
      } else {
        current = input[current][0];
      }
    }

    return current;
  };

  let count = 0;

  let start = "AAA";

  while (start !== "ZZZ") {
    start = getNext(instructions, start);
    count += instructions.length;
  }

  console.log(count);
};

problem1(instructions, input);

const problem2 = (instructions, input) => {
  const getNext = (instructions, current) => {
    for (const instruction of instructions) {
      if (instruction === "R") {
        current = input[current][1];
      } else {
        current = input[current][0];
      }
    }

    return current;
  };

  const startingNodes = Object.keys(input).filter((v) => v.endsWith("A"));

  const counts = [];
  for (const start of startingNodes) {
    let count = 0;
    let current = start;
    while (!current.endsWith("Z")) {
      current = getNext(instructions, current);
      count += instructions.length;
    }
    counts.push(count);
  }

  const gcd = (a, b) => {
    if (b === 0) return a;
    return gcd(b, a % b);
  };

  const lcm = (a, b) => {
    return (a * b) / gcd(a, b);
  };

  console.log(counts.reduce(lcm));
};

problem2(instructions, input);
