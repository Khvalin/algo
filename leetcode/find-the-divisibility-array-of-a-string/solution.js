/**
 * @param {string} word
 * @param {number} m
 * @return {number[]}
 */
var divisibilityArray = function(word, m) {
    const res = new Array(word.length)
    res.fill(0)

    let rem = 0

    for(let i = 0; i < word.length; i++) {
        rem *= 10
        rem += (1 * word[i])
        rem %= m
        
        if (rem === 0) {
            res[i] = 1
        }
    }
    
    return res
};
