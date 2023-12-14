import { readFileSync } from "fs";

const getDistance = (time, distance) => {
  let pass = 0;
  for (let i = 0; i < time; i++) {
    const timeLeft = time - i;
    const dist = timeLeft * i;
    if (dist > distance) {
      pass++;
    }
  }
  return pass;
};

const problem1 = () => {
  const input = readFileSync("./input.txt", "utf8")
    .trim()
    .split("\n")
    .map((line) =>
      line
        .split(" ")
        .slice(1)
        .filter((x) => x !== "")
        .map((x) => +x)
    );
  const [time, distance] = input;

  let sum = 1;
  for (let i = 0; i < time.length; i++) {
    const t = time[i];
    const d = distance[i];
    const count = getDistance(t, d);

    sum *= count;
  }

  console.log(sum);
};

problem1();

const problem2 = () => {
  const input = readFileSync("./input.txt", "utf8")
    .trim()
    .split("\n")
    .map(
      (line) =>
        +line
          .split(" ")
          .slice(1)
          .filter((x) => x !== "")
          .join("")
    );
  const [time, distance] = input;

  const count = getDistance(time, distance);

  console.log(count);
};

problem2();
