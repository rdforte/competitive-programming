const { solveNQueens } = require("../index");

it.each([
  {
    boardDimensions: 4,
    expectedOutput: [
      [".Q..", "...Q", "Q...", "..Q."],
      ["..Q.", "Q...", "...Q", ".Q.."],
    ],
    expectedOutputLen: 2,
  },
])(
  "find all distinct positions can place queen on $boardDimensions x $boardDimensions board",
  ({ boardDimensions, expectedOutput, expectedOutputLen }) => {
    const numPermutations = solveNQueens(boardDimensions);
    expect(numPermutations.length).toEqual(expectedOutputLen);
    expect(numPermutations).toEqual(expect.arrayContaining(expectedOutput));
  }
);
