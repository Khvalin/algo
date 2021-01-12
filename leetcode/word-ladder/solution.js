/**
 * @param {string} beginWord
 * @param {string} endWord
 * @param {string[]} wordList
 * @return {number}
 */
var ladderLength = function (beginWord, endWord, wordList) {
  const ABC = "abcdefghijklmnopqrstuvwxyz";
  const words = new Map();
  const dist = [];
  const adj = [];
  let endIndex = -1;

  const q = [0];
  wordList.unshift(beginWord);
  for (let i = 0; i < wordList.length; i++) {
    dist.push(Number.POSITIVE_INFINITY);
    adj.push([]);
    const w = wordList[i];
    if (words.has(w)) {
      continue;
    }
    if (w === endWord) {
      endIndex = i;
    }

    words.set(w, i);
  }

  for (const w of wordList) {
    const a = w.split("");
    for (let i = 0; i < a.length; i++) {
      const t = a[i];
      for (const ch of ABC) {
        if (ch === t) {
          continue;
        }
        a[i] = ch;
        const key = a.join("");
        if (words.has(key)) {
          adj[words.get(w)].push(words.get(key));
        }
      }
      a[i] = t;
    }
  }

  dist[0] = 0;
  while (q.length) {
    const ind = q.shift();
    let d = dist[ind];

    for (const i of adj[ind]) {
      if (dist[i] > d + 1) {
        dist[i] = d + 1;
        q.push(i);
      }
    }
  }

  let res = 0;
  if (endIndex > -1 && Number.isFinite(dist[endIndex])) {
    res = dist[endIndex] + 1;
  }

  return res;
};
