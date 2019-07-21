/**
 * @param {string[]} set1
 * @param {string[]} set2
 * @return {string[]}
 */
const longestCommonSubsequence = function(set1, set2) {
  // Init LCS matrix.
  const lcsMatrix = Array(set2.length + 1).fill(null).map(() => Array(set1.length + 1).fill(null))

  // Fill first row with zeros.
  for (let columnIndex = 0; columnIndex <= set1.length; columnIndex += 1) {
    lcsMatrix[0][columnIndex] = 0
  }

  // Fill first column with zeros.
  for (let rowIndex = 0; rowIndex <= set2.length; rowIndex += 1) {
    lcsMatrix[rowIndex][0] = 0
  }

  // Fill rest of the column that correspond to each of two strings.
  for (let rowIndex = 1; rowIndex <= set2.length; rowIndex += 1) {
    for (let columnIndex = 1; columnIndex <= set1.length; columnIndex += 1) {
      if (set1[columnIndex - 1] === set2[rowIndex - 1]) {
        lcsMatrix[rowIndex][columnIndex] = lcsMatrix[rowIndex - 1][columnIndex - 1] + 1
      } else {
        lcsMatrix[rowIndex][columnIndex] = Math.max(
          lcsMatrix[rowIndex - 1][columnIndex],
          lcsMatrix[rowIndex][columnIndex - 1]
        )
      }
    }
  }

  // Calculate LCS based on LCS matrix.
  if (!lcsMatrix[set2.length][set1.length]) {
    // If the length of largest common string is zero then return empty string.
    return [ '' ]
  }

  const longestSequence = []
  let columnIndex = set1.length
  let rowIndex = set2.length

  while (columnIndex > 0 || rowIndex > 0) {
    if (set1[columnIndex - 1] === set2[rowIndex - 1]) {
      // Move by diagonal left-top.
      longestSequence.unshift(set1[columnIndex - 1])
      columnIndex -= 1
      rowIndex -= 1
    } else if (lcsMatrix[rowIndex][columnIndex] === lcsMatrix[rowIndex][columnIndex - 1]) {
      // Move left.
      columnIndex -= 1
    } else {
      // Move up.
      rowIndex -= 1
    }
  }

  return longestSequence
}

var shortestCommonSupersequence = function(set1, set2) {
  // Let's first find the longest common subsequence of two sets.
  const lcs = longestCommonSubsequence(set1, set2)

  // If LCS is empty then the shortest common supersequence would be just
  // concatenation of two sequences.
  if (lcs.length === 1 && lcs[0] === '') {
    return set1.concat(set2)
  }

  // Now let's add elements of set1 and set2 in order before/inside/after the LCS.
  let supersequence = []

  let setIndex1 = 0
  let setIndex2 = 0
  let lcsIndex = 0
  let setOnHold1 = false
  let setOnHold2 = false

  while (lcsIndex < lcs.length) {
    // Add elements of the first set to supersequence in correct order.
    if (setIndex1 < set1.length) {
      if (!setOnHold1 && set1[setIndex1] !== lcs[lcsIndex]) {
        supersequence.push(set1[setIndex1])
        setIndex1 += 1
      } else {
        setOnHold1 = true
      }
    }

    // Add elements of the second set to supersequence in correct order.
    if (setIndex2 < set2.length) {
      if (!setOnHold2 && set2[setIndex2] !== lcs[lcsIndex]) {
        supersequence.push(set2[setIndex2])
        setIndex2 += 1
      } else {
        setOnHold2 = true
      }
    }

    // Add LCS element to the supersequence in correct order.
    if (setOnHold1 && setOnHold2) {
      supersequence.push(lcs[lcsIndex])
      lcsIndex += 1
      setIndex1 += 1
      setIndex2 += 1
      setOnHold1 = false
      setOnHold2 = false
    }
  }

  // Attach set1 leftovers.
  if (setIndex1 < set1.length) {
    supersequence = supersequence.concat(set1.slice(setIndex1))
  }

  // Attach set2 leftovers.
  if (setIndex2 < set2.length) {
    supersequence = supersequence.concat(set2.slice(setIndex2))
  }

  return supersequence.join('')
}

let res = shortestCommonSupersequence('abac', 'cab')
console.log(res)
