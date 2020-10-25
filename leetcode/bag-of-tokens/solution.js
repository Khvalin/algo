/**
 * @param {number[]} tokens
 * @param {number} P
 * @return {number}
 */
var bagOfTokensScore = function(tokens, P) {
  tokens = tokens.sort((a, b) => a-b)
  
  let p = P
  let score = 0
  let res = 0
  let i = 0
  let j = tokens.length - 1
  
  let f = true
  while (f && (i <= j)) {
    f = false
    if (p >= tokens[0]) {
      p -= tokens[i]
      i++
      score++
      f = true
    } else {
      if (score > 0) {
        p += tokens[j]
        j--
        score--
        f = true
      }
    }
    
    res = Math.max(res, score)
  }

  return res
};

let tokens = [100,200,300,400], P = 200
console.log(bagOfTokensScore(tokens, P))
