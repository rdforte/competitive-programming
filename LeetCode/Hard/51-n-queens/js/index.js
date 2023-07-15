const QUEEN = "Q";
const EMPTY_SQUARE = ".";

module.exports.solveNQueens = function (n) {
  const attackedColumns = new Set();
  const attackedDiagonals = new Set();
  const attackedAntiDiagonals = new Set();
  const board = [];

  for (let r = 0; r < n; r++) {
    board.push([]);
    for (let c = 0; c < n; c++) {
      board[r].push(EMPTY_SQUARE);
    }
  }

  return backtrack(
    0,
    board,
    attackedColumns,
    attackedDiagonals,
    attackedAntiDiagonals
  );
};

const backtrack = (
  row,
  board,
  attackedColumns,
  attackedDiagonals,
  attackedAntiDiagonals
) => {
  const solutions = [];
  for (let col = 0; col < board.length; col++) {
    const diagonalQueenAttacks = row - col;
    const antiDiagonalQueenAttacks = row + col;
    if (
      attackedColumns.has(col) ||
      attackedDiagonals.has(diagonalQueenAttacks) ||
      attackedAntiDiagonals.has(antiDiagonalQueenAttacks)
    ) {
      continue;
    }

    // add queen to board
    board[row][col] = QUEEN;
    attackedColumns.add(col);
    attackedDiagonals.add(diagonalQueenAttacks);
    attackedAntiDiagonals.add(antiDiagonalQueenAttacks);
    // queen added to board and attacked cells marked.
    // add queen to next row
    if (row < board.length - 1) {
      const sol = backtrack(
        row + 1,
        board,
        attackedColumns,
        attackedDiagonals,
        attackedAntiDiagonals
      );
      sol.forEach((s) => solutions.push(s));
    } else {
      const formattedSol = board.map((s) => s.join(""));
      solutions.push(formattedSol);
    }

    // remove queen from board and proceed to next col
    board[row][col] = EMPTY_SQUARE;
    attackedColumns.delete(col);
    attackedDiagonals.delete(diagonalQueenAttacks);
    attackedAntiDiagonals.delete(antiDiagonalQueenAttacks);
  }

  return solutions;
};
