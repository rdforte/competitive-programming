plugins {
    kotlin("jvm") version "2.1.21"
}

repositories {
    mavenCentral()
}

sourceSets {
    main {
        kotlin.srcDirs("2025") // Add multiple advent of code src dirs here.
    }
}
