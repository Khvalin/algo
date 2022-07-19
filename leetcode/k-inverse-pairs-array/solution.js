/**
 * @param {number} n
 * @param {number} k
 * @return {number}
 */
 var kInversePairs = function(n, k) {
    const MOD = 1_000_000_000 + 7
    let a = [1]

    const total = (n * (n - 1)) >> 1
    const mid = Math.ceil(total / 2) + 1

    for (let i = 2; i <= n; i++) {
        const total = (i * (i - 1)) >> 1
        const mid = total / 2 + 1
        const next = []
        
        for (let k = 0; k <= total; k++) {
            let t = 0
            for (let j = 1; j <= i; j++) {
                
                const c = i - j
                let b = 0
                if (c > k || k -c >= a.length) {
                    continue
                }


                b = a[k -c]


                t += b % MOD
                t %= MOD
            }

            next.push(t)
        }

        a = next
    }


    if (k >= mid) {
        k = total - k
    }

    return a[k]
 }

// console.log(kInversePairs(500,990))

for (let i = 2; i < 9; i++) {
    const a = []
    const total = (i * (i - 1)) >> 1

    for (let j = 0; j <= total; j++) {
        a.push(kInversePairs(i,j))
    }

    console.log(i, a.join(' '))
}

// console.log(kInversePairs(5,0))
// console.log(kInversePairs(5,1))
// console.log(kInversePairs(5,2))
// console.log(kInversePairs(5,3))
// console.log(kInversePairs(5,4))
// console.log(kInversePairs(5,5))
// console.log(kInversePairs(5,6))
// console.log(kInversePairs(5,7))
// console.log(kInversePairs(5,8))
// console.log(kInversePairs(5,9))
// console.log(kInversePairs(5,10))