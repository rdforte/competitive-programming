package day1

import utils.readInputLines
import kotlin.math.absoluteValue

const val isExample = false
const val day = 1

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

    val lines = readInputLines(day, isExample)
    lines.forEach {
        val (id1, id2) = it.split(Regex("\\s+"))
        list1 += id1.toInt()
        list2 += id2.toInt()
    }

    return Pair(list1.sorted(), list2.sorted())
}