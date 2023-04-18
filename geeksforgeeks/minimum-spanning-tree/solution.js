//User function Template for javascript
/**
 * @param {number[][]} arr
 * @param {number} v
 * @param {number} e
 * @returns {number}
 */
class Solution {
  spanningTree(arr, V, e) {
    const sets = [];
    const parent = [];
    for (let i = 0; i < V; i++) {
      sets.push(new Set([i]));
      parent.push(i);
    }

    arr = arr.map((a) => [1 * a[0], 1 * a[1], 1 * a[2]]);
    arr.sort((a, b) => a[2] - b[2]);

    let sum = 0;
    for (let [u, v, w] of arr) {
      u = parent[u];
      v = parent[v];
      if (u === v) {
        continue;
      }

      sum += w;

      let p = u;
      if (sets[p].size < sets[v].size) {
        p = v;
      }

      if (sets[p].length >= V) {
        break;
      }

      if (u !== p) {
        for (const node of sets[u]) {
          sets[p].add(node);
          parent[node] = p;
        }
      }

      if (v !== p) {
        for (const node of sets[v]) {
          sets[p].add(node);
          parent[node] = p;
        }
      }
    }

    return sum;
  }
}

/*
let v = 7
let e = 7
let arr = [[0, 1, 3],
[1, 3, 3],
[1, 5, 10],
[2, 4, 6],
[2, 6, 9],
[3, 6, 8],
[4, 5, 6]]*/

let v = 7;
let e = 9;
let arr = [
  [0, 2, 7],
  [0, 4, 1],
  [0, 6, 5],
  [1, 2, 10],
  [1, 6, 3],
  [2, 4, 10],
  [2, 5, 10],
  [3, 5, 8],
  [3, 6, 6],
];

const sol = new Solution();
let res = sol.spanningTree(arr, v, e);
console.log(res);
