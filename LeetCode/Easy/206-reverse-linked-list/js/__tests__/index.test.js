const { reverseList, createNode } = require("../index");

it("should reverse the linked list", () => {
  const node5 = createNode(5, null);
  const node4 = createNode(4, node5);
  const node3 = createNode(3, node4);
  const node2 = createNode(2, node3);
  const node1 = createNode(1, node2);

  expect(reverseList(node1).value).toEqual(node5.value);
});
