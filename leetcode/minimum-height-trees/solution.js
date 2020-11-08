const findMinHeightTrees = (n, edges) => {
  let res = [];

  const adj = [];
  const degree = [];

  for (let i = 0; i < n; i++) {
    res.push(i);
    adj.push([]);
    degree.push(0);
  }

  if (!edges.length) {
    return res;
  }

  for (const [ a, b ] of edges) {
    degree[a]++;
    degree[b]++;

    adj[a].push(b);
    adj[b].push(a);
  }

  let leaves = [];
  for (let i = 0; i < n; i++) {
    if (degree[i] === 1) {
      leaves.push(i);
    }
  }

  while (leaves.length) {
    res = [];
    const len = leaves.length;
    const next = new Set();

    for (let i = 0; i < len; i++) {
      const l = leaves[i];
      res.push(l);
      degree[l]--;

      for (const e of adj[l]) {
        degree[e]--;
        if (degree[e] === 1 && !next.has(e)) {
          next.add(e);
          leaves.push(e);
        }
      }
    }

    leaves.splice(0, len);
  }

  return res;
};

console.log(findMinHeightTrees(1, []));
console.log(findMinHeightTrees(4, [ [ 1, 0 ], [ 1, 2 ], [ 1, 3 ] ]));
console.log(findMinHeightTrees(6, [ [ 3, 0 ], [ 3, 1 ], [ 3, 2 ], [ 3, 4 ], [ 5, 4 ] ]));

console.log(findMinHeightTrees(7, [ [ 0, 1 ], [ 1, 2 ], [ 1, 3 ], [ 2, 4 ], [ 3, 5 ], [ 4, 6 ] ]));
