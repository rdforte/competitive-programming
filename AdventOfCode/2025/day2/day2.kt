package day2

import kotlin.io.path.Path
import kotlin.io.path.readLines
import kotlin.math.absoluteValue

const val basePath = "AdventOfCode/2025/day2"
const val isExample = false

/*
* LEARNINGS:
* 5 scope level functions - let, run, with, apply, also.
* scope level function takes an Extension Lambda which gives us access to the receiver.
*
* mutable lists have a method called removeAt which allows us to easily remove the element at that index.
* */

fun main() {
    partOne()
    partTwo()
}

fun partTwo() {
    val reports = readInputLines().map { it.split(" ").map { level -> level.toInt() } }

    var numSafeReports = 0

    for (report in reports) {
        if (isValidReport2(report)) {
            numSafeReports++
        }
    }

    println("Answer Part Two: $numSafeReports")
}

// Brute force by removing 1 element at a time and running isValidReport.
// The input is not that large so this works and keeps the code relatively simple.
fun isValidReport2(report: List<Int>): Boolean {
    if (isValidReport(report)) {
        return true
    }

    for (i in 0..report.lastIndex) {
        if (isValidReport(report.toMutableList().apply { removeAt(i) })) { // use scope level function apply
            return true
        }
    }

    return false
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
