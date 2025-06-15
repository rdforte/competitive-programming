package day2

import kotlin.io.path.Path
import kotlin.io.path.readLines
import kotlin.math.absoluteValue

const val basePath = "AdventOfCode/2025/day2"
const val isExample = false

fun main() {
    partOne()
}

fun partOne(): Unit {
    val reports = readInputLines().map { it.split(Regex("\\s")).map { level -> level.toInt() } }

    var numSafeReports = 0

    outer@ for (report in reports) {
        if (report.size <= 1) {
            numSafeReports++
            continue
        }

        val isIncreasing = report[1] > report[0]

        if (isIncreasing) {
            for (i in 1..<report.size) {
                if (report[i] < report[i - 1]) {
                    continue@outer
                }

                if (report[i] - report[i - 1] == 0 || (report[i] - report[i - 1]).absoluteValue > 3) {
                    continue@outer
                }
            }
        }

        if (!isIncreasing) {
            for (i in 1..report.size - 1) {
                if (report[i] > report[i - 1]) {
                    continue@outer
                }

                if (report[i] - report[i - 1] == 0 || (report[i] - report[i - 1]).absoluteValue > 3) {
                    continue@outer
                }
            }
        }

        numSafeReports++
    }

    println("Answer Part One: $numSafeReports")
}

fun readInputLines(): List<String> {
    val fileName = if (isExample) "$basePath/inputExample.txt" else "$basePath/input.txt"
    val path = Path(fileName)
    return path.readLines()
}
