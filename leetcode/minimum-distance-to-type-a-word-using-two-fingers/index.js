/**
 * @param {string} word
 * @return {number}
 */
const minimumDistance = (word) => {
  const getCoords = (ch) => {
    const ind = ch.length ===1? ch.charCodeAt(0) - 'A'.charCodeAt(0):0

    return [ Math.floor(ind / 6) , (ind % 6) ]
  }

  const states = [{l:'', r:'', cost:0}]

  let r = Number.POSITIVE_INFINITY
  for (const ch of word) {
    const l = states.length
    const coord = getCoords(ch)
    for (let i = 0; i < l;i++) {
      const state = states.shift();

      let dl = 0
      if (state.l !== '') {

      }
    }
  }

  return r
};


let input = 'CAKE'
console.log(minimumDistance(input))
