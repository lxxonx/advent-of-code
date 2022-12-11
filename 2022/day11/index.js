import fs from "fs";

const gcd = (a, b) => (a ? gcd(b % a, a) : b);

const lcm = (a, b) => (a * b) / gcd(a, b);

class Monkey {
  constructor(monkeyStr) {
    const lines = monkeyStr.split("\n");

    this.div = +lines[3].split("by ")[1];
    this.items = lines[1].split(": ")[1].split(", ").map(Number);
    this.trueMonkey = +lines[4].split("monkey ")[1];
    this.falseMonkey = +lines[5].split("monkey ")[1];

    this.operation = (old, lcm) => {
      old = old % BigInt(lcm);
      const op = lines[2].split("= ")[1];
      const opArr = op.split(" ");
      opArr[2] = isNaN(Number(opArr[2])) ? opArr[2] : opArr[2] + "n";
      const opStr = opArr.join("");
      let e;
      try {
        e = eval(opStr);
      } catch {
        e = old;
      }

      return BigInt(e);
    };

    this.test = (old) => {
      if (old % BigInt(this.div) === BigInt(0)) {
        return [old, this.trueMonkey];
      } else {
        return [old, this.falseMonkey];
      }
    };
  }

  count = 0;

  round(monkeys, lcm) {
    this.count += this.items.length;
    while (this.items.length > 0) {
      const item = this.items.shift();
      const result = this.operation(BigInt(item), lcm);
      const [newItem, nextMonkey] = this.test(result);
      monkeys[nextMonkey].items.push(newItem);
    }
  }
}

// const PART1_MAX_ROUND = 20;

// const part1 = () => {
//   const monkeys = [];
//   const input = fs.readFileSync("input.txt", "utf8");
//   const monkeyStr = input.split("\n\n");

//   for (const monk of monkeyStr) {
//     const monkey = new Monkey(monk);
//     monkeys.push(monkey);
//   }

//   for (let round = 0; round < PART1_MAX_ROUND; round++) {
//     for (const monkey of monkeys) {
//       monkey.round(monkeys);
//     }
//   }

//   monkeys.forEach((monkey, i) => console.log(i, monkey.count));
// };

// part1();

const PART2_MAX_ROUND = 10_000;

const part2 = () => {
  const monkeys = [];
  const input = fs.readFileSync("input.txt", "utf8");
  const monkeyStr = input.split("\n\n");

  for (const monk of monkeyStr) {
    const monkey = new Monkey(monk);
    monkeys.push(monkey);
  }

  const l = monkeys.map((monkey) => monkey.div).reduce(lcm);

  for (let round = 0; round < PART2_MAX_ROUND; round++) {
    for (const monkey of monkeys) {
      monkey.round(monkeys, l);
    }
  }

  const counts = monkeys.map((monkey) => monkey.count).sort((a, b) => b - a);
  console.log(counts[0] * counts[1]);
};

part2();
