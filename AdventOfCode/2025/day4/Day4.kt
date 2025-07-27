package day4

import utils.readInputLines

private const val IS_EXAMPLE = false

fun main() {
    part1()
    part2()
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
            Pair(-1, 1), // Diagonal Right UP
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

    println("PART ONE ANSWER: $count")
}

fun part2() {
    val rows = readInputLines(day = 4, isExample = IS_EXAMPLE)

    var count = 0

    for (rowIdx in 1 until rows.lastIndex) {
        for (colIdx in 1 until rows[rowIdx].lastIndex) {
            val char = rows[rowIdx][colIdx]
            if (char == 'A') {
                val diagonalTLToBR =
                    (rows[rowIdx - 1][colIdx - 1] == 'M' && rows[rowIdx + 1][colIdx + 1] == 'S') ||
                        (rows[rowIdx - 1][colIdx - 1] == 'S' && rows[rowIdx + 1][colIdx + 1] == 'M')
                val diagonalTRToBL =
                    (rows[rowIdx - 1][colIdx + 1] == 'M' && rows[rowIdx + 1][colIdx - 1] == 'S') ||
                        (rows[rowIdx - 1][colIdx + 1] == 'S' && rows[rowIdx + 1][colIdx - 1] == 'M')

                if (diagonalTRToBL && diagonalTLToBR) {
                    count++
                }
            }
        }
    }

    println("PART TWO ANSWER: $count")
}
