package day3

import utils.readInputLines

const val isExample = false
const val day = 3
const val dontMul = "don't()"
const val doMul = "do()"

/*
* LEARNINGS:
* fold is a nice way to multiple all elements.
* */

fun main() {
    partOne()
    partTwo()
}


fun partTwo() {
    val rgx = Regex("mul\\(\\d+,\\d+\\)|don't\\(\\)|do\\(\\)")
    val line = readInputLines(day, isExample)

    var shouldInclude = true
    var res = 0
    line.forEach { line ->
        rgx.findAll(line).forEach {
            when (it.value) {
                doMul -> shouldInclude = true
                dontMul -> shouldInclude = false
                else -> if (shouldInclude) {
                    res += getMultipliedVal(it.value)
                }
            }
        }
    }

    println("Answer Part Two: $res")
}

fun getMultipliedVal(str: String): Int {
    val digitRgx = Regex("\\d+,\\d+")
    return digitRgx.find(str)?.value?.split(",")?.fold(1) { acc, num -> acc * num.toInt() } ?: 0
}

fun partOne() {
    val rgx = Regex("mul\\(\\d+,\\d+\\)")
    val line = readInputLines(day, isExample)

    val digitRgx = Regex("\\d+,\\d+")

    var res = 0
    line.forEach {
        res += rgx.findAll(it)
            .toMutableList().map { mul ->
                digitRgx.findAll(mul.value).flatMap { nums -> nums.value.split(",") }.toMutableList()
            }.sumOf { nums -> nums[0].toInt() * nums[1].toInt() }
    }

    println("Answer Part One: $res")
}