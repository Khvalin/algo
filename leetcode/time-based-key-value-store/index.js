/**
 * Initialize your data structure here.
 */
var TimeMap = function () {
  this.stores = {};
  this.timestamps = {};
  this.binarySearch = (arr = [], n) => {
    let [lo, hi] = [0, arr.length - 1];
    while (lo < hi && !(arr[lo] <= n && arr[lo + 1] >= n)) {
      const mid = (hi + lo) >> 1;
      if (arr[mid] > n) {
        hi = mid - 1;
      } else {
        lo = mid;
        if (arr[mid] < n) {
          lo++;
        }
      }
    }

    return lo;
  };
};

/**
 * @param {string} key
 * @param {string} value
 * @param {number} timestamp
 * @return {void}
 */
TimeMap.prototype.set = function (key, value, timestamp) {
  this.stores[`${key}_${timestamp}`] = value;
  this.timestamps[key] = this.timestamps[key] || [];

  if (!this.timestamps[key].length || this.timestamps[key][0] >= timestamp) {
    this.timestamps[key].unshift(timestamp);
    return;
  }

  if (this.timestamps[key][this.timestamps[key].length - 1] <= timestamp) {
    this.timestamps[key].push(timestamp);
    return;
  }

  const i = this.binarySearch(this.timestamps[key], timestamp);
  this.timestamps[key].splice(i + 1, 0, timestamp);
};

/**
 * @param {string} key
 * @param {number} timestamp
 * @return {string}
 */
TimeMap.prototype.get = function (key, timestamp) {
  if (!this.timestamps[key] || !this.timestamps[key].length || this.timestamps[key][0] > timestamp) {
    return '';
  }

  if (this.timestamps[key][this.timestamps[key].length - 1] <= timestamp) {
    const ts = this.timestamps[key][this.timestamps[key].length - 1];
    return this.stores[`${key}_${ts}`];
  }

  const i = this.binarySearch(this.timestamps[key], timestamp);
  timestamp = this.timestamps[key][i];

  return this.stores[`${key}_${timestamp}`] || '';
};

let timeMap = new TimeMap();
let stamps = [10, 7, 4, 2, 3, 13, 5];
for (const i of stamps) {
  timeMap.set('k', i.toString(), i);
}

console.log(timeMap.timestamps);


/**
 * Your TimeMap object will be instantiated and called as such:
 * var obj = new TimeMap()
 * obj.set(key,value,timestamp)
 * var param_2 = obj.get(key,timestamp)
 */
