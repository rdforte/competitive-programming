const { searchMatrix } = require("../index");

it.each([
  {
    matrix: [
      [1, 3, 5, 7],
      [10, 11, 16, 20],
      [23, 30, 34, 60],
    ],
    target: 10,
    expectedOutput: true,
  },
  {
    matrix: [
      [1, 3, 5, 7],
      [10, 11, 16, 20],
      [23, 30, 34, 60],
    ],
    target: 13,
    expectedOutput: false,
  },
])(
  "should determine whether target is in matrix",
  ({ matrix, target, expectedOutput }) => {
    const output = searchMatrix(matrix, target);
    expect(output).toEqual(expectedOutput);
  }
);
