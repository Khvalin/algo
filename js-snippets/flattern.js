const flattern = function(array) {
  const q = [...array]
  const res = []

  while (q.length > 0) {
  	const elem = q.shift()
  	if (Array.isArray(elem)) {
  		q.unshift(...elem)
  		continue
  	}

  	res.push(elem)
  }

  return res
};

const res = flattern([ 1, [ 2, [ 3,[4,5],6 ] ], 10 ]);
console.log(res.join(' '));
