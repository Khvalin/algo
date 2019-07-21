/**
 * @param {number} n
 * @param {number[][]} red_edges
 * @param {number[][]} blue_edges
 * @return {number[]}
 */
var shortestAlternatingPaths = function(n, red_edges, blue_edges) {
  let nodes = []

  const createEdge = (color, dest) => {
    return {
      color,
      dest
    }
  }

  const res = [],
    visited = {}

  for (let i = 0; i < n; i++) {
    res.push(Number.POSITIVE_INFINITY)
    nodes.push({ edges: [] })
  }

  for (const e of red_edges) {
    nodes[e[0]].edges.push(createEdge('r', e[1]))
  }

  for (const e of blue_edges) {
    nodes[e[0]].edges.push(createEdge('e', e[1]))
  }

  let q = [ { last: '', node: 0, len: 0 } ]
  while (q.length > 0) {
    const cur = q.shift()

    const key = `${cur.node} ${cur.last}`
    if (visited[key]) {
      continue
    }
    //  console.log(key, cur.len, res[cur.node])

    visited[key] = true
    res[cur.node] = Math.min(res[cur.node], cur.len)

    const edges = nodes[cur.node].edges
    for (const e of edges) {
      if (e.color !== cur.last) {
        q.push({ last: e.color, node: e.dest, len: cur.len + 1 })
      }
    }
  }

  return res.map((len) => (len > 10 + n ? -1 : len))
}
/* 
console.log(shortestAlternatingPaths(3, [ [ 0, 1 ], [ 1, 2 ] ], []))
console.log(shortestAlternatingPaths(3, [ [ 0, 1 ] ], [ [ 2, 1 ] ]))
console.log(shortestAlternatingPaths(3, [ [ 1, 0 ] ], [ [ 2, 1 ] ]))
console.log(shortestAlternatingPaths(3, [ [ 0, 1 ] ], [ [ 1, 2 ] ]))
console.log(shortestAlternatingPaths(3, [ [ 0, 1 ], [ 0, 2 ] ], [ [ 1, 0 ] ])) */
console.log(shortestAlternatingPaths(5, [ [ 0, 1 ], [ 1, 2 ], [ 2, 3 ], [ 3, 4 ] ], [ [ 1, 2 ], [ 2, 3 ], [ 3, 1 ] ]))
