function isValid(stale, latest, otjson) {
  let pos = 0
  const text = [...stale]

  const operations = JSON.parse(otjson) // assuming valid JSON

  for (const { op, count, chars } of operations) {
    switch (op) {
      case 'skip':
        pos += count;
        if (pos < 0 || pos > text.length) {
          return false
        }
        break;

      case 'delete':
        if (count < 0 || pos + count > text.length) {
          return false
        }
        text.splice(pos, count)
        break;

      case 'insert':
        if (typeof chars !== 'string') {
          return false
        }
        text.splice(pos, 0, ...chars)
        pos += chars.length
        break;

      default:
        return false
    }
  }

  if (text.length !== latest.length) {
    return false
  }

  for (let i = 0; i < text.length; i++) {
    if (text[i] !== latest[i]) {
      return false
    }
  }

  return true
}
let res = ''

res = isValid(
  '',
  'Hello, human!',
  '[{"op": "insert", "chars": "Hello, human!"}, {"op": "insert", "chars": " Hello, human!"}]'
); // true
console.log(res, true)

res = isValid(
  'a',
  '',
  '[{"op": "skip", "count": 0}, {"op": "delete", "count": 1}]'
); // true
console.log(res, true)

res = isValid(
  'Repl.it uses operational transformations to keep everyone in a multiplayer repl in sync.',
  'Repl.it uses operational transformations.',
  '[{"op": "skip", "count": 40}, {"op": "delete", "count": 47}]'
); // true
console.log(res, true)

res = isValid(
  'Repl.it uses operational transformations to keep everyone in a multiplayer repl in sync.',
  'Repl.it uses operational transformations.',
  '[{"op": "skip", "count": 45}, {"op": "delete", "count": 47}]'
); // false, delete past end
console.log(res, false)

res = isValid(
  'Repl.it uses operational transformations to keep everyone in a multiplayer repl in sync.',
  'Repl.it uses operational transformations.',
  '[{"op": "skip", "count": 40}, {"op": "delete", "count": 47}, {"op": "skip", "count": 2}]'
); // false, skip past end
console.log(res, false)

res = isValid(
  'Repl.it uses operational transformations to keep everyone in a multiplayer repl in sync.',
  'We use operational transformations to keep everyone in a multiplayer repl in sync.',
  '[{"op": "delete", "count": 7}, {"op": "insert", "chars": "We"}, {"op": "skip", "count": 4}, {"op": "delete", "count": 1}]'
); // true
console.log(res, true)

res = isValid(
  'Repl.it uses operational transformations to keep everyone in a multiplayer repl in sync.',
  'We can use operational transformations to keep everyone in a multiplayer repl in sync.',
  '[{"op": "delete", "count": 7}, {"op": "insert", "chars": "We"}, {"op": "skip", "count": 4}, {"op": "delete", "count": 1}]'
); // false
console.log(res, false)

res = isValid(
  'Repl.it uses operational transformations to keep everyone in a multiplayer repl in sync.',
  'Repl.it uses operational transformations to keep everyone in a multiplayer repl in sync.',
  '[]'
); // true
console.log(res, true)
