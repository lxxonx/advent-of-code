import { readFileSync } from "fs";

const input = readFileSync("./input.txt", "utf8").split("\n");

const problem1 = (input) => {
  const bag = {
    red: 12,
    green: 13,
    blue: 14,
  };

  let sum = 0;

  for (const game of input) {
    const [num, ...rounds] = game.split(": ");

    const roundNum = +num.split(" ")[1];

    const parsed = rounds
      .map((m) => m.split("; "))
      .flatMap((m) => m.map((n) => n.split(", ")));

    let pass = true;

    for (const round of parsed) {
      round.forEach((m) => {
        const [count, color] = m.split(" ");
        if (bag[color] < +count) {
          pass = false;
        }
      });
    }

    if (pass) {
      sum += roundNum;
    }
  }

  console.log(sum);
};

// problem1(input);

const problem2 = (input) => {
  let sum = 0;

  for (const game of input) {
    let red = 0;
    let green = 0;
    let blue = 0;

    const rounds = game
      .split(": ")
      .slice(1)
      .flatMap((m) => m.split("; ").map((n) => n.split(", ")));

    for (const round of rounds) {
      round.forEach((m) => {
        const [count, color] = m.split(" ");
        if (color === "red") {
          red = Math.max(red, +count);
        } else if (color === "green") {
          green = Math.max(green, +count);
        } else if (color === "blue") {
          blue = Math.max(blue, +count);
        }
      });
    }

    const multiplier = red * green * blue;
    sum += multiplier;
  }

  console.log(sum);
};

problem2(input);
