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
// const copyLeft = new TreeNode(2)
// const copyHead = new TreeNode(1,copyLeft, copyRight)

const isSameTree = (primary, secondary) => {
  if (!primary || !secondary) return !primary && !secondary;
  if (primary.val !== secondary.val) return false;
  return (
    isSameTree(primary.left, secondary.left) &&
    isSameTree(primary.right, secondary.right)
  );
};

const isSubTree = (root, subRoot) => {
  const findSubtree = (r, s) => {
    if (r == null) return false;
    if (isSameTree(r, s)) return true;
    return findSubtree(r.left, s) || findSubtree(r.right, s);
  };

  return findSubtree(root, subRoot);
};

console.log(isSubTree(head, copyRight));
