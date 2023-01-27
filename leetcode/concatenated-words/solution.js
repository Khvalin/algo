/**
 * @param {string[]} words
 * @return {string[]}
 */
var findAllConcatenatedWordsInADict = function (words) {
  const END = "$";
  const trie = {};

  let maxLen = 0;
  for (const w of words) {
    maxLen = Math.max(w.length, maxLen);
  }

  const buckets = new Array(maxLen + 1);
  for (let i = 0; i < words.length; i++) {
    const w = words[i];
    const len = w.length;
    buckets[len] = buckets[len] || [];
    buckets[len].push(i);
  }

  const canSplit = (wi, ind) => {
    const w = words[wi];
    if (ind >= w.length) {
      return true;
    }

    let res = false;
    let t = trie;
    for (let i = ind; !res && (i < w.length); i++) {
      const ch = w[i];

      if (!(ch in t)) {
        break;
      }

      t = t[ch];
      if (t[END]) {
        res = res || canSplit(wi, i + 1);
      }
    }

    return res;
  };

  const res = [];
  for (const bucket of buckets) {
    if (!bucket || !bucket.length) {
      continue;
    }

    for (const i of bucket) {
      const w = words[i];
      if (canSplit(i, 0)) {
        res.push(w);
      }

      let t = trie;
      for (const ch of w) {
        t[ch] = t[ch] || {};
        t = t[ch];
      }
      t[END] = true;
    }
  }

  return res;
};

let words = ["cat", "cats", "catsdogcats", "dog"];
console.log(findAllConcatenatedWordsInADict(words));

