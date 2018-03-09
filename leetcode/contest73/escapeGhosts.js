/**
 * @param {number[][]} ghosts
 * @param {number[]} target
 * @return {boolean}
 */
var escapeGhosts = function(ghosts, target) {
  const distanceSquared = function(pointA, pointB) {
    const r = Math.abs(pointA[0] - pointB[0]) + Math.abs(pointA[1] - pointB[1]);
    return r;
  };

  const myDistance = distanceSquared([ 0, 0 ], target);

  const result = ghosts.every((g) => distanceSquared(target, g) > myDistance);

  return result;
};

console.log(escapeGhosts([ [ 1, 0 ], [ 0, 3 ] ], [ 0, 1 ]));
