/**
 * @param {string} start
 * @param {string} end
 * @param {string[]} bank
 * @return {number}
 */
var minMutation = function(start, end, bank) {
	const d = (a,b) => {
		let res = Math.abs(a.length - b.length)
		for (let i = 0; i < a.length; i++) {
			if (a[i] !== b[i]) {
				res++
			}
		}
		
		return res
	}
	
    let res  = bank.length + 10
    const q = [[start, new Set([start]), 0]]
    
    while (q.length) {
		const [cur, visited, dist] = q.pop()
		if (cur === end) {
			res = Math.min(res, dist)
			continue
		}
		
		for (const dna of bank) {
			if (visited.has(dna) || d(cur, dna) !== 1) {
				continue
			}
			
			const copy = new Set([...visited, dna])
			q.push([dna, copy, dist + 1])
		}
	}
    
    if (res > bank.length) {
        return -1
    }
    
    return res
};

let start = "AAAAACCC", end = "AACCCCCC", bank = ["AAAACCCC","AAACCCCC","AACCCCCC"];
console.log(minMutation(start, end,bank))
