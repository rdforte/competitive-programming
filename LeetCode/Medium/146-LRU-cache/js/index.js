module.exports.createLRUCache = (capacity) => new LRUCache(capacity);

/**
 * @param {number} capacity
 */
var LRUCache = function (capacity) {
  this.capacity = capacity;
  this.cache = new Map();
  this.queue = new Queue();
};

/**
 * @param {number} key
 * @return {number}
 */
LRUCache.prototype.get = function (key) {
  if (this.cache.has(key)) {
    const node = this.cache.get(key);
    this.queue.remove(node);
    const newNode = this.queue.add(key, node.value);
    this.cache.set(key, newNode);
    return newNode.value;
  }

  return -1;
};

/**
 * @param {number} key
 * @param {number} value
 * @return {void}
 */
LRUCache.prototype.put = function (key, value) {
  if (this.cache.has(key)) {
    const node = this.cache.get(key);
    this.queue.remove(node);
    const newNode = this.queue.add(key, value);
    this.cache.set(key, newNode);
    return;
  }

  if (this.queue.size() === this.capacity) {
    const lruNode = this.queue.top();
    this.cache.delete(lruNode.key);
    this.queue.remove(lruNode);
  }

  const node = this.queue.add(key, value);
  this.cache.set(key, node);
};

/**
 * Your LRUCache object will be instantiated and called as such:
 * var obj = new LRUCache(capacity)
 * var param_1 = obj.get(key)
 * obj.put(key,value)
 */

class Queue {
  #front;
  #back;
  #length = 0;

  constructor() {
    this.#front = new LLNode(-1, -1);
    this.#back = new LLNode(-1, -1);
    this.#front.next = this.#back;
    this.#back.prev = this.#front;
  }

  add(key, value) {
    const node = new LLNode(key, value);
    node.prev = this.#back.prev;
    node.next = this.#back;
    this.#back.prev.next = node;
    this.#back.prev = node;

    this.#length++;

    return node;
  }

  size() {
    return this.#length;
  }

  remove(node) {
    node.prev.next = node.next;
    node.next.prev = node.prev;
    this.#length = Math.max(0, this.#length - 1);
  }

  top() {
    return this.#front.next;
  }
}

function LLNode(key, value, next, prev) {
  this.value = value;
  this.key = key;
  this.next = next ?? null;
  this.prev = prev ?? null;
}
