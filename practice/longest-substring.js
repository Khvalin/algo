const longestSubstring = function(array = []) {
  let minLength = Number.POSITIVE_INFINITY;
  let minLenthString = '';

  for (let i = 0; i < array.length; i++) {
    if (array[i].length < minLength) {
      minLength = array[i].length;
      minLenthString = array[i];
    }
  }

  const matchAll = function(subst) {
    let matchCount = 0;
    for (let i = 0; i < array.length; i++) {
      if (array[i].startsWith(subst)) {
        matchCount++;
      }
    }

    return matchCount === array.length;
  };

  let result = minLenthString;
  if (matchAll(result)) {
    return result;
  }

  minLength = Math.trunc(minLength / 2);

  let matchesAll = false;
  do {
    matchesAll = matchAll(minLenthString.substr(0, minLength));
    if (matchesAll) {
      minLength = Math.trunc(minLength + (minLenthString.length - minLength) / 2);
    } else {
      minLength = Math.trunc(minLength / 2);
    }
  } while (minLength > 0 && !(matchesAll && !matchAll(minLenthString.substr(0, minLength + 1))));

  return minLenthString.substr(0, minLength);
};

console.log(longestSubstring([ 'asvdfasdf', 'asvdfasd', 'asvv' ]));
