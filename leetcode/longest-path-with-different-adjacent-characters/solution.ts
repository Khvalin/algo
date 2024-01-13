function longestPath(parent: number[], s: string): number {
  const N = parent.length;
  const children = parent.map(() => []);
  const depth = parent.map(() => -1);

  let root = 0;
  for (let i = 0; i < parent.length; i++) {
    const p = parent[i];
    if (p < 0) {
      root = i;
      continue;
    }
    children[p].push(i);
  }
  let res = 0;

  const walk = (ind) => {
    if (ind >= N || ind < 0) {
      return 0;
    }
    if (depth[ind] >= 0) {
      return depth[ind];
    }
    let r = 1;
    for (const i of children[ind]) {
      const d = walk(i);
      if (s[ind] !== s[i]) {
        r = Math.max(r, 1 + d);
      }
    }

    depth[ind] = r;
    res = Math.max(res, r);
    return r;
  };
  walk(root);

  for (let i = 0; i < N; i++) {
    const a = children[i];
    if (a.length < 2) {
      continue;
    }

    const b = [0, 0, 0];
    for (const j of a) {
      if (s[i] === s[j]) {
        continue;
      }
      b[0] = depth[j];
      if (b[0] > b[1]) {
        [b[0], b[1]] = [b[1], b[0]];
      }

      if (b[1] > b[2]) {
        [b[1], b[2]] = [b[2], b[1]];
      }
      res = Math.max(res, 1 + b[2] + b[1]);
    }
  }

  return res;
}

let res;
/*res = longestPath( [-1,0,0,1,1,2], "abacbe")
console.log(res)

res = longestPath( [-1,0,0,0], "aabc")
console.log(res)
*/
res = longestPath([-1, 0, 1], "aab");
console.log(res);
