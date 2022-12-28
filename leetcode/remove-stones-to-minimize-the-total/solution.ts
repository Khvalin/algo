class Heap {
  a = [] as number[];
  less = (a, b) => a < b;
  constructor(less) {
    if (less) {
      this.less = less;
    }
  }

  push(n) {
    this.a.push(n);
    return this.siftUp(this.a.length - 1);
  }

  siftUp(ind) {
    const a = this.a;
    let i = ind;
    while (i > 0 && !this.less(a[(i - 1) >> 1], a[i])) {
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

      if (r < a.length && this.less(a[r], a[minInd])) {
        minInd = r;
      }

      if (l < a.length && this.less(a[l], a[minInd])) {
        minInd = l;
      }

      [a[minInd], a[ind]] = [a[ind], a[minInd]];
    } while (minInd !== ind);
  }

  size() {
    return this.a.length;
  }
}

function minStoneSum(piles: number[], k: number): number {
  const heap = new Heap((a, b) => {
    return (a > b);
  });

  for (const n of piles) {
    heap.push(n);
  }

  for (let i = 0; i < k; i++) {
    const n = heap.pop();
    heap.push(n >> 1 + n % 1);
  }

  let res = 0;
  for (const n of heap.a) {
    res += n;
  }

  return res;
}
