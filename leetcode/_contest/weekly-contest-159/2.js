/**
 * @param {string[]} folder
 * @return {string[]}
 */
var removeSubfolders = function(folder) {
  folder = folder.sort((a, b) => a.localeCompare(b))

  const res = []
  let i = 0
  while (i < folder.length) {
    const f = folder[i]
    res.push(f)

    i++
    while (i < folder.length && folder[i].startsWith(f + '/')) {
      i++
    }
  }

  return res
}

console.log(removeSubfolders([ '/a', '/a/b', '/c/d', '/c/d/e', '/c/f' ]))
console.log(removeSubfolders([ '/a', '/a/b/c', '/a/b/d' ]))
