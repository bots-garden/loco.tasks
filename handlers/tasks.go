// Package handlers ...
package handlers

import (
	"loco-tasks/data"
	"loco-tasks/models"
	"loco-tasks/status"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetTasks is a handler for /admin/tasks/list
//   - It returns a list of all tasks
//   - Usage:
//     curl http://localhost:8080/admin/tasks/list
func GetTasks(c *gin.Context) {
	tasksList := data.GetTasks()
	c.JSON(http.StatusAccepted, &tasksList)
}

// GetTasksOfGroup is a handler for /admin/tasks/list/:group
//   - It returns a list of all tasks for a group
//   - Usage:
//     curl http://localhost:8080/admin/tasks/list/:group
//     curl http://localhost:7070/admin/tasks/list/hey-group
func GetTasksOfGroup(c *gin.Context) {
	tasksList := data.GetTasksOfGroup(c.Param("group"))
	c.JSON(http.StatusAccepted, &tasksList)
}

// DeleteTasksOfGroup is a handler for /admin/tasks/list/:group
func DeleteTasksOfGroup(c *gin.Context) {
	tasksList := data.DeleteTasksOfGroup(c.Param("group"))
	c.JSON(http.StatusAccepted, &tasksList)
}

/*
CreateTask is a handler for /admin/tasks/registration (POST)

# It creates a task record in the tasks map

Usage:
curl -X POST http://localhost:7070/admin/tasks/start \
-H 'Content-Type: application/json; charset=utf-8' \
-d @- << EOF

	{
	    "name":"start-hello",
	    "description":"start-hello wasm service",
	    "path":"this is the path",
	    "args": ["arg1", "arg2"],
	    "env": ["ONE=1", "TWO=2"]
	}

EOF
*/

// CreateTask is a handler for /admin/tasks/start
func CreateTask(c *gin.Context) {
	task := models.Task{}

	if err := c.BindJSON(&task); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	log.Println("📝", task)

	taskRecord := models.TaskRecord{
		Group:       task.Group,
		Name:        task.Name,
		Description: task.Description,
		Path:        task.Path,
		Args:        task.Args,
		Env:         task.Env,
		CreatedAt:   time.Now(),
	}

	data.SetTask(uuid.New().String(), taskRecord)

	c.JSON(http.StatusAccepted, &taskRecord)

}

// StopProcessesOfGroup stops all the processes of a given task
func StopProcessesOfGroup(c *gin.Context) {

	//tasksList := data.GetTasksOfGroup(c.Param("group"))

	for keyTask, taskRecord := range data.GetTasksOfGroup(c.Param("group")) {
		if taskRecord.CurrentStatus == status.Started {
			// kill process
			log.Println("🔥", "killing", keyTask, taskRecord.Pid)

			err := taskRecord.Cmd.Process.Kill()

			//err := syscall.Kill(taskRecord.Pid, syscall.Signal(syscall.SIGTERM))
			if err != nil {
				log.Println("💥", err)
			}
			data.RemoveTask(keyTask)
		}

		//taskRecord.Cmd.Process.Pid

		//taskRecord.Cmd.Process.Kill()
		//data.RemoveTask(keyTask)

	}

	//c.JSON(http.StatusAccepted, &tasksList)
	c.JSON(http.StatusAccepted, "")


}
