package jaetan.fr.plugins.routing

import io.ktor.http.*
import io.ktor.server.application.*
import io.ktor.server.request.*
import io.ktor.server.response.*
import io.ktor.server.routing.*
import jaetan.fr.models.Task
import jaetan.fr.models.generateId
import jaetan.fr.models.taskStorage

fun Route.taskRouting() {
    route("/tasks") {
        get {
            if (taskStorage.isNotEmpty()) {
                call.respond(taskStorage)
            } else {
                call.respondText("No tasks found")
            }
        }

        get("{id}") {
            val taskId = call.parameters["id"] ?: return@get call.respondText(
                text = "Missing id",
                status = HttpStatusCode.BadRequest
            )
            val task = taskStorage.find { it.id == taskId } ?: return@get call.respondText(
                text = "No task with id: $taskId",
                status = HttpStatusCode.NotFound
            )

            call.respond(task)
        }

        post {
            val task = call.receiveNullable<Task>() ?: return@post call.respondText(
                text = "Task not Found",
                status = HttpStatusCode.BadRequest
            )

            task.generateId()
            taskStorage.add(task)
            call.respondText("Task stored correctly", status = HttpStatusCode.Created)
        }

        put("done/{id}/{isDone}") {
            val id = call.parameters["id"] ?: return@put call.respondText(
                text = "Missing id",
                status = HttpStatusCode.BadRequest
            )
            val isDone = call.parameters["isDone"]?.toBooleanStrictOrNull() ?: return@put call.respondText(
                text = "Missing isDone",
                status = HttpStatusCode.BadRequest
            )

            val taskId = taskStorage.indexOfFirst { it.id == id }

            if (taskId != -1) {
                taskStorage.find { it.id == id }?.isDone = isDone
                call.respondText("Task updated correctly", status = HttpStatusCode.Accepted)
            } else {
                call.respondText(
                    text = "Task not Found",
                    status = HttpStatusCode.BadRequest
                )
            }
        }

        delete("{id}") {
            val id = call.parameters["id"] ?: return@delete call.respond(HttpStatusCode.BadRequest)

            if (taskStorage.removeIf { it.id == id }) {
                call.respondText("Task removed correctly", status = HttpStatusCode.Accepted)
            } else {
                call.respondText("Task not Found", status = HttpStatusCode.NotFound)
            }
        }
    }
}