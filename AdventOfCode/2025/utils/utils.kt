package utils

import kotlin.io.path.Path
import kotlin.io.path.readLines

const val basePath = "AdventOfCode/2025/day"

fun readInputLines(day: Int, isExample: Boolean = true): List<String> {
    val fileName = if (isExample) "${basePath + day}/inputExample.txt" else "${basePath + day}/input.txt"
    val path = Path(fileName)
    return path.readLines()
}
