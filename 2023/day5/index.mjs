import { readFileSync } from "fs";

const input = readFileSync("./input.txt", "utf8").trim();

const [seeds, ...maps] = input.split("\n\n");

const seedsList = seeds.split(": ")[1].split(" ").map(Number);

const mapsList = maps.map((map) =>
  map
    .split("\n")
    .slice(1)
    .map((row) => row.split(" ").map(Number))
);

const problem1 = (seeds, maps) => {
  let result = Infinity;

  for (let seed of seeds) {
    for (const map of maps) {
      for (const row of map) {
        const [destination, source, range] = row;
        if (seed < source + range && seed >= source) {
          const value = destination - source;
          seed += value;
          break;
        }
      }
    }
    result = Math.min(result, seed);
  }

  console.log(result);
};

// problem1(seedsList, mapsList);

const problem2 = (seeds, maps) => {
  let result = Infinity;

  const seedRange = seeds.reduce((acc, cur, idx) => {
    if (idx % 2 === 0) {
      acc.push(cur);
    } else {
      acc[acc.length - 1] = [acc[acc.length - 1], cur];
    }

    return acc;
  }, []);

  for (const [start, end] of seedRange) {
    for (let i = start; i < start + end; i++) {
      let seed = i;
      for (const map of maps) {
        for (const row of map) {
          const [destination, source, range] = row;
          if (i < source + range && i >= source) {
            const value = destination - source;
            seed += value;
            break;
          }
        }
      }
      result = Math.min(result, seed);
    }
  }
  console.log(result);
};

problem2(seedsList, mapsList);
