module.exports.searchInsert = function (nums, target) {
  let middle;
  for (let l = 0, r = nums.length - 1; l <= r; ) {
    middle = Math.trunc((l + r) / 2);

    if (nums[middle] === target) return middle;

    if (target > nums[middle]) {
      l = middle + 1;
    } else {
      r = middle - 1;
    }
  }

  return target > nums[middle] ? middle + 1 : middle;
};
