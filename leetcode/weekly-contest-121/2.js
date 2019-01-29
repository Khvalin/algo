/**
 * Initialize your data structure here.
 */
var TimeMap = function() {
  this._storage = {}
}

/** 
 * @param {string} key 
 * @param {string} value 
 * @param {number} timestamp
 * @return {void}
 */
TimeMap.prototype.set = function(key, value, timestamp) {
  this._storage[key] = this._storage[key] || []

  this._storage[key].push({ timestamp, value })
}

/** 
 * @param {string} key 
 * @param {number} timestamp
 * @return {string}
 */
TimeMap.prototype.get = function(key, timestamp) {
  const arr = this._storage[key]

  const binarySearch = (left, right) => {
    if (left >= right) {
      return right
    }

    const mid = Math.floor((right - left) / 2)
    // console.log(left, right, mid)
    if (arr[mid].timestamp <= timestamp) {
      return binarySearch(mid, right)
    } else {
      return binarySearch(left, mid - 1)
    }
  }

  if (!arr || !arr.length || arr[0].timestamp > timestamp) {
    return ''
  }

  let i = binarySearch(0, arr.length - 1)

  if (Number.isInteger(i)) {
    while (i < arr.length - 2 && arr[i].timestamp < timestamp) {
      i++
    }

    return arr[i] ? arr[i].value : ''
  }

  return ''
}

/** 
 * Your TimeMap object will be instantiated and called as such:
 * var obj = Object.create(TimeMap).createNew()
 * obj.set(key,value,timestamp)
 * var param_2 = obj.get(key,timestamp)
 */
/* 
var obj = new TimeMap()
obj.set('foo', 'bar', 1)
console.log(obj.get('foo', 1))
console.log(obj.get('foo', 3))

obj.set('foo', 'bar2', 4)
console.log(obj.get('foo', 4))
console.log(obj.get('foo', 5))
 */

var obj = new TimeMap()

obj.set('love', 'high', 10)
obj.set('love', 'low', 20)

console.log(obj.get('love', 5))
console.log(obj.get('love', 10))
console.log(obj.get('love', 15))
console.log(obj.get('love', 20))
console.log(obj.get('love', 25))
