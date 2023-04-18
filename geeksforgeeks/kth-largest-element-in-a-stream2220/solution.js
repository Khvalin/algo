
class Heap {
  arr = [];
  #less = (a, b) => a < b;
  get count() {
	  return this.arr.length
  };

  constructor(size, less) {
    this.arr = Array.from(Array(size));
    if (less) {
      this.#less = less;
    }
  }

  leftChild(x) {
    return 2 * x + 1;
  }

  rightChild(x) {
    return this.leftChild(x) + 1;
  }

  parent(x) {
    return (x - 1) >> 1;
  }

  insert(val) {
	let i = this.count;
    this.arr.push(val)

    while (i != 0 && this.#less(this.arr[this.parent(i)], this.arr[i])) {
      const parent = this.parent(i);

      // Swap child node with parent
      const temp = this.arr[parent];
      this.arr[parent] = this.arr[i];
      this.arr[i] = temp;

      i = parent;
    }
  }

  peek() {
    return this.arr[0]
  }

  pop() {
    const res = this.arr[0]
    this.arr[0] = this.arr[this.count - 1];
    this.arr.pop();
    this.heapify();

    return res
  }

  heapify(i = 0) {
    if (this.count < 2) {
	  //For empty or Heap with single element we need not perform any operation
      return;
    }

    while (i < this.count) {
      let largest = i;
      let left = this.leftChild(i);
      let right = this.rightChild(i);

      if (left < this.count && this.#less(this.arr[i], this.arr[left])) {
        largest = left;
      }
      if (right < this.count && this.#less(this.arr[largest], this.arr[right])) {
        largest = right;
      }

      if (largest === i) {
        break;
      }

      // Swap root node with larger child
      const temp = this.arr[largest];
      this.arr[largest] = this.arr[i];
      this.arr[i] = temp;

      // Heapify child nodes to re-order subtree and maintain Heap integrity
      i = largest;
    }
  }
}



class Solution {
    kthLargest(arr,k,n){
        const stack = []
        const heap = new Heap(0, (a, b) => a > b)
        const res = []
        for (const num of arr) {
            heap.insert(num)
            if (heap.count < k) {
                res.push(-1)
                continue
            }
            console.log(heap)
            if (heap.count >= k) {
                heap.pop()
            }
            heap.heapify()
            res.push(heap.peek())
        }

        return res
    }
}

const sol = new Solution()
const res = sol.kthLargest([1,2,3,4,5,6,7,8,7,7], 4, 6)
console.log(res)

