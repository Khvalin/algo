'use strict'

const createTree = (nodeValues) => {
  const nodes = {}
  for (let i = 1; i <= nodeValues.length; i++) {
    const val = nodeValues[i - 1]
    if (val === null){
       continue
    }
    const node = { val, left: null, right: null }
    if (i > 1) {
      let parent = i >> 1

      if (i % 2 == 0) {
        nodes[parent].left = node
      } else {
        nodes[parent].right = node
      }
    }
    nodes[i] = node
  }

  return nodes[1] || null
}

export { createTree }
