const covert = (s) => {
	const res = {}
	const lines = s.split('\n')

	
	let cur = 0
	const stack = [res]
	for (const line of lines) {
		let obj = stack[stack.length - 1]
		let [key, val] = line.split(':')
		const k = key.trimLeft()
		val = val.trim()
		const level = (key.length - k.length) >> 1
		if (val) {
			cur = level
			obj[k] = val
			continue
		}
		
		

		if (level > cur) {
			obj[k] = {}
			stack.push(obj[k])
			
			continue
		}
		
		while(level < cur) {
			cur--
			stack.pop()
		}
		obj = stack[stack.length - 1]
		obj[k] = {}
		stack.push(obj[k])
	}
	
	return res
}

const input = 
`K1:V1
K2:V2
K3:
  K31:V31
  K32:
    K321:V321
    K322:V322
    K33:V33
K4:
  K41:V41
  K42:V42`

console.log(covert(input))
