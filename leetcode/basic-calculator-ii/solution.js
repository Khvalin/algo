/**
 * @param {string} s
 * @return {number}
 */
var calculate = function (s) {
  let n = null;

  const arr = [];

  const calc = (op, a, b) => {
    switch (op) {
      case ("-"):
        return a - b;
      case ("+"):
        return a + b;
      case ("/"):
        return Math.floor(a / b);
      case ("*"):
        return a * b;
    }
  };

  const calcTop = () => {
    const b = arr.pop();
    const op = arr.pop();
    const a = arr.pop();

    arr.push(calc(op, a, b));
  };

  const calcBottom = () => {
    const [a, op] = arr.splice(0, 2);
    const b = arr[0];

    arr[0] = calc(op, a, b);
  };

  for (let i = 0; i <= s.length; i++) {
    const ch = s[i] ?? "~";
    const d = 1 * ch;

    if (!Number.isNaN(d)) {
      n = 10 * (n || 0) + d;
      continue;
    }

    if (n !== null) {
      arr.push(n);
      n = null;
    }

    if (
      arr.length > 2 &&
      (arr[arr.length - 2] === "*" || arr[arr.length - 2] === "/")
    ) {
      // a + b * c
      calcTop();
    }

    if (arr.length >= 7 && (arr[1] === "+" || arr[1] === "-")) {
      // a + b + c * d
      calcBottom();
    }

    switch (ch) {
      case ("-"):
      case ("+"):
      case ("/"):
      case ("*"):
        arr.push(ch);
        break;
    }
  }

  while (arr.length > 2) {
    calcBottom();
  }

  return arr[0];
};

//console.log(calculate(" 3+5 / 2 "))
//console.log(calculate(" 1-1 + 1"))
//console.log(calculate("1-1+1"))
console.log(calculate("   3 / 2 "));
