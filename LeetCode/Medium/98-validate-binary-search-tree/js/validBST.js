module.exports.isValidBST = (root) => {
  return validateBSTRecursively(
    root,
    Number.NEGATIVE_INFINITY,
    Number.POSITIVE_INFINITY
  );
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
// min = 6
// max = Infinity
const validateBSTRecursively = (node, min, max) => {
  if (node == null) return true;

  let isValid = node.val > min && node.val < max;

  if (!isValid) return false;

  if (!validateBSTRecursively(node.left, min, node.val)) {
    isValid = false;
  }

  if (!validateBSTRecursively(node.right, node.val, max)) {
    isValid = false;
  }

  return isValid;
};
