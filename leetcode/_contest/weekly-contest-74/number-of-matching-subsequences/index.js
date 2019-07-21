/**
 * @param {string} S
 * @param {string[]} words
 * @return {number}
 */
var numMatchingSubseq = function(S, words) {
  let count = 0;

  const positions = words.map(() => 0);
  const curSymbols = words.reduce((acc, w, i) => {
    const cur = acc.get([ w[0] ]) || {
      indices: []
    };

    if (!acc.has(w[0])) {
      acc.set(w[0], {
        indices: []
      });
    }

    acc.get(w[0]).indices.push(i);
    return acc;
  }, new Map());

  let chars = Array.from(curSymbols.keys());

  for (var i = 0; i < S.length; i++) {
    const ch = S[i];
    const index = chars.indexOf(ch);
    if (index > -1) {
      const search = curSymbols.get(ch);
      if (search) {
        const len = search.indices.length;

        for (let j = 0; j < len; j++) {
          //console.log(search.indices);
          const k = search.indices.shift();
          positions[k]++;

          //      console.log(words[k]);
          if (positions[k] < words[k].length) {
            const next = words[k][positions[k]];
            if (!curSymbols.has(next)) {
              curSymbols.set(next, {
                indices: []
              });
            }

            curSymbols.get(next).indices.push(k);
          } else {
            count += 1;
          }
        }
      }

      chars = Array.from(curSymbols.keys());
    }
  }

  // console.log(curSymbols);

  return count;
};

console.log(numMatchingSubseq('abcde', [ 'a', 'bb', 'acd', 'ace' ]));
