package utils

import kotlin.io.path.Path
import kotlin.io.path.readLines

const val BASE_PATH = "AdventOfCode/2025/day"

fun readInputLines(
    day: Int,
    isExample: Boolean = true,
): List<String> {
    val fileName = if (isExample) "${BASE_PATH + day}/inputExample.txt" else "${BASE_PATH + day}/input.txt"
    val path = Path(fileName)
    return path.readLines()
}
