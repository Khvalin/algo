const flattern = function(array, index = 0, result = []) {
  if (index >= array.length) {
    return result;
  }

  const next = array[index];
  if (next.length) {
    flattern(next, 0, result);
  } else {
    result.push(next);
  }
  return flattern(array, index + 1, result);
};

const res = flattern([ 1, [ 2, [ 3 ] ], 4, 5 ]);
console.log(res.join(' '));
