package part1

import kotlinx.coroutines.coroutineScope

suspend fun main() {
    for (i in 1..10) {
        test(i)
    }
}

suspend fun test(i: Int) =
    coroutineScope {
        println("World! $i")
    }
