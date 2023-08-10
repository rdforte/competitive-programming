module.exports.reverseList = (head) => {
  let prevNode = null;

  while (head != null) {
    const nextHead = head.next;
    head.next = prevNode;
    prevNode = head;
    head = nextHead;
  }

  return prevNode;
};

module.exports.createNode = (value, next) => new node(value, next);

function node(value, next) {
  this.value = value ?? 0;
  this.next = next ?? null;
}
