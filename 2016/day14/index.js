import os from "os";

import { spawn, Thread, Worker } from "threads";

console.time("ended");

async function part1(workers, doorId, chunk) {
  const concurrency = workers.length;
  console.log(concurrency);

  const hashTable = [];
  let found = 0;
  for (let start = 0; ; start += chunk * concurrency) {
    const jobs = workers.map((worker, i) =>
      worker.hash(doorId, start + chunk * i, start + chunk * (i + 1), hashTable)
    );
    for await (const r of jobs)
      for (const [i, h, j, h2] of r) {
        found++;
        console.log(i, h, j, h2);
        if (found === 64) {
          return i;
        }
      }
  }
}

async function part2(workers, doorId, chunk) {
  const concurrency = workers.length;
  const hashTable = [];
  let found = 0;
  for (let start = 0; ; start += chunk * concurrency) {
    const jobs = workers.map((worker, i) =>
      worker.stretchedHash(
        doorId,
        start + chunk * i,
        start + chunk * (i + 1),
        hashTable
      )
    );
    for await (const r of jobs)
      for (const [i, h, j, h2] of r) {
        found++;
        console.log(i, h, j, h2);
        if (found === 64) {
          return i;
        }
      }
  }
}

const workers = await Promise.all(
  os.cpus().map((_) => spawn(new Worker("./scan")))
);
const chunk = 3000;
const doorId = "jlmsuwbz";
console.log(await part2(workers, doorId, chunk));
await Promise.all(workers.map(Thread.terminate));
console.timeEnd("ended");
