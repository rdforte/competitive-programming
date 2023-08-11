function TreeNode(val, left, right) {
  this.val = val === undefined ? 0 : val;
  this.left = left === undefined ? null : left;
  this.right = right === undefined ? null : right;
}

const right1 = new TreeNode(7);
const left1 = new TreeNode(15);
const right = new TreeNode(20, left1, right1);
const left = new TreeNode(9);
const head = new TreeNode(3, left, right);

const maxDepth = (root, level = 0) => {
  if (root === null) {
    return level;
  }

  return Math.max(
    maxDepth(root.left, level + 1),
    maxDepth(root.right, level + 1)
  );
};

console.log(maxDepth(head));
