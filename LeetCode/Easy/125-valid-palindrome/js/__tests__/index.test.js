const { isPalindrome } = require("../index");

it.each([
  {
    inputStr: "",
    expectOutput: true,
  },
  {
    inputStr: "A man, a plan, a canal: Panama",
    expectOutput: true,
  },
  {
    inputStr: "race a car",
    expectOutput: false,
  },
  {
    inputStr: "ab_a",
    expectOutput: true,
  },
  {
    inputStr: "0P",
    expectOutput: false,
  },
])(
  "should determine whether or not the string is a valid palindrome",
  ({ inputStr, expectOutput }) => {
    const isPal = isPalindrome(inputStr);
    expect(isPal).toEqual(expectOutput);
  }
);
