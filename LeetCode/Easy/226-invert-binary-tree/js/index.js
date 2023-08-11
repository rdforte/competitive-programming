function TreeNode(val, left, right) {
  this.val = val;
  this.left = left ?? null;
  this.right = right ?? null;
}

const right3 = new TreeNode(9);
const left3 = new TreeNode(6);
const right2 = new TreeNode(3);
const left2 = new TreeNode(1);
const right1 = new TreeNode(7, left3, right3);
const left1 = new TreeNode(2, left2, right2);
const head = new TreeNode(4, left1, right1);

// Utilise BFS

const invertTree = (root) => {
  if (root == null) return root;
  const queue = [];
  queue.push(root);

  while (queue.length > 0) {
    const levelLength = queue.length;

    for (let i = 0; i < levelLength; i++) {
      const node = queue.pop();
      const left = node.left;
      const right = node.right;
      node.left = right;
      node.right = left;

      if (node.left) {
        queue.unshift(node.left);
      }

      if (node.right) {
        queue.unshift(node.right);
      }
    }
  }

  return root;
};

invertTree(head);

console.log("left", head.left);
console.log("right", head.right);
