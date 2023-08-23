module.exports.isValidSudoku = function (board) {
  const cols = new Array(9);
  const rows = new Array(9);
  const grid = [
    [new Set(), new Set(), new Set()],
    [new Set(), new Set(), new Set()],
    [new Set(), new Set(), new Set()],
  ];

  for (let r = 0; r < 9; r++) {
    for (let c = 0; c < 9; c++) {
      if (r === 0) {
        cols[c] = new Set();
        rows[c] = new Set();
      }

      if (Number.isNaN(+board[r][c])) {
        continue;
      }

      if (rows[c].has(board[r][c])) {
        return false;
      }

      if (cols[r].has(board[r][c])) {
        return false;
      }

      const gridRow = Math.trunc(r / 3);
      const gridCol = Math.trunc(c / 3);

      if (grid[gridRow][gridCol].has(board[r][c])) {
        return false;
      }

      rows[c].add(board[r][c]);
      cols[r].add(board[r][c]);
      grid[gridRow][gridCol].add(board[r][c]);
    }
  }

  return true;
};
