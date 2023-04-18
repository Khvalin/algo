
//User function Template for javascript

/**
 * @param {string} S
 * @return {string[]}
 */

class Solution {
    find_permutation(S){
         const chars = [...S]
         chars.sort()

         const unique = new Set()

         const solve = (a) => {
             if (!a.length) {
                 return []
             }

            if (a.length === 1) {
                 return [...a]
            }

            const results = []
            for (let i = 0; i < a.length; i++) {
                const copy = [...a]
                copy.splice(i, 1)
                
				const res = solve(copy)
                for (const s of res) {
                    const str = a[i] + s

                    if (unique.has(str)) {
                        continue
                    }

                    unique.add(str)
                    results.push(str)
                }
            }

            return results
         }

        return solve(chars)
    }
}

const obj = new Solution()
console.log(obj.find_permutation('ABCD'))
