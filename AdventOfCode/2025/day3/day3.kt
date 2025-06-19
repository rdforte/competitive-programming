package day3

import utils.readInputLines

const val isExample = false
const val day = 3

const val dontMul = "don't()"
const val doMul = "do()"

const val mulRgx = """mul\((\d+),(\d+)\)"""
const val dontRgx = """don't\(\)"""
const val doRgx = """do\(\)"""

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
    val line = readInputLines(day, isExample)

    var shouldInclude = true
    var res = 0
    line.forEach { line ->
        "$mulRgx|$dontRgx|$doRgx".toRegex().findAll(line).forEach {
            when (it.value) {
                doMul -> shouldInclude = true
                dontMul -> shouldInclude = false
                else -> if (shouldInclude) {
                    res += it.multiplyNumbers()
                }
            }
        }
    }

    println("Answer Part Two: $res")
}

fun partOne() {
    val line = readInputLines(day, isExample)

    var res = 0
    line.forEach { line ->
        res += mulRgx.toRegex().findAll(line).sumOf {
            it.multiplyNumbers()
        }
    }

    println("Answer Part One: $res")
}

private fun MatchResult.multiplyNumbers(): Int {
    val (first, second) = this.destructured
    return first.toInt() * second.toInt()
}