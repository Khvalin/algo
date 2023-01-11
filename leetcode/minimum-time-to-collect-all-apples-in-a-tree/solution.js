/**
 * @param {number} n
 * @param {number[][]} edges
 * @param {boolean[]} hasApple
 * @return {number}
 */
var minTime = function (n, edges, hasApple) {
  const visited = new Set();
  const counts = new Array(n);
  const parents = (new Array(n)).map(() => -1);
  const adj = new Array(n);
  for (const [a, b] of edges) {
    adj[a] = adj[a] || [];
    adj[b] = adj[b] || [];
    adj[a].push(b);
    adj[b].push(a);

    counts[a] = -1;
  }

  hasApple[0] = false; //collected by default
  const root = 0;
  let total = 0;

  const count = (ind) => {
    visited.add(ind);
    if (counts[ind] >= 0) {
      return counts[ind];
    }
    let res = 0;
    if (hasApple[ind]) {
      res++;
      total++;
    }

    for (const i of adj[ind]) {
      if (!visited.has(i)) {
        parents[i] = ind;
        res += count(i);
      }
    }
    counts[ind] = res;

    return res;
  };
  count(root);

  visited.clear();

  let node = root;
  let res = 0;
  outer:
  while (total > 0 || node !== root) {
    visited.add(node);
    res++;
    if (hasApple[node]) {
      total--;
      hasApple[node] = false;
      counts[node]--;
    }

    const a = adj[node];
    for (let i of a) {
      if (counts[i] > 0 && !visited.has(i)) {
        node = i;
        continue outer;
      }
    }
    if (parents[node] > -1) {
      node = parents[node];
    }
  }

  return res;
};

let n = 7,
  edges = [[0, 1], [0, 2], [1, 4], [1, 5], [2, 3], [2, 6]],
  hasApple = [false, false, true, false, true, true, false];
n = 4;
edges = [[0, 1], [1, 2], [0, 3]];
hasApple = [true, true, true, true];

n = 7;
edges = [[0, 1], [1, 2], [1, 3], [2, 4], [4, 5], [4, 6]];
hasApple = [false, true, false, false, true, false, true];

n = 7;
edges = [[0, 1], [0, 2], [1, 4], [1, 5], [2, 3], [2, 6]];
hasApple = [true, false, false, false, false, false, false];

let res = minTime(n, edges, hasApple);
console.log(res);
