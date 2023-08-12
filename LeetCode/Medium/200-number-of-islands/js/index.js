const grid = [
  ["1", "1", "0", "0", "0"],
  ["1", "1", "0", "0", "0"],
  ["0", "0", "1", "0", "0"],
  ["0", "0", "0", "1", "1"],
];

const numIslands = (grid) => {
  let islands = 0;

  const dfs = (r, c) => {
    if (
      r < 0 ||
      r >= grid.length ||
      c < 0 ||
      c >= grid[0].length ||
      grid[r][c] === "0"
    )
      return;

    grid[r][c] = "0";

    for (let d of directions) {
      dfs(r + d.row, c + d.col);
    }
  };

  for (let r = 0; r < grid.length; r++) {
    for (let c = 0; c < grid[r].length; c++) {
      if (grid[r][c] === "0") continue;

      if (grid[r][c] === "1") {
        islands++;
        dfs(r, c);
      }
    }
  }
  return islands;
};

const directions = [
  { row: -1, col: 0 },
  { row: 1, col: 0 },
  { row: 0, col: -1 },
  { row: 0, col: 1 },
];

console.log(numIslands(grid));
