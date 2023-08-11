module.exports.sortArray = function (nums) {
  return mergeSort(nums, 0, nums.length - 1);
};

const mergeSort = (nums, left, right) => {
  if (left === right) {
    return [nums[left]];
  }

  // break down array
  const middle = left + Math.trunc((right - left) / 2);

  const leftArr = mergeSort(nums, left, middle);
  const rightArr = mergeSort(nums, middle + 1, right);

  const merge = [];

  for (let l = 0, r = 0, n = left; n <= right; n++) {
    if (l >= leftArr.length) {
      merge.push(rightArr[r]);
      r++;
      continue;
    }
    if (r >= rightArr.length) {
      merge.push(leftArr[l]);
      l++;
      continue;
    }

    if (leftArr[l] <= rightArr[r]) {
      merge.push(leftArr[l]);
      l++;
      continue;
    }

    if (leftArr[l] >= rightArr[r]) {
      merge.push(rightArr[r]);
      r++;
      continue;
    }
  }

  return merge;
};
