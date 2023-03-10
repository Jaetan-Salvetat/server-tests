package jaetan.fr

import io.ktor.server.application.*
import io.ktor.server.engine.*
import io.ktor.server.netty.*
import jaetan.fr.plugins.*
import jaetan.fr.plugins.routing.configureRouting

fun main() {
    embeddedServer(Netty, port = 8080, host = "0.0.0.0", module = Application::module)
        .start(wait = true)
}

fun Application.module() {
    configureSerialization()
    configureRouting()
}
