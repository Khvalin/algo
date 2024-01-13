const generateGrid = (n, m) => {
	const res = []
	for (let i = 0; i < n; i++) {
		const a = []
		res.push(a)
		for (let j = 0; j < n; j++) {
			a.push((Math.floor(j/m) % 2) * 255)
		}
	}
	
	return res
}


console.log(generateGrid(10, 2))
