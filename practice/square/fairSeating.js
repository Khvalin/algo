const fairSeating = (h, b) => {
	const seating = new Array(h + b - 1)
	seating.fill('h')
	if (b === 0) {
		return seating
	}
	
	const d = h/b
	let i = 0
	let f = true
	while (i < seating.length && b > 0) {
		i += (f? Math.floor(d): Math.ceil(d))
		
		seating[i] = 'B'
		b--
		
		f = !f
	}
	
	return seating
}


let res = fairSeating(6, 4)
console.log(res.toString())
