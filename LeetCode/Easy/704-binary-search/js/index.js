module.exports.search = function (nums, target) {
  for (let l = 0, r = nums.length; l <= r; ) {
    const middle = Math.trunc(l + (r - l) / 2);
    if (nums[middle] === target) return middle;

    if (target > nums[middle]) {
      l = middle + 1;
    } else {
      r = middle - 1;
    }
  }

  return -1;
};
