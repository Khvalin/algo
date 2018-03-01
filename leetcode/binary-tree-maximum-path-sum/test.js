var List = require('leetcode').List;

// { val: 1, next: { val: 2, next: { val: 3, next: null } } }
var l = List.create([1, 2, 3]);

// [1, 2, 3]
console.log(l.toArray());