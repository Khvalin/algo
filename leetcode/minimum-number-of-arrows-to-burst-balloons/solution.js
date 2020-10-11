/**
 * @param {number[][]} points
 * @return {number}
 */
var findMinArrowShots = function(points) {
  points = points.sort((a, b) => {
    if (a[0] === b[0]) {
      return a[1] - b[1];
    }

    return a[0] - b[0];
  });

  let res = 0;
  let right = Number.NEGATIVE_INFINITY;
  for (const p of points) {
    if (p[0] > right) {
      res++;
      right = p[1];
      continue;
    }

    right = Math.min(p[1], right);
  }

  return res;
};

//console.log(findMinArrowShots([[10,16],[2,8],[1,5], [1,6],[7,12]]));
console.log(findMinArrowShots([ [ 9, 12 ], [ 1, 10 ], [ 4, 11 ], [ 8, 12 ], [ 3, 9 ], [ 6, 9 ], [ 6, 7 ] ]));
