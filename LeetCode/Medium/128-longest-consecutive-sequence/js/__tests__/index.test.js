const { longestConsecutiveSequence } = require("../index");

it.each([
  {
    input: [100, 4, 200, 1, 3, 2],
    expectedOutput: 4,
  },
  {
    input: [0, 3, 7, 2, 5, 8, 4, 6, 0, 1],
    expectedOutput: 9,
  },
])(
  "should find the longest consecutive output $input",
  ({ input, expectedOutput }) => {
    expect(longestConsecutiveSequence(input)).toEqual(expectedOutput);
  }
);
