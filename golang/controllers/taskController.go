package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang/models"
	"net/http"
	"strconv"
)

func GetAll(c *gin.Context) {
	if len(models.TaskStorage) != 0 {
		c.JSON(http.StatusOK, gin.H{"data": models.TaskStorage})
	} else {
		c.JSON(http.StatusOK, "No tasks found")
	}
}

func GetById(c *gin.Context) {
	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Missing id")
		fmt.Println(err)
		return
	}

	for _, task := range models.TaskStorage {
		if task.Id == taskId {
			c.JSON(http.StatusOK, gin.H{"data": task})
			return
		}
	}

	c.JSON(http.StatusBadRequest, "No task with id: "+strconv.Itoa(taskId))
}

func Add(c *gin.Context) {
	var task models.Task
	err := c.ShouldBindJSON(&task)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Task not Found")
		fmt.Println(err)
		return
	}

	task.Id = len(models.TaskStorage)

	models.TaskStorage = append(models.TaskStorage, task)
	c.JSON(http.StatusCreated, "Task stored correctly")
}

func UpdateDone(c *gin.Context) {
	var id int
	isDone, err := strconv.ParseBool(c.Param("isDone"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Missing isDone")
		fmt.Println(err)
		return
	}
	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Missing id")
		fmt.Println(err)
		return
	}

	for index, task := range models.TaskStorage {
		if task.Id == id {
			models.TaskStorage[index].IsDone = isDone
			c.JSON(http.StatusAccepted, "Task updated correctly")
			return
		}
	}

	c.JSON(http.StatusBadRequest, "Task not Found")
}

func DeleteById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Missing id")
		fmt.Println(err)
		return
	}

	for index, task := range models.TaskStorage {
		if task.Id == id {
			models.TaskStorage = append(models.TaskStorage[:index], models.TaskStorage[index+1:]...)
			c.JSON(http.StatusAccepted, "Task deleted correctly")
			return
		}
	}

	c.JSON(http.StatusBadRequest, "Task not Found")
}
