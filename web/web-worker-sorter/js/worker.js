const LEN = 1e5;
const a = new Array(LEN);
const CHUNK_SIZE = LEN / 20;

for (let i = 0; i < a.length; i++) {
  a[i] = Math.floor(1e6 * Math.random());
}

let sortingIntervalId = null;

const C = [0, 0];

const newNums = [];

const sort = () => {
  let k = 0;
  while (C[0] < a.length) {
    if (k >= CHUNK_SIZE) {
      break;
    }
    k++;
    let i = C[0];
    C[1] = i + 1;
    while (C[1] < a.length) {
      const j = C[1];
      if (a[j] < a[i]) {
        i = j;
      }
      C[1]++;
    }

    const t = a[C[0]];
    a[C[0]] = a[i];
    a[i] = t;

    if (C[0] % Math.floor(LEN / 34) === 0) {
      postMessage({ action: "onprogress", complete: C[0], total: a.length });
    }
    C[0]++;
  }

  if (C[0] >= a.length) {
    postMessage({ action: "done", array: a });
    return true;
  }

  return false;
};

const continueSort = (doSort) => {
  if (doSort) {
    const sorted = sort();

    if (!sorted) {
      sortingIntervalId = setTimeout(continueSort, 0, true);
    }

    return;
  }

  sortingIntervalId = setTimeout(continueSort, 0, true);
};

onmessage = function (e) {
  console.log("Worker: action data: ", JSON.stringify(e.data));

  if (e.data.action === "sort") {
    continueSort(true);
  }

  if (e.data.action === "push") {
    newNums.push(e.data.num);
  }

  // const result = e.data[0] * e.data[1];
  // if (isNaN(result)) {
  //   postMessage('Please write two numbers');
  // } else {
  //   const workerResult = 'Result: ' + result;
  //   console.log('Worker: Posting message back to main script');
  //   postMessage(workerResult);
  // }
};
