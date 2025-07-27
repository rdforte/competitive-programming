package day3

import utils.readInputLines

const val IS_EXAMPLE = false
const val DAY = 3

const val DONT_MUL = "don't()"
const val DO_MUL = "do()"

const val MUL_RGX = """mul\((\d+),(\d+)\)"""
const val DONT_RGX = """don't\(\)"""
const val DO_RGX = """do\(\)"""

/*
* LEARNINGS:
* fold to multiple all elements.
* Regex capturing groups for getting elements you need from a string.
* extension functions make for cleaner / more readable code.
* using """some-rgx""" triple quotes to avoid double \\ for rgx - String Literal.
* converting string to Regex with .toRegex().
* */

fun main() {
    partOne()
    partTwo()
}

fun partTwo() {
    val line = readInputLines(DAY, IS_EXAMPLE)

    var shouldInclude = true
    var res = 0
    line.forEach { line ->
        "$MUL_RGX|$DONT_RGX|$DO_RGX".toRegex().findAll(line).forEach {
            when (it.value) {
                DO_MUL -> shouldInclude = true
                DONT_MUL -> shouldInclude = false
                else ->
                    if (shouldInclude) {
                        res += it.multiplyNumbers()
                    }
            }
        }
    }

    println("Answer Part Two: $res")
}

fun partOne() {
    val line = readInputLines(DAY, IS_EXAMPLE)

    var res = 0
    line.forEach { line ->
        res +=
            MUL_RGX.toRegex().findAll(line).sumOf {
                it.multiplyNumbers()
            }
    }

    println("Answer Part One: $res")
}

private fun MatchResult.multiplyNumbers(): Int {
    val (first, second) = this.destructured
    return first.toInt() * second.toInt()
}
