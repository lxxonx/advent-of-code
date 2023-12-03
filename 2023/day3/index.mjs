import { readFileSync } from "fs";

const input = readFileSync("./input.txt", "utf8")
  .split("\n")
  .map((m) => m.split(""));

const problem1 = (input) => {
  const findNumberEnd = (i, line) => {
    let value = line[i];
    while (Number.isInteger(+line[i + 1])) {
      value += line[i + 1];
      i++;
    }
    return [i, value];
  };

  const isAdjacent = (symbolCoords, numberCoords) => {
    const [symbolLine, symbolIndex] = symbolCoords;
    const [numberLine, numberStart, numberEnd] = numberCoords;
    if (
      symbolLine <= numberLine + 1 &&
      symbolLine >= numberLine - 1 &&
      symbolIndex <= numberEnd + 1 &&
      symbolIndex >= numberStart - 1
    ) {
      return true;
    }
  };
  const symbols = [];
  const numbers = [];
  for (let lineNum = 0; lineNum < input.length; lineNum++) {
    const line = input[lineNum];
    for (let i = 0; i < line.length; i++) {
      if (line[i] === ".") {
        continue;
      }
      if (!Number.isInteger(+line[i])) {
        symbols.push([lineNum, i]);
        continue;
      }
      if (Number.isInteger(+line[i])) {
        const [end, value] = findNumberEnd(i, line);
        numbers.push({ line: lineNum, start: i, end, value: +value });
        i = end;
      }
    }
  }
  let sum = 0;
  for (const symbol of symbols) {
    for (const number of numbers) {
      if (isAdjacent(symbol, [number.line, number.start, number.end])) {
        sum += number.value;
      }
    }
  }
  console.log(sum);
};

// problem1(input);

const problem2 = (input) => {
  const findNumberEnd = (i, line) => {
    let value = line[i];
    while (Number.isInteger(+line[i + 1])) {
      value += line[i + 1];
      i++;
    }
    return [i, value];
  };

  const isAdjacent = (symbolCoords, numberCoords) => {
    const [symbolLine, symbolIndex] = symbolCoords;
    const [numberLine, numberStart, numberEnd] = numberCoords;
    if (
      symbolLine <= numberLine + 1 &&
      symbolLine >= numberLine - 1 &&
      symbolIndex <= numberEnd + 1 &&
      symbolIndex >= numberStart - 1
    ) {
      return true;
    }
  };
  const symbols = [];
  const numbers = [];
  for (let lineNum = 0; lineNum < input.length; lineNum++) {
    const line = input[lineNum];
    for (let i = 0; i < line.length; i++) {
      if (line[i] === "*") {
        symbols.push([lineNum, i]);
        continue;
      }
      if (Number.isInteger(+line[i])) {
        const [end, value] = findNumberEnd(i, line);
        numbers.push({ line: lineNum, start: i, end, value: +value });
        i = end;
      }
    }
  }
  let sum = 0;
  for (const symbol of symbols) {
    const multiplier = [];
    for (const number of numbers) {
      if (isAdjacent(symbol, [number.line, number.start, number.end])) {
        multiplier.push(number.value);
      }
      if (multiplier.length === 3) {
        break;
      }
    }
    if (multiplier.length === 2) {
      sum += multiplier.reduce((a, b) => a * b, 1);
    }
  }
  console.log(sum);
};

problem2(input);
