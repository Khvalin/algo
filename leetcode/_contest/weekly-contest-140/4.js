/**
 * @param {string} text
 * @return {string}
 */
var smallestSubsequence = function(text) {
  const dict = {}
  // Count all distinct characters.
  let dist_count = 0

  for (let i = 0; i < text.length; i++) {
    let ch = text[i]
    if (!dict[ch]) {
      dict[ch] = 1
      dist_count++
    }
  }

  let start = 0,
    start_index = -1,
    min_len = Number.POSITIVE_INFINITY

  let count = 0
  let curr_count = {}

  for (let j = 0; j < text.length; j++) {
    let ch = text[j]

    curr_count[ch] = curr_count[ch] || 0
    curr_count[ch]++

    // If any distinct character matched,
    // then increment count
    if (curr_count[ch] == 1) {
      count++
    }

    // if all the characters are matched
    if (count == dist_count) {
      // Try to minimize the window i.e., check if
      // any character is occurring more no. of times
      // than its occurrence in pattern, if yes
      // then remove it from starting and also remove
      // the useless characters.
      while (curr_count[text[start]] > 1) {
        if (curr_count[text[start]] > 1) curr_count[text[start]]--
        start++
      }

      // Update window size
      let len_window = j - start + 1
      if (min_len > len_window) {
        min_len = len_window
        start_index = start
      }
    }
  }

  // Return substring starting from start_index
  // and length min_len
  return text.substring(start_index, start_index + min_len)
}

console.log(smallestSubsequence('aabcbcdbca'))
console.log(smallestSubsequence('cdadabcc'))
