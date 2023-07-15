class LLNode {
    value;
    next;
    constructor(value, next) {
        this.value = value;
        this.next = next;
    }
}
module.exports.createLLNode = (value, next) => new LLNode(value, next);

module.exports.swapEverySecondNode = (node) => {
    return swapSecondNode(node);
};

const swapSecondNode = (node) => {
    if (node == null || node.next == null) {
        return node;
    }

    const nextNode = node.next;
    const nextNextNode = nextNode.next;
    nextNode.next = node;

    node.next = swapSecondNode(nextNextNode);
    return nextNode;
}