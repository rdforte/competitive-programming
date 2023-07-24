const { threeSum } = require("../index");

it.each([
  {
    input: [-1, 0, 1, 2, -1, -4],
    expectOutput: [
      [-1, -1, 2],
      [-1, 0, 1],
    ],
  },
  {
    input: [0, 1, 1],
    expectOutput: [],
  },
  {
    input: [0, 0, 0],
    expectOutput: [[0, 0, 0]],
  },
  {
    input: [0, 0, 0, 0],
    expectOutput: [[0, 0, 0]],
  },
  {
    input: [-2, 0, 0, 2, 2],
    expectOutput: [[-2, 0, 2]],
  },
])(
  "should return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0",
  ({ input, expectOutput }) => {
    const output = threeSum(input);
    expect(output.length).toEqual(expectOutput.length);
    const sortedOutput = output.map((o) => o.sort());
    expect(sortedOutput).toEqual(expect.arrayContaining(expectOutput));
  }
);
