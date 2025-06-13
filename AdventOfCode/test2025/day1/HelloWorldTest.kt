package day1

import org.junit.jupiter.api.Assertions.*
import org.junit.jupiter.api.Test

class HelloWorldTest {

    @Test
    fun `should return hello`() {
        assertEquals("hello", HelloWorld().hello())
    }
}