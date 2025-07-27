package day4

import day2.isExample
import utils.readInputLines

private const val IS_EXAMPLE = false

fun main() {
    part1()
}

fun part1() {
    val rows = readInputLines(day = 4, isExample = IS_EXAMPLE)

    val rowSize = rows.size
    val colSize = rows[0].length

    val wantNext =
        mapOf(
            'X' to 'M',
            'M' to 'A',
            'A' to 'S',
        )

    var count = 0

    val directions =
        listOf(
            Pair(-1, 0), // Up
            Pair(-1, -1), // Diagonal Right UP
            Pair(0, 1), // Right
            Pair(1, 1), // Diagonal Right Down
            Pair(1, 0), // Down
            Pair(1, -1), // Diagonal Down Left
            Pair(0, -1), // Left
            Pair(-1, -1), // Diagonal Left Up
        )

    fun backtrack(
        row: Int,
        col: Int,
        next: Char?,
        dir: Pair<Int, Int>,
    ) {
        if (next == null) {
            return
        }

        if (row < 0 || row >= rowSize || col < 0 || col >= colSize) {
            return
        }

        if (rows[row][col] != next) {
            return
        }

        if (next == 'S') {
            count++
            return
        }

        backtrack(row + dir.first, col + dir.second, wantNext[next], dir)
    }

    for ((rowIdx, rowVal) in rows.withIndex()) {
        for (colIdx in rowVal.indices) {
            for (dir in directions) {
                backtrack(rowIdx, colIdx, 'X', dir)
            }
        }
    }

    println(count)
}
