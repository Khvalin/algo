class Heap {
  a = [] as number[];
  constructor() {
  }

  push(n) {
    this.a.push(n);
    return this.siftUp(this.a.length - 1);
  }

  siftUp(ind) {
    const a = this.a;
    let i = ind;
    while (i > 0 && a[(i - 1) >> 1] > a[i]) {
      const p = (i - 1) >> 1;
      [a[p], a[i]] = [a[i], a[p]];
      i = p;
    }
  }

  pop() {
    const a = this.a;
    if (!a.length) {
      return null;
    }

    const res = a[0];
    let last = a.pop();
    if (a.length) {
      a[0] = last;
      this.siftDown(0);
    }
    return res;
  }

  siftDown(ind) {
    const a = this.a;

    let minInd = ind;
    do {
      ind = minInd;
      const l = (ind << 1) + 1;
      const r = (ind << 1) + 2;

      minInd = ind;

      if (r < a.length && a[r] < a[minInd]) {
        minInd = r;
      }

      if (l < a.length && a[l] < a[minInd]) {
        minInd = l;
      }

      [a[minInd], a[ind]] = [a[ind], a[minInd]];
    } while (minInd !== ind);
  }

  size() {
    return this.a.length;
  }
}

function connectSticks(sticks: number[]): number {
  const heap = new Heap();
  for (const n of sticks) {
    heap.push(n);
  }

  if (sticks.length < 2) {
    return 0;
  }

  let res = 0;
  while (heap.size() > 1) {
    const a = heap.pop();
    const b = heap.pop();
    const s = a + b;

    res += s;
    heap.push(s);
  }

  return res;
}

let res;
//res = connectSticks([6,5,4,3,2,1])
res = connectSticks([
  3354,
  4316,
  3259,
  4904,
  4598,
  474,
  3166,
  6322,
  8080,
  9009,
]);
console.log(res, 151646);
//res = connectSticks([5])
//res = connectSticks([1,4, 3, 5, 5, 3])
//console.log(res)
