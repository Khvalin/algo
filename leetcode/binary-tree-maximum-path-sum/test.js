var createTree = require('leetcode').Tree.create
var maxPathSum = require('./index').maxPathSum

var l = createTree([ 1, 2, 3 ])
console.log(maxPathSum(l))

l = createTree([ 1, 2, 3, null, null, 4, null, null, 5 ])
console.log(maxPathSum(l))

l = createTree([ 0, 1, 2, 3, 100, 50, 101 ])
console.log(maxPathSum(l))
