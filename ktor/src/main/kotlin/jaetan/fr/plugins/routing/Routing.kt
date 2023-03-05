package jaetan.fr.plugins.routing

import io.ktor.server.application.*
import io.ktor.server.routing.*

fun Application.configureRouting() {
    routing {
        taskRouting()
    }
}
