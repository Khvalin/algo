/**
 * @param {number} target
 * @param {number} startFuel
 * @param {number[][]} stations
 * @return {number}
 */
var minRefuelStops = function(target, startFuel, stations,  stops = 0) {
  console.log(target, startFuel)
  if (startFuel >= target){
    return stops ;
  }


  let nextStation = stations.shift() ;
  if (! nextStation ){
    return -1;
  }
  console.log(nextStation)

  startFuel -= nextStation[0];

  if (startFuel >= 0) {  
    var fillUp = minRefuelStops(target - nextStation[0], startFuel + nextStation[1], stations, stops + 1);

    var skip = minRefuelStops(target - nextStation[0], startFuel, stations, stops);

    if (skip < 0){
      return fillUp;
    } else {
      return Math.min(skip, fillUp)
    }
  }

  console.log(startFuel)

  return -1;
};

console.log(minRefuelStops(100, 10, [[10,60],[20,30],[30,30],[60,40]]))