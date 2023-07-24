module.exports.removeDuplicates = function (nums) {
  let totalNums = nums.length;
  if (totalNums <= 1) return totalNums;

  let removed = 0;

  for (let i = 1, n = 1; i < totalNums; i++) {
    if (nums[i] === nums[i - 1]) {
      removed++;
    } else {
      nums[n] = nums[i];
      n++;
    }
  }

  return totalNums - removed;
};
