import md5 from "md5";
import { expose } from "threads/worker";

const hasConsecutive = (hash, length) => {
  for (let i = 0; i <= hash.length - length; i++) {
    if (
      hash
        .slice(i, i + length)
        .split("")
        .every((x) => x === hash[i])
    ) {
      return hash[i];
    }
  }
  return null;
};

expose({
  hash(doorId, lb, ub, table) {
    let h = "";
    const hashes = [];
    for (let i = lb; i < ub; i++) {
      const id = `${doorId}${i}`;
      if (!table[id]) {
        table[id] = md5(id);
      }
      h = table[id];
      const c = hasConsecutive(h, 3);
      if (!c) continue;
      for (let j = 1; j <= 1000; j++) {
        const tid = `${doorId}${i + j}`;
        if (!table[tid]) {
          table[tid] = md5(tid);
        }
        if (table[tid].includes(c.repeat(5))) {
          hashes.push([i, h, i + j, table[tid]]);
          break;
        }
      }
    }
    return hashes;
  },

  stretchedHash(doorId, lb, ub, table) {
    const hashes = [];
    for (let i = lb; i < ub; i++) {
      const id = `${doorId}${i}`;
      if (!table[id]) {
        let h = id;
        for (let j = 0; j < 2017; j++) {
          h = md5(h);
        }
        table[id] = h;
      }
      const c = hasConsecutive(table[id], 3);
      if (!c) continue;
      for (let j = 1; j <= 1000; j++) {
        const tid = `${doorId}${i + j}`;
        if (!table[tid]) {
          let th = tid;
          for (let j = 0; j < 2017; j++) {
            th = md5(th);
          }
          table[tid] = th;
        }
        if (table[tid].includes(c.repeat(5))) {
          hashes.push([i, table[id], i + j, table[tid]]);
          break;
        }
      }
    }
    return hashes;
  },
});
