plugins {
    kotlin("jvm") version "2.1.21"
}

group = "com.github.rdforte"

repositories {
//    mavenCentral()
}

dependencies {
//    testImplementation(kotlin("test"))
}

//tasks.test {
//    useJUnitPlatform()
//}

sourceSets {
    main {
        kotlin.srcDirs("2025") // Add multiple advent of code src dirs here.
    }
}
