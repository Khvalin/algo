/**
 * Initialize your data structure here.
 */
var MyQueue = function() {
  this.stacks = [];
  this.dirty = false;
};

/**
 * Push element x to the back of queue. 
 * @param {number} x
 * @return {void}
 */
MyQueue.prototype.push = function(x) {
  if (!this.dirty) {
    this.stacks.push([])
    this.dirty = true;
  }
  this.stacks[this.stacks.length - 1].push(x)
};

/**
 * Removes the element from in front of queue and returns that element.
 * @return {number}
 */
MyQueue.prototype.pop = function() {
  if (this.dirty){
    this.peek()
  }

  let l = this.stacks.length;
  const last = this.stacks[l - 1];
  const r = last.pop()
  
  while(l > 0 && this.stacks[l-1].length === 0) {
    this.stacks.pop()
    l--
  }
  
  return r
};

/**
 * Get the front element.
 * @return {number}
 */
MyQueue.prototype.peek = function() {
  let l = this.stacks.length;
  while(l > 0 && this.stacks[l-1].length === 0) {
    this.stacks.pop()
    l--
  }
  if (l === 0) {
    return
  }

  if (this.dirty) {
    this.dirty = false;
    const temp = [];
    while (this.stacks[l-1].length > 0){
      const top = this.stacks[l-1].pop()
      temp.push(top)
    }

    this.stacks.pop()
    const f = this.stacks.pop()
    this.stacks.push(temp)
    if (f) {
      this.stacks.push(f)
    }
  }
  const last = this.stacks[l - 1];
  
  return last[last.length - 1]
};

/**
 * Returns whether the queue is empty.
 * @return {boolean}
 */
MyQueue.prototype.empty = function() {
  return this.stacks.length === 0 || this.stacks[0].length === 0
};

/** 
 * Your MyQueue object will be instantiated and called as such:
 * var obj = new MyQueue()
 * obj.push(x)
 * var param_2 = obj.pop()
 * var param_3 = obj.peek()
 * var param_4 = obj.empty()
 */
