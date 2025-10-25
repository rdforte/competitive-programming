plugins {
    kotlin("jvm") version "2.2.20"
}

repositories {
    mavenCentral()
}

sourceSets {
    main {
        kotlin.srcDirs("2025") // Add multiple advent of code src dirs here.
    }
}

dependencies {
    implementation("org.jetbrains.kotlinx:kotlinx-coroutines-core:1.10.2")
}
