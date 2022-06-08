/**
 * THIS DOESN'T SOLVE THE PROBLEM AS IT REMOVES SUBSTRINGS NOT subsequences
 * BUT IT LIKE IT MUCH MORE
 * @param {string} s
 * @return {number}
 */
var removePalindromeSub = function (s) {
    const getLongest = (s, ind) => {
        const res = [ind, ind]
        for (let k = 0; k < 2; k++) {
            let i = ind
            let j = ind + k

            while (i >= 0 && j < s.length && s[i] === s[j]) {
                // make sure the invariant is always valid
                if (j > res[1]) {
                    res[0] = i
                    res[1] = j
                }

                i--
                j++
            }
        }
        return res
    }

    const q = [[s.split(''), 0]]
    let res = s.length + 10

    while (q.length) {
        let [s, len] = q.pop();
        if (!s.length) {
            res = Math.min(res, len)
            continue
        }


        for (let i = 0; i < s.length; i++) {
            const a = getLongest(s, i)
            const next = [...s]
            next.splice(a[0], a[1]-a[0]+1)
            q.push([next, 1 + len])
        }
    }

    return res
};

console.log(removePalindromeSub("bbaabaaa",2))
// console.log(removePalindromeSub("abb"),2)
// console.log(removePalindromeSub("abaaaaaabbba"),2)
// console.log(removePalindromeSub("abbba"),1)