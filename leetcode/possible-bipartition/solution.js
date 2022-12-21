var possibleBipartition = function(n, dislikes) {
	const adj = []
	const groups = []
	for (let i = 0; i <= n; i++) {
		adj.push([])
		groups.push(0)
	}
	for (const [u, v] of dislikes) {
		adj[u].push(v)
        adj[v].push(u)
	}
		
	const dfs = (ind) => {
        const seen = new Set()
		const q = [ind]
		groups[ind] = groups[ind] || 1
		while (q.length) {
			const u = q.pop()
			const g = groups[u]
			
			for (const v of adj[u]) {
                if (seen.has(v)) {
                    continue
                }
				if (!groups[v]) {
					q.push(v)
					groups[v] = g === 1?2:1
				}
				if (groups[v] === g) {
					return false
				}
                seen.add(v)
			}
		}
		
		return true
	}
	
	for (let i = 1; i <= n; i++) {
		const res = dfs(i)
		if  (!res) {
			return false
		}

	}
	
	return true
}

let res = 0
res = possibleBipartition(4, [[1,2],[1,3],[2,4]])
console.log(res, true)

res = possibleBipartition(3, [[1,2],[1,3],[2,3]])
console.log(res, false)
