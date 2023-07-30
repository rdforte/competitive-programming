module.exports.containsNearbyDuplicate = function (nums, k) {
  const mp = new Map();

  for (let i = 0; i < nums.length; i++) {
    if (mp.has(nums[i]) && Math.abs(mp.get(nums[i]) - i) <= k) {
      return true;
    }
    mp.set(nums[i], i);
  }

  return false;
};
