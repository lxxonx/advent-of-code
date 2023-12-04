import { readFileSync } from "fs";

const input = readFileSync("./input.txt", "utf8").split("\n");

const problem1 = (input) => {
  const list = input.map((line) => {
    //   console.log(line);
    line = line.replaceAll("  ", " ");
    const [gameInfo, ...games] = line.split(": ");
    const gameNum = +gameInfo.split(" ")[1];

    const [hands, lucky] = games.flatMap((a) => a.trim().split(" | "));

    const handsNum = hands.split(" ").map((a) => +a);
    const luckyNum = lucky.split(" ").map((a) => +a);

    return { gameNum, handsNum, luckyNum };
  });

  let sum = 0;

  for (const l of list) {
    let mult = 0;
    for (const hand of l.handsNum) {
      for (const lucky of l.luckyNum) {
        if (hand === lucky) {
          if (mult === 0) {
            mult = 1;
          } else {
            mult *= 2;
          }
        }
      }
    }
    sum += mult;
  }

  console.log(sum);
};

// problem1(input);

const problem2 = (input) => {
  const list = input.map((line) => {
    line = line.replaceAll("  ", " ");
    const [gameInfo, ...games] = line.split(": ");
    const gameNum = +gameInfo.split(" ")[1];

    const [hands, lucky] = games.flatMap((a) => a.trim().split(" | "));

    const handsNum = hands.split(" ").map((a) => +a);
    const luckyNum = lucky.split(" ").map((a) => +a);

    return { gameNum, handsNum, luckyNum };
  });

  const copies = new Array(list.length + 1).fill(1);

  for (let i = 1; i < list.length + 1; i++) {
    const l = list[i - 1];
    const copy = copies[i];
    let matches = 0;
    for (const hand of l.handsNum) {
      for (const lucky of l.luckyNum) {
        if (hand === lucky) {
          matches++;
        }
      }
    }
    for (let j = 1; j < matches + 1; j++) {
      copies[i + j] += copy;
    }
  }

  console.log(copies.reduce((a, b) => a + b, 0) - 1);
};

problem2(input);
