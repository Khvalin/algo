/**
 * // This is the master's API interface.
 * // You should not implement it, or speculate about its implementation
 * function Master() {
 *
 *     @param {string[]} wordlist
 *     @param {Master} master
 *     @return {integer}
 *     this.guess = function(word) {
 *         ...
 *     };
 * };
 */
/**
 * @param {string[]} wordlist
 * @param {Master} master
 * @return {void}
 */
var findSecretWord = function(wordlist, master) {
  const wordMatch = (w1, w2) => {
    let res = 0;
    for (let j = 0; j < word.length; j++) {
      if (w2[j] === w1[i]) {
        res++;
      }
    }

    return res;
  };

  let res = 0;
  let prev = 0;
  let tries = 0;
  let w = wordlist.pop();
  let candidates = [];
  do {
    tries++;
    res = master.guess(w);
    if (res === 6) {
      return;
    }
    if (res === 0) {
      let i = 0;
      while (i < wordlist.length) {
        if (wordMatch(wordlist[i], w) > 0) {
          wordlist.splice(i, 1);
          continue;
        }
        i++;
      }
      w = wordlist.pop();
      continue;
    }
    if (prev < res) {
      for (let i = 0; i < wordlist.length; i++) {
        if (wordMatch(wordlist[i], w) >= res) {
          candidates.push(wordlist[i]);
        }
      }
    }

    prev = res;
  } while (tries < 10 && res < 6);
};
