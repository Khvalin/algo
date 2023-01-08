function maxPoints(points: number[][]): number {
	const EPS = 1e-12;
	
	let res = Math.min(2, points.length)
	for (let i = 0; i < points.length; i++) {
		const [x1, y1] = points[i]
		for (let j = i + 1; j < points.length; j++) {
			const [x2, y2] = points[j]
			if (x1 === x2) {
				let c = 2
				for (let l = j + 1; l < points.length; l++) {
					const [x3, y3] = points[l]
					if (x3 === x2) {
						c++
					}
				}
				res = Math.max(c, res)
				continue
			}
			
			let c = 2
			const k = (y2 - y1) / (x2 - x1)
			const b = y1 - k * x1
			
			for (let l = j + 1; l < points.length; l++) {
				const [x3, y3] = points[l]
				if (x3 === x2 || x3 === x1) {
					continue
				}
				
				const k2 = (y3 - y1) / (x3 - x1)
				const b2 = y3 - k * x3

				if (Math.abs(k2 - k) < EPS && Math.abs(b2 - b) < EPS) {
					c++
				}
			}
			
			res = Math.max(res, c)
		}
	}
	
	return res
};

let points = [[1,1],[2,2],[3,3]]
points = [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
points = [[1,1],[1,2],[1,3],[1,10],[1,30],[1,40]]
points = [[-6,-1],[3,1],[12,3]]

console.log(maxPoints(points))
