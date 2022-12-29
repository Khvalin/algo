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

function getOrder(tasks: number[][]): number[] {
  const indices = tasks.map((_, i) => i);
  indices.sort((i, j) => {
    if (tasks[i][0] === tasks[j][0]) {
      return tasks[i][1] - tasks[j][1];
    }

    return tasks[i][0] - tasks[j][0];
  });

  const heap = new Heap((i, j) => {
    if (tasks[i][1] === tasks[j][1]) {
      return i < j;
    }
    return tasks[i][1] < tasks[j][1];
  });
  heap.push(indices[0]);

  const res = [];
  let i = 1;
  let t = 0;
  while (res.length < tasks.length) {
    const j = heap.pop();
    res.push(j);
    t = Math.max(t, tasks[j][0]) + tasks[j][1];

    while (
      i < tasks.length && (tasks[indices[i]][0] <= t || heap.size() === 0)
    ) {
      heap.push(indices[i]);
      i++;
    }
  }

  return res;
}

let input = [];
//input =   [[1,2],[2,4],[3,2],[4,1]]
// deno-fmt-ignore
input = [[19,13],[16,9],[21,10],[32,25],[37,4],[49,24],[2,15],[38,41],[37,34],[33,6],[45,4],[18,18],[46,39],[12,24]]
let res = getOrder(input);
console.log(res);
