/**
 * @param {number[]} arr
 * @return {number}
 */
var minJumps = function(arr) {
    const values = new Map()
    
    for (let i = 0; i < arr.length; i++) {
        const n = arr[i]
        if (!values.has(n)) {
            values.set(n, [])
        }
        
        values.get(n).push(i)
    }
    
    const visited = new Set([0])

    let res = 0
    let nodes = [0]
    outer:
    while (nodes.length) {
        const next = new Set()
        for (const i of nodes) {
            visited.add(i)
            if (i === arr.length - 1) {
                break outer
            }
            
            if (i > 0 && !visited.has(i-1)) {
                next.add(i-1)
                visited.add(i-1)
            }
            
            if ((i < arr.length - 1) && !visited.has(i+1)) {
                next.add(i+1)
                visited.add(i+1)
            }
            
            for (const j of values.get(arr[i])) {
                if (!visited.has(j)) {
                    next.add(j)
                    visited.add(j)
                }
            }
            values.set(arr[i], [])
        }
        nodes = [...next]
        res++
    }
    
    return res
};
