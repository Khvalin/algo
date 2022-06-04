/**
 * @param {number} n
 * @return {string[][]}
 */
var solveNQueens = function (n) {
    const a = []
    let p = []
    for (let i = 0; i < n; i++) {
        a.push(0)
        p.push('.')
    }


    const res = []
    const addAns = (b) => {
        const r = []
        for (let i = 0; i < n; i++) {
            r.push([...p])
        }

        for (let i = 0; i < n; i++) {
            r[i][b[i]] = 'Q'
        }

        for (let i = 0; i < n; i++) {
            r[i] = r[i].join('')
        }

        res.push(r)
    }

    const q = [[0, [...a]]]

    while (q.length) {
        const cur = q.pop()
        const [pos, a] = cur

        if (pos === n) {
            addAns(a)
            continue
        }

        outer:
        for (let i = 0; i < n; i++) {

            for (let j = 0; j < pos; j++) {
                const f = (a[j] === i) || (j + a[j] === i + pos) || (j + i === a[j] + pos)

                if (f) {
                    continue outer
                }

            }

            const c = [...a]
            c[pos] = i

            q.push([pos + 1, c])
        }
    }

    return res
};


console.log(solveNQueens(1))