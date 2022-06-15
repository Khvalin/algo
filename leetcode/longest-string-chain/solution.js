/**
 * @param {string[]} words
 * @return {number}
 */
var longestStrChain = function (words) {
    words = words.sort((a, b) => a.length - b.length)

    const canChain = (w1, w2) => {
        if (w1.length !== w2.length - 1) {
            return false
        }

        let i = 0
        let j = 0
        while (j < w2.length) {
            if (w1[i] === w2[j]) {
                i++
            }

            j++
        }

        return i === w1.length && j === w2.length
    }

    const memo = words.map(() => -1)
    const chainable = new Map()

    const solve = (ind) => {
        if (ind >= words.length) {
            return 0
        }

        if (memo[ind] > 0) {
            return memo[ind]
        }
        const w = words[ind];

        let res = 1
        let i = ind + 1
        while (i < words.length && words[i].length !== w.length + 1) {
            i++
        }

        while (i < words.length && words[i].length === w.length + 1) {
            const key = ind * 1001 + i
            let f = chainable.get(key) || false
            if (!chainable.has(key)) {
                f = canChain(w, words[i])
                chainable.set(key, f)
            }

            if (f) {
                res = Math.max(res, 1 + solve(i))
            }
            i++
        }

        memo[ind] = res
        return res
    }

    let res = 0
    for (let i = 0; i < words.length; i++) {
        if (memo[i] < 0) {
            res = Math.max(res, solve(i))
        }
    }

    return res

};

console.log(longestStrChain(["a", "b", "ba", "bca", "bda", "bdca"]), 4)