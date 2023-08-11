function TreeNode(val, left, right) {
  this.val = val;
  this.left = left ?? null;
  this.right = right ?? null;
}

const right1 = new TreeNode(5);
const left1 = new TreeNode(4);
const right = new TreeNode(3, left1, right1);
const left = new TreeNode(2);
const head = new TreeNode(1, left, right);

const isBalanced = (root) => {
  if (root == null) return true;

  let balanced = true;

  const dfs = (node, level) => {
    if (node == null) return level;

    const heightLeft = dfs(node?.left, level + 1);
    const heightRight = dfs(node?.right, level + 1);

    if (Math.abs(heightLeft - heightRight) > 1) {
      balanced = false;
    }

    return Math.max(heightRight, heightLeft);
  };

  dfs(root, 0);

  return balanced;
};

console.log(isBalanced(head));
