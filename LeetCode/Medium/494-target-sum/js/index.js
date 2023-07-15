module.exports.findTargetSumWays = function (nums, target) {
  return findAllSums(nums, target, 0, 0);
};

const findAllSums = (nums, target, sum, i) => {
  if (i === nums.length) {
    return sum === target ? 1 : 0;
  }

  let total = 0;
  total += findAllSums(nums, target, sum - nums[i], i + 1);
  total += findAllSums(nums, target, sum + nums[i], i + 1);

  return total;
};
