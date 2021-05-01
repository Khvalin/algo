
/**
 * @param {string[]} words
 */
var WordFilter = function(words) {
  this.dict = new Map()
  for (let i = 0; i < words.length; i++) {
    const w = words[i]
    for (let j = 1; j <= w.length; j++) {
      const prefix = w.substring(0, j)
      for (let k = 0; k < w.length; k++) {
        const suffix = w.substring(k)
        this.dict.set(`${prefix}-${suffix}`, i)
      }
    }
  }
};

/** 
 * @param {string} prefix 
 * @param {string} suffix
 * @return {number}
 */
WordFilter.prototype.f = function(prefix, suffix) {
  return this.dict.get(`${prefix}-${suffix}`) ?? -1
};

/** 
 * Your WordFilter object will be instantiated and called as such:
 * var obj = new WordFilter(words)
 * var param_1 = obj.f(prefix,suffix)
 */
