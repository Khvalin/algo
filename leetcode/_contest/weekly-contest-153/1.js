/**
 * @param {number[]} distance
 * @param {number} start
 * @param {number} destination
 * @return {number}
 */
var distanceBetweenBusStops = function(distance, start, destination) {
  let d1 = 0
  let d2 = 0

  let i = start
  while (i != destination) {
    d1 += distance[i]
    i = (i + 1) % distance.length
  }

  i = destination
  while (i != start) {
    d2 += distance[i]
    i = (i + 1) % distance.length
  }

  return Math.min(d1, d2)
}
