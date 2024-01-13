function kmp(haystack, needle) {
  // Preprocess the needle to create the prefix function
  let pi = [0];
  let i = 1, j = 0;
  while (i < needle.length) {
    if (needle[i] === needle[j]) {
      pi[i] = j + 1;
      i++;
      j++;
    } else if (j > 0) {
      j = pi[j-1];
    } else {
      pi[i] = 0;
      i++;
    }
  }

  // Search for the needle in the haystack
  i = 0, j = 0;
  while (i < haystack.length) {
    if (needle[j] === haystack[i]) {
      i++;
      j++;
      if (j === needle.length) {
        return i - j;
      }
    } else if (j > 0) {
      j = pi[j-1];
    } else {
      i++;
    }
  }

  // Needle not found
  return -1;
}


/**
 * @param {string} haystack
 * @param {string} needle
 * @return {number}
 */
var strStr = function(haystack, needle) {
    return kmp(haystack, needle)
};

//console.log(strStr('hello', 'll'))
console.log(strStr("mississippi", "issip"))
