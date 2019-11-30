/**
 * @param {number[][]} intervals
 * @param {number[]} toBeRemoved
 * @return {number[][]}
 */
var removeInterval = function(intervals, toBeRemoved) {
  const res = [];

  const [ a, b ] = toBeRemoved;
  for (let i = 0; i < intervals.length; i++) {
    let [ x, y ] = intervals[i];

    if (x > a && y < b) {
      continue;
    }

    if (a > x && b < y) {
      res.push([ x, a ]);
      res.push([ b, y ]);
      continue;
    }

    if (a > x && a < y) {
      y = a;
    }

    if (b > x && b < y) {
      x = b;
    }

    res.push([ x, y ]);
  }

  return res;
};

let intervals = [ [ 0, 2 ], [ 3, 4 ], [ 5, 7 ] ],
  toBeRemoved = [ 1, 6 ];

console.log(removeInterval(intervals, toBeRemoved));
(intervals = [ [ 0, 5 ] ]), (toBeRemoved = [ 2, 3 ]);
console.log(removeInterval(intervals, toBeRemoved));
