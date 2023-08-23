module.exports.productExceptSelf = (nums) => {
  let prefix = new Array(nums.length);
  let postfix = new Array(nums.length);

  for (let i = 0, j = nums.length - 1; i < nums.length; i++, j--) {
    prefix[i] = i === 0 ? 1 : nums[i - 1] * prefix[i - 1];
    postfix[j] = j === nums.length - 1 ? 1 : nums[j + 1] * postfix[j + 1];
  }

  const sol = [];
  for (let i = 0; i < nums.length; i++) {
    sol.push(prefix[i] * postfix[i]);
  }

  return sol;
};
