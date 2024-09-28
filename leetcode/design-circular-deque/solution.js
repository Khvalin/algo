/**
 * @param {number} k
 */
var MyCircularDeque = function (k) {
    this.items = new Map()
    this.first = 0
    this.len = 0
    this.maxLen = k
    this._isEmpty = true
};

/** 
 * @param {number} value
 * @return {boolean}
 */
MyCircularDeque.prototype.insertFront = function (value) {
    if (this.isFull()) {
        return false
    }

    this.first--
    this.len++

    this.items.set(this.first, value)

    return true
};

/** 
 * @param {number} value
 * @return {boolean}
 */
MyCircularDeque.prototype.insertLast = function (value) {
    if (this.isFull()) {
        return false
    }


    this.items.set(this.first + this.len, value)
    this.len++
    return true
};

/**
 * @return {boolean}
 */
MyCircularDeque.prototype.deleteFront = function () {
    if (this.isEmpty()) {
        return false
    }
    this.first++
    this.len--

    return true
};

/**
 * @return {boolean}
 */
MyCircularDeque.prototype.deleteLast = function () {
    if (this.isEmpty()) {
        return false
    }
    this.len--

    return true
};

/**
 * @return {number}
 */
MyCircularDeque.prototype.getFront = function () {
    if (this.isEmpty()) {
        return -1
    }
    return this.items.get(this.first)
};

/**
 * @return {number}
 */
MyCircularDeque.prototype.getRear = function () {
    if (this.isEmpty()) {
        return -1
    }
    return this.items.get(this.first + this.len - 1)
};

/**
 * @return {boolean}
 */
MyCircularDeque.prototype.isEmpty = function () {
    return this.len == 0
};

/**
 * @return {boolean}
 */
MyCircularDeque.prototype.isFull = function () {
    return this.len === this.maxLen
};

/** 
 * Your MyCircularDeque object will be instantiated and called as such:
 * var obj = new MyCircularDeque(k)
 * var param_1 = obj.insertFront(value)
 * var param_2 = obj.insertLast(value)
 * var param_3 = obj.deleteFront()
 * var param_4 = obj.deleteLast()
 * var param_5 = obj.getFront()
 * var param_6 = obj.getRear()
 * var param_7 = obj.isEmpty()
 * var param_8 = obj.isFull()
 */