/**
 * @param {string} S
 * @param {string} T
 * @return {string}
 */
var customSortString = function(S, T) {
  const sortFunction = (a, b) => {
    if (S.indexOf(a) === -1 && S.indexOf(b) > -1) {
      return 1;
    }

    if (S.indexOf(b) === -1) {
      if (S.indexOf(a) > -1) {
        return -1;
      } else {
        return 0;
      }
    }

    return S.indexOf(a) - S.indexOf(b);
  };

  return T.split('')
    .sort((a, b) => {
      const result = sortFunction(a, b);
      //console.log(a, b, result);

      return result;
    })
    .join('');
};

console.log(customSortString('exv', 'xwvee'));
