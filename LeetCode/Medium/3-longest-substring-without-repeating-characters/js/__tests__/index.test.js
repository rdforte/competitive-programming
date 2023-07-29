const { lengthOfLongestSubstring } = require("../index");

it.each([
  {
    input: "abcabcbb",
    expectedOutput: 3,
  },
  {
    input: "bbbbb",
    expectedOutput: 1,
  },
  {
    input: "pwwkew",
    expectedOutput: 3,
  },
  {
    input: "abba",
    expectedOutput: 2,
  },
])(
  "should find longest substring without repeating characters",
  ({ input, expectedOutput }) => {
    const substingLen = lengthOfLongestSubstring(input);
    expect(substingLen).toEqual(expectedOutput);
  }
);
