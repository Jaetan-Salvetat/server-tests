package jaetan.fr.models

import jaetan.fr.extensions.isNotNull
import kotlinx.serialization.Serializable
import java.util.UUID

@Serializable
data class Task(
    var text: String,
    var isDone: Boolean,
    var id: String = ""
)

fun Task.generateId() {
    id = UUID.randomUUID().toString()

    if (taskStorage.find { it.id == id }.isNotNull()) {
        generateId()
    }
}

val taskStorage = mutableListOf<Task>()