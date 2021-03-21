function mobileRemote(text) {
  const abc = [ '', '1', '2abc', '3def', '4ghi', '5jkl', '6mno', '7pqrs', '8tuv', '9wxyz', '*', '0', '#' ];
  const cost = [
    [],
    [ 0, 0, 1, 1, 1, 2, 2, 2, 3, 3, 1, 2, 2 ],
    [ 0, 1, 0, 1, 2, 1, 2, 3, 2, 3, 2, 1, 2 ],
    [ 0, 1, 1, 0, 2, 2, 1, 3, 4, 2, 2, 2, 1 ],
    [ 0, 1, 2, 2, 0, 1, 1, 1, 2, 2, 2, 3, 3 ],
    [ 0, 2, 1, 2, 1, 0, 1, 2, 1, 2, 3, 2, 3 ],
    [ 0, 2, 2, 1, 1, 1, 0, 2, 3, 1, 3, 3, 2 ],
    [ 0, 2, 3, 3, 1, 2, 2, 0, 1, 2, 1, 2, 2 ],
    [ 0, 3, 2, 3, 2, 1, 2, 1, 0, 1, 2, 1, 2 ],
    [ 0, 3, 3, 2, 2, 2, 1, 1, 1, 0, 2, 2, 1 ],
    [ 0, 1, 2, 2, 2, 3, 2, 1, 2, 2, 0, 1, 1 ],
    [ 0, 2, 1, 2, 3, 2, 3, 2, 1, 2, 1, 0, 1 ],
    [ 0, 2, 2, 1, 3, 3, 2, 3, 2, 1, 1, 1, 0 ]
  ];
  let res = 0;
  let caps = false;
  let cur = 1;

  let i = 0;
  while (i < text.length) {
    const ch = text[i];
    let needCaseChange = false;
    if (ch.toLowerCase() !== ch.toUpperCase()) {
      needCaseChange = caps ? ch === ch.toLowerCase() : ch === ch.toUpperCase();
    }
    if (needCaseChange) {
      res += cost[cur][10] + 2;
      cur = 10;
      caps = !caps;
      continue;
    }

    let next = -1;
    let pos = -1;
    for (let j = 1; j < abc.length; j++) {
      pos = abc[j].indexOf(ch.toLowerCase());
      if (pos > -1) {
        next = j;
        break;
      }
    }
    console.log(ch, cur, next, pos);

    res += cost[cur][next] + 1 + pos + 1;
    cur = next;
    i++;
  }

  return res;
}

const ref = {
  mobileremote: 71
  // c: 6,
  // yandex: 34,

  // MobileRemote: 87,
  // Yandex: 39,
  // C: 10,
  // YANDEX: 36,
  // MOBILEREMOTE: 75,
  // Mobile111Remote: 93,
  // Yandex77: 44,
  // '12345': 15
};

for (const [ key, value ] of Object.entries(ref)) {
  console.log(key, mobileRemote(key), value);
}
