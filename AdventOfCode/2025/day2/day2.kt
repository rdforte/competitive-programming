package day2

import kotlin.io.path.Path
import kotlin.io.path.readLines
import kotlin.math.absoluteValue

const val basePath = "AdventOfCode/2025/day2"
const val isExample = false

fun main() {
    partOne()
}

fun partOne() {
    val reports = readInputLines().map { it.split(" ").map { level -> level.toInt() } }

    var numSafeReports = 0

    for (report in reports) {
        if (isValidReport(report)) {
            numSafeReports++
        }
    }

    println("Answer Part One: $numSafeReports")
}

fun isValidReport(report: List<Int>): Boolean {
    if (report.size <= 1) {
        return true
    }

    var isIncreasing = true
    var isDecreasing = true
    var rangeLessThanThree = true

    for (i in 1..<report.size) {
        val before = report[i - 1]
        val after = report[i]
        rangeLessThanThree = rangeLessThanThree && (after - before).absoluteValue <= 3
        isIncreasing = isIncreasing && after > before
        isDecreasing = isDecreasing && after < before
    }

    return rangeLessThanThree && (isDecreasing || isIncreasing)
}

fun readInputLines(): List<String> {
    val fileName = if (isExample) "$basePath/inputExample.txt" else "$basePath/input.txt"
    val path = Path(fileName)
    return path.readLines()
}
