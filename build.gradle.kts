plugins {
    kotlin("jvm") version "2.1.21"
}

group = "com.github.rdforte"

repositories {
    mavenCentral()
}

dependencies {
    testImplementation(kotlin("test"))
}

tasks.test {
    useJUnitPlatform()
}

kotlin {
    sourceSets {
        println(this.forEach { it.name })
    }
}

sourceSets {
    main {
        kotlin.srcDirs("AdventOfCode/2025")
    }
}
