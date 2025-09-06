package day5

import utils.readInputLines

private const val IS_EXAMPLE = false
private const val DAY = 5

fun main() {
    part1()
    part2()
}

fun part1() {
    val (rules, updates) = readInput()
    val (pos, _) = getPosNegUpdates(rules, updates)

    val correctUpdates = pos.sumOf { it[it.size / 2] }
    println("ANSWER PART 1: $correctUpdates")
}

fun part2() {
    val (rules, updates) = readInput()
    val (_, neg) = getPosNegUpdates(rules, updates)

    fun compare(
        a: Int,
        b: Int,
    ): Int {
        if (rules[b]?.contains(a) == true) {
            return -1
        }
        return 1
    }

    val correctUpdates = neg.map { it.sortedWith(::compare) }.sumOf { it[it.size / 2] }
    println("ANSWER PART 2: $correctUpdates")
}

fun readInput(): Pair<Map<Int, List<Int>>, List<List<Int>>> {
    val (pageOrderingRules, updates) =
        readInputLines(DAY, IS_EXAMPLE).fold(
            initial = Pair(mutableListOf<List<Int>>(), mutableListOf<List<Int>>()),
        ) { acc, cur ->
            if (cur.isEmpty()) {
                acc
            } else if (cur.contains("|")) {
                acc.first.add(cur.split("|").map { it.toInt() })
                acc
            } else {
                acc.second.add(cur.split(",").map { it.toInt() })
                acc
            }
        }

    val rules = pageOrderingRules.groupBy { it[1] }.mapValues { (_, values) -> values.map { it[0] } }

    return Pair(rules, updates)
}

fun getPosNegUpdates(
    rules: Map<Int, List<Int>>,
    updates: List<List<Int>>,
): Pair<List<List<Int>>, List<List<Int>>> {
    fun isCorrectOrder(nums: List<Int>): Boolean {
        for (idx in 1 until nums.size) {
            if (rules[nums[idx]]?.contains(nums[idx - 1]) != true) {
                return false
            }
        }
        return true
    }

    return updates.partition(::isCorrectOrder)
}
