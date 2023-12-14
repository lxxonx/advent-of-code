import { readFileSync } from "fs";

const input = readFileSync("./input.txt", "utf-8")
  .trim()
  .split("\n")
  .map((line) => line.split(""));

const pushToTop = (input, i, j) => {
  let dest = i - 1;
  while (true) {
    if (dest < 0) {
      break;
    }
    if (input[dest][j] === "#") {
      break;
    }
    if (input[dest][j] === "O") {
      break;
    }
    input[dest + 1][j] = ".";
    input[dest][j] = "O";
    dest = dest - 1;
  }
  return input;
};

const pushToLeft = (input, i, j) => {
  let dest = j - 1;
  while (true) {
    if (dest < 0) {
      break;
    }
    if (input[i][dest] === "#") {
      break;
    }
    if (input[i][dest] === "O") {
      break;
    }
    input[i][dest + 1] = ".";
    input[i][dest] = "O";
    dest = dest - 1;
  }

  return input;
};

const pushToBottom = (input, i, j) => {
  let dest = i + 1;
  while (true) {
    if (dest > input.length - 1) {
      break;
    }
    if (input[dest][j] === "#") {
      break;
    }
    if (input[dest][j] === "O") {
      break;
    }
    input[dest - 1][j] = ".";
    input[dest][j] = "O";
    dest = dest + 1;
  }
  return input;
};

const pushToRight = (input, i, j) => {
  let dest = j + 1;
  while (true) {
    if (dest > input[0].length - 1) {
      break;
    }
    if (input[i][dest] === "#") {
      break;
    }
    if (input[i][dest] === "O") {
      break;
    }
    input[i][dest - 1] = ".";
    input[i][dest] = "O";
    dest = dest + 1;
  }

  return input;
};

const problem1 = (input) => {
  const tilt = (input) => {
    for (let i = 0; i < input.length; i++) {
      for (let j = 0; j < input[0].length; j++) {
        if (input[i][j] === "O" && i > 0) {
          input = pushToTop(input, i, j);
        }
      }
    }

    return input;
  };
  const tilted = tilt(input);

  let sum = 0;
  for (let i = 0; i < tilted.length; i++) {
    const countO = tilted[i].filter((char) => char === "O").length;

    sum += countO * (tilted.length - i);
    // console.log(countO, tilted.length - i, countO * (tilted.length - i));
  }
  console.log(sum);
};

problem1(input);

const problem2 = (input) => {
  const tiltTop = (input) => {
    for (let i = 0; i < input.length; i++) {
      for (let j = 0; j < input[0].length; j++) {
        if (input[i][j] === "O") {
          input = pushToTop(input, i, j);
        }
      }
    }

    return input;
  };
  const tiltLeft = (input) => {
    for (let i = 0; i < input.length; i++) {
      for (let j = 0; j < input[0].length; j++) {
        if (input[i][j] === "O") {
          input = pushToLeft(input, i, j);
        }
      }
    }

    return input;
  };
  const tiltBottom = (input) => {
    for (let j = 0; j < input[0].length; j++) {
      for (let i = input.length - 1; i >= 0; i--) {
        if (input[i][j] === "O") {
          input = pushToBottom(input, i, j);
        }
      }
    }

    return input;
  };
  const tiltRight = (input) => {
    for (let i = 0; i < input.length; i++) {
      for (let j = input[0].length - 1; j >= 0; j--) {
        if (input[i][j] === "O") {
          input = pushToRight(input, i, j);
        }
      }
    }

    return input;
  };

  const compareMatrix = (matrix, matrix2) => {
    for (let i = 0; i < matrix.length; i++) {
      const row = matrix[i];
      const row2 = matrix2[i];

      if (row.join("") !== row2.join("")) {
        return false;
      }
    }
    return true;
  };

  const hasItInArray = (array, matrix) => {
    for (let i = 0; i < array.length; i++) {
      const row = array[i];
      if (compareMatrix(row, matrix)) {
        return true;
      }
    }
    return false;
  };

  const matrixToString = (matrix) => {
    return matrix.map((row) => row.join("")).join("\n");
  };

  const history = [matrixToString(input)];
  let indexStart = 0,
    indexEnd = 0;

  for (let i = 0; i < 1000000000; i++) {
    let tilted = tiltTop(input);
    tilted = tiltLeft(input);
    tilted = tiltBottom(input);
    tilted = tiltRight(input);

    const copyStr = matrixToString(tilted);
    if (!history.includes(copyStr)) {
      history.push(copyStr);
    } else {
      const idx = history.indexOf(copyStr);
      indexStart = idx;
      indexEnd = i;
      break;
    }
    input = tilted;
  }

  let count = 1000000000;

  while (indexEnd - indexStart < count) {
    count -= indexEnd - indexStart;
  }

  input = history[indexStart + count - 1]
    .split("\n")
    .map((row) => row.split(""));

  let sum = 0;
  for (let i = 0; i < input.length; i++) {
    const countO = input[i].filter((char) => char === "O").length;

    sum += countO * (input.length - i);
  }
  console.log(sum);
};

problem2(input);
