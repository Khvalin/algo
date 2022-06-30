/**
 * @param {number[]} target
 * @return {boolean}
 */
var isPossible = function (target) {
    target.sort((a, b) => a - b)
    const a = target.map(() => 1)

    let sum = a.length

    let cur = 0
    while (a.length && a[0] === target[0]) {
        a.shift()
        target.shift()
    }

    if (!a.length) {
        return true
    }

    const q = [[a, sum, 0]]
    while (q.length) {
        let [a, sum, cur] = q.pop()
      //  console.log(a, sum)

        if (cur >= target.length) {
            return true
        }

        if (sum > target[cur]) {
            continue
        }

        if (sum === target[cur]) {
            sum += sum - a[0]
            const copy = [...a]
            copy.shift()
            q.push([copy, sum, cur + 1])
            continue
        }

        let visited = new Set()
        let res = false
        for (let i = 0; !res && i < a.length; i++) {
            if (visited.has(a[i])) {
                continue
            }
            visited.add(a[i])
            const copy = [...a]
            let t = sum
            copy[i] = t
            t += copy[i] - a[i]

            q.push([copy, t, cur])
        }
    }


    return false

};


let res = isPossible([1,1000000000])
console.log(res)