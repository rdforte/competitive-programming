module.exports.searchMatrix = function (matrix, target) {
  let row = 0;
  for (let i = 0; i < matrix.length; i++) {
    const matrixRowLastNum = matrix[i][matrix[i].length - 1];
    if (target <= matrixRowLastNum) {
      row = i;
      break;
    }
  }

  for (let l = 0, r = matrix[row].length - 1; l <= r; ) {
    const middle = l + Math.trunc((r - l) / 2);
    if (matrix[row][middle] === target) {
      return true;
    }

    if (target > matrix[row][middle]) {
      l = middle + 1;
    } else {
      r = middle - 1;
    }
  }

  return false;
};
