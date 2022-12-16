var assignBikes = function(workers, bikes) {
	const dist = (i, j) => {
		return Math.abs(workers[i][0] - bikes[j][0]) + Math.abs(workers[i][1] - bikes[j][1])
	}
	
    const N = workers.length + 2
    
	const memo = new Map()
    const solve = (w, b) => {
		if (!w || !b) {
			return 0
		}
        const key = (w << N) | b
        
        if (memo.has(key)) {
			return memo.get(key)
		}
		
		let res = Number.POSITIVE_INFINITY
		for (let i = 0; i < workers.length; i++) {
            if (!(w & (1 << i))) {
                continue
            }
			for (let j = 0; j < bikes.length; j++) {
                if (!(b & (1 << j))) {
                    continue
                }
				const w2 = w - (1 << i)
				const b2 = b - (1 << j)
				
				const d = Math.abs(workers[i][0] - bikes[j][0]) + Math.abs(workers[i][1] - bikes[j][1])
				const r = d + solve(w2, b2)

				res = Math.min(res, r)
			}
		}
		
		memo.set(key, res)
		
		return res
    }
    
    let w = 0
    let b = 0
    
    for (let i = 0; i < workers.length; i++) {
        w = w << 1
        w += 1
    }
    
    for (let i = 0; i < bikes.length; i++) {
        b = b << 1
        b += 1
    }
    
    const res = solve(w, b)
    
    return res
};
let res = 0
/*
res = assignBikes([[0,0],[2,1]], [[1,2],[3,3]])
console.log(res)

res = assignBikes([[0,0],[1,0],[2,0],[3,0],[4,0]], [[0,999],[1,999],[2,999],[3,999],[4,999]])
console.log(res)
*/

res = assignBikes([[0,0],[1,0],[2,0],[3,0],[4,0],[5,0]],
[[0,999],[1,999],[2,999],[3,999],[4,999],[5,999],[6,999],[7,999]])

res = assignBikes([[0,0],[1,0],[2,0],[3,0],[4,0],[5,0],[6,0],[7,0],[8,0],[9,0]],
[[0,999],[1,999],[2,999],[3,999],[4,999],[5,999],[6,999],[7,999],[8,999],[9,999]])

console.log(res)
