/**
 * @param {number} day
 * @param {number} month
 * @param {number} year
 * @return {string}
 */
var dayOfTheWeek = function(day, month, year) {
  const date = new Date(year, month - 1, day)

  return [ 'Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday' ][date.getDay()]
}

console.log(dayOfTheWeek(15, 8, 1993))
