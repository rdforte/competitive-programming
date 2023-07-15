const {
    createLLNode,
    swapEverySecondNode,
} = require("../index.js");

it("should swap every two adjacent nodes in linked list", () => {
    const nodeFour = createLLNode(4, null);
    const nodeThree = createLLNode(3, nodeFour);
    const nodeTwo = createLLNode(2, nodeThree);
    const nodeOne = createLLNode(1, nodeTwo);

    const nodeOneInList = swapEverySecondNode(nodeOne);

expect(nodeOneInList?.value).toEqual(2);
    const nodeTwoInList = nodeOneInList.next;
    expect(nodeTwoInList?.value).toEqual(1);
    const nodeThreeInList = nodeTwoInList.next;
    expect(nodeThreeInList?.value).toEqual(4);
    const nodeFourInList = nodeThreeInList.next;
    expect(nodeFourInList?.value).toEqual(3);
});