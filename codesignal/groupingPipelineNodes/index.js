function groupingPipelineNodes(n, g, v) {
  n += 1
  const graph = []
  for (let i = 0; i <= n; i++) {
    graph.push(new Set())
  }

  const mergeVertices = new Set([ ...v ])
  let mergePoint = 0

  //  console.log(mergePoint, mergeVertices)
  for (const edge of g) {
    let [ from, to ] = edge
    if (mergeVertices.has(from)) {
      from = mergePoint
    }

    if (mergeVertices.has(to)) {
      to = mergePoint
    }
    if (from !== to) {
      graph[from].add(to)
    }
  }

  // console.log(graph)

  const bfs = (s, component, recStack = new Set()) => {
    if (recStack.has(s)) {
      return false
    }

    if (component.has(s)) {
      return true
    }

    component.add(s)
    recStack.add(s)

    for (let next of graph[s].values()) {
      if (!bfs(next, component, recStack)) {
        return false
      }
    }
    recStack.delete(s)

    return true
  }

  let res = true
  let component = new Set()
  for (let i = 0; res && i < n; i++) {
    // console.log(visited)
    if (!component.has(i)) {
      res = res && bfs(i, component)
    }
  }

  return res
}

const t1 = groupingPipelineNodes(6, [ [ 1, 2 ], [ 2, 3 ], [ 2, 4 ], [ 2, 5 ], [ 4, 6 ], [ 5, 6 ] ], [ 4, 5 ])
console.log(true, t1)

const t2 = groupingPipelineNodes(6, [ [ 1, 2 ], [ 2, 3 ], [ 2, 4 ], [ 2, 5 ], [ 4, 6 ], [ 5, 6 ] ], [ 2, 6 ])
console.log(false, t2)

const t5 = groupingPipelineNodes(6, [ [ 1, 2 ], [ 2, 3 ], [ 2, 4 ], [ 2, 5 ], [ 4, 6 ], [ 5, 6 ] ], [ 4 ])
console.log(t5)

const t6 = groupingPipelineNodes(6, [ [ 1, 2 ], [ 2, 3 ], [ 2, 4 ], [ 2, 5 ], [ 4, 6 ], [ 5, 6 ] ], [])
console.log(t6)

const _t1 = groupingPipelineNodes(3, [ [ 1, 1 ], [ 2, 2 ], [ 3, 3 ] ], [ 1, 2 ])
console.log(true, _t1)
