/**
 * @param {number[]} A
 * @return {number}
 */
var largestComponentSize = function(A) {
  const primes = [];

  const n = Math.max(...A) >> 1;
  const isPrime = new Map();

  const buckets = new Map();
  const sets = new Map();

  for (let i = 2; i <= n + 1; i += 2) {
    if (!isPrime.has(i)) {
      primes.push(i);

      sets.set(i, [ i ]);
      buckets.set(i, i);

      for (let j = 3; i * j <= n; j += 2) {
        isPrime.set(i * j, false);
      }
    }

    if (i == 2) {
      i = 1; //haha
    }
  }

  let res = 0;

  const counts = new Map();
  for (const n of A) {
    let targetId = buckets.get(n) || -1;
    let s = sets.get(targetId) || [];

    for (let i = 0; primes[i] <= n >> 1; i++) {
      let p = primes[i];
      if (n % p === 0) {
        const bIndex = buckets.get(p);

        if (targetId === bIndex) {
          continue;
        }

        s.push(...sets.get(bIndex));
        sets.set(p, []);
        if (targetId === -1) {
          targetId = bIndex;
        }
      }
    }

    if (targetId === -1) {
      continue;
    }

    let c = (counts.get(targetId) || 0) + 1;
    for (const j of s) {
      if (j !== targetId) {
        c += counts.get(j) || 0;
        counts.set(j, 0);
      }
      buckets.set(j, targetId);
    }

    sets.set(targetId, s);
    counts.set(targetId, c);

    res = Math.max(res, c);
  }

  return res;
};

/* TESTS */

// prettier-ignore
let a = [4,6,15,35]
// console.log(4, largestComponentSize(a));

// a = [ 65, 35, 43, 76, 15, 11, 81, 22, 55, 92, 31 ];
// console.log(9, largestComponentSize(a));

// a = [ 2, 3, 6, 7, 4, 12, 21, 39 ];
// console.log(8, largestComponentSize(a));

a = [ 1, 2, 3, 4, 5, 6, 7, 8, 9 ];
console.log(6, largestComponentSize(a));

// prettier-ignore
a = [512,129,939,138,907,652,781,398,399,16,857,239,185,161,290,547,932,945,299,429,559,432,49,950,953,319,864,325,993,200,460,461,463,722,853,313,729,348,608,225,871,489,750,751,371,884,376,506,737]
console.log(40, largestComponentSize(a));
