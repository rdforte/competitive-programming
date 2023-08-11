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

const copyRight1 = new TreeNode(5);
const copyLeft1 = new TreeNode(4);
const copyRight = new TreeNode(3, copyLeft1, copyRight1);
const copyLeft = new TreeNode(2);
const copyHead = new TreeNode(1, copyLeft, copyRight);

const isSameTree = (primary, secondary) => {
  let sameTree = true;

  const dfsSameTree = (p, s) => {
    if (p == null && s == null) {
      return;
    }

    if (
      (p == null && s != null) ||
      (p != null && s == null) ||
      p?.val != s?.val
    ) {
      sameTree = false;
      return;
    }

    if (sameTree) dfsSameTree(p.left, s.left);
    if (sameTree) dfsSameTree(p.right, s.right);
  };

  dfsSameTree(primary, secondary);

  return sameTree;
};

console.log(isSameTree(head, copyHead));
