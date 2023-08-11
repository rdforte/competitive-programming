module.exports.isValidBST = (root) => dfsBST(root);

const dfsBST = (
  node,
  min = Number.NEGATIVE_INFINITY,
  max = Number.POSITIVE_INFINITY
) => {
  if (node == null) return true;

  if (!(node.val > min && node.val < max)) return false;

  const leftIsValid = dfsBST(node.left, min, node.val);
  const rightIsValid = dfsBST(node.right, node.val, max);

  return leftIsValid && rightIsValid;
};

module.exports.createNode = (val, left, right) => new Node(val, left, right);

class Node {
  val;
  left;
  right;

  constructor(val, left, right) {
    this.val = val;
    this.left = left;
    this.right = right;
  }
}
