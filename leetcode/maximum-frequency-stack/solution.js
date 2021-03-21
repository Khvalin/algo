var FreqStack = function() {
  this.head = { val: '', prev: null, next: null };
  this.tail = this.head;
  this.counts = new Map();
  this.counts.set('', -1);
};

/** 
* @param {number} x
* @return {void}
*/
FreqStack.prototype.push = function(x) {
  const count = (this.counts.has(x) ? this.counts.get(x) : 0) + 1;
  const node = { val: x, next: null, count };
  this.counts.set(x, count);

  if (this.counts.get(this.tail.val) <= count) {
    node.prev = this.tail;
    this.tail.next = node;
    this.tail = node;

    return;
  }

  let cur = this.head;
  do {
    cur = cur.next;
  } while (cur.count <= count);

  node.prev = cur.prev;
  node.next = cur;
  cur.prev.next = node;
  cur.prev = node;
};

/**
* @return {number}
*/
FreqStack.prototype.pop = function() {
  // let cur = this.head;
  // const a = [];
  // while (cur) {
  //   a.push(cur.val);
  //   cur = cur.next;
  // }
  // console.log(a);
  const res = this.tail.val;
  this.counts.set(res, this.counts.get(res) - 1);

  this.tail = this.tail.prev;
  this.tail.next = null;

  return res;
};

const obj = new FreqStack();
obj.push(5);
obj.push(7);
obj.push(5);
obj.push(7);
obj.push(4);
obj.push(5);
console.log(obj.pop());
console.log(obj.pop());
console.log(obj.pop());
console.log(obj.pop());

/** 
* Your FreqStack object will be instantiated and called as such:
* var obj = new FreqStack()
* obj.push(x)
* var param_2 = obj.pop()
*/
