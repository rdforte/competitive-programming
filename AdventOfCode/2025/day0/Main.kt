package day0

fun main() {
    val names = listOf("Ryan", "Alicia", "Bianca")
    val ages = listOf(31, 29, 29)

    val people = names.zip(ages) { name, age -> Person(name, age) }

    val peopleMap = people.groupBy(Person::group)

    println(peopleMap)
}

data class Person(
    val name: String,
    val age: Int,
)

fun Person.group(): String = "$name $age"
