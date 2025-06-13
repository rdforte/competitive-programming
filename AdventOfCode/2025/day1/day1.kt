package day1

import kotlin.io.path.Path
import kotlin.io.path.readLines
import kotlin.math.absoluteValue

const val basePath = "AdventOfCode/2025/day1"
const val isExample = false

fun main() {
    partOne()
    partTwo()
}

fun partTwo() {
    val (list1, list2) = getListsSorted()

    val seen = mutableMapOf<Int, Int>()

    var total = 0
    list1.forEach { list1Item ->
        val similarityScore = seen.getOrPut(list1Item) {
            list2.filter { list2Item -> list1Item == list2Item }.sum()
        }
        total += similarityScore
    }


    println("Answer Part Two: $total")
}

fun partOne() {
    val (list1, list2) = getListsSorted()

    var sum = 0
    list1.forEachIndexed { idx, id -> sum += (id - list2[idx]).absoluteValue }

    println("Answer Part One: $sum")
}

fun getListsSorted(): Pair<List<Int>, List<Int>> {
    val list1 = mutableListOf<Int>()
    val list2 = mutableListOf<Int>()

    val lines = readInputLines(isExample)
    lines.forEach {
        val (id1, id2) = it.split(Regex("\\s+"))
        list1 += id1.toInt()
        list2 += id2.toInt()
    }

    return Pair(list1.sorted(), list2.sorted())
}

fun readInputLines(isExample: Boolean): List<String> {
    val fileName = if (isExample) "$basePath/inputExample.txt" else "$basePath/input.txt"
    val path = Path(fileName)
    return path.readLines()
}