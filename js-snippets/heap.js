class Heap {
  arr = [];
  count = 0;
  less = (a, b) => a < b;

  constructor(size, less) {
    this.arr = Array.from(Array(size));
    if (less) {
      this.less = less;
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
    let i = this.count++;
    this.arr[i] = val;
    while (i != 0 && this.less(this.arr[this.parent(i)], this.arr[i])) {
      const parent = this.parent(i);

      // Swap child node with parent
      const temp = this.arr[parent];
      this.arr[parent] = this.arr[i];
      this.arr[i] = temp;

      i = parent;
    }
  }

  peek() {
    return  this.arr[0]
  }

  pop() {
    const res = this.arr[0]
    this.arr[0] = this.arr[this.arr.length - 1];
    this.count--;
    this.arr.pop();
    this.heapify();

    return res
  }

  heapify(i = 0) {
    if (this.count < 2) {
      return;
    }

    while (i < this.count) {
      //For empty or Heap with single element we need not perform any operation

      let largest = i;
      let left = this.leftChild(i);
      let right = this.rightChild(i);

      if (left < this.count && this.less(this.arr[i], this.arr[left])) {
        largest = left;
      }
      if (right < this.count && this.less(this.arr[right], this.arr[largest])) {
        largest = right;
      }

      if (largest === i) {
        break;
      }

      // Swap root node with larger child
      const temp = this.arr[largest];
      this.arr[largest] = this.arr[i];
      this.arr[i] = temp;

      // Heapify child nodes to re-order subtree and maintain MinHeap integrity
      i = largest;
    }
  }
}
