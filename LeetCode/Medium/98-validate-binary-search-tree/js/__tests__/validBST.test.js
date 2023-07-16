const { isValidBST, createNode } = require("../validBST.js");

it("should return true if valid Binary Search Tree", () => {
  const seven = createNode(7);
  const five = createNode(5);
  const three = createNode(3);
  const one = createNode(1);
  const six = createNode(6, five, seven);
  const two = createNode(2, one, three);
  const root = createNode(4, two, six);

  const isValidTree = isValidBST(root);
  expect(isValidTree).toBe(true);
});
