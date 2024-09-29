var AllOne = function () {
  this.count = new Map();
};

/**
 * @param {string} key
 * @return {void}
 */
AllOne.prototype.inc = function (key) {
  const c = this.count.get(key) ?? 0;
  this.count.set(key, c + 1);
};

/**
 * @param {string} key
 * @return {void}
 */
AllOne.prototype.dec = function (key) {
  const c = this.count.get(key);
  if (c == 1) {
    this.count.delete(key);
    return;
  }

  this.count.set(key, c - 1);
};

/**
 * @return {string}
 */
AllOne.prototype.getMaxKey = function () {
  let res = "";
  let max = 0;
  for (const [key, value] of this.count.entries()) {
    if (value > max) {
      max = value;
      res = key;
    }
  }

  return res;
};

/**
 * @return {string}
 */
AllOne.prototype.getMinKey = function () {
  let res = "";
  let min = 0;
  for (const [key, value] of this.count.entries()) {
    if (res == "" || value < min) {
      min = value;
      res = key;
    }
  }

  return res;
};

/**
 * Your AllOne object will be instantiated and called as such:
 * var obj = new AllOne()
 * obj.inc(key)
 * obj.dec(key)
 * var param_3 = obj.getMaxKey()
 * var param_4 = obj.getMinKey()
 */
