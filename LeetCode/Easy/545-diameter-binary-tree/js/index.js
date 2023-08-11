function TreeNode(val, left, right) {
  this.val = val;
  this.left = left ?? null;
  this.right = right ?? null;
}

const right1 = new TreeNode(5);
const left1 = new TreeNode(4);
const right = new TreeNode(3);
const left = new TreeNode(2, left1, right1);
const head = new TreeNode(1, left, right);

// used a closure to keep track of the max length between 2 nodes. Could have also used a class

const diameterOfBinaryTree = (node) => {
  let max = 0;
  const findDiameterOfBinaryTree = (n) => {
    if (n === null) return 0;

    const leftMostNodeLevel = findDiameterOfBinaryTree(n.left);
    const rightMostNodeLevel = findDiameterOfBinaryTree(n.right);

    const diameter = leftMostNodeLevel + rightMostNodeLevel;
    max = Math.max(max, diameter);
    return Math.max(leftMostNodeLevel, rightMostNodeLevel) + 1;
  };

  findDiameterOfBinaryTree(node);

  return max;
};

console.log(diameterOfBinaryTree(head));
