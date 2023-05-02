## [2661. First Completely Painted Row or Column](https://leetcode.com/problems/first-completely-painted-row-or-column/description/)

For this one working backwards from the end goal helped me to come up with a solution.

End Goal: determine when all the cells in a row or column is completely painted.

1. identify the rows and columns
2. keep track of when a new cell in a row is painted. If the total painted cells in the row / col is the total length of
the row / col then we have our answer.
3. in order to find in which row / col a particular number belongs to we can get its coordinates. Once we have the coordinates
for each number in the matrix then we can just loop through the array, refer to its coordinates we calculated and update the
count for the row / col it belongs to.