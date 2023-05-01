// Package handlers ...
package handlers

import (
	"loco-tasks/data"
	"loco-tasks/models"
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

/*

 */

/*
	RegisterTask is a handler for /admin/tasks/registration (POST)

# It creates a task record in the tasks map

Usage:
curl -X POST http://localhost:7070/admin/tasks/registration \
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
func RegisterTask(c *gin.Context) {
	task := models.Task{}

	if err := c.BindJSON(&task); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	log.Println("📝", task)

	taskRecord := models.TaskRecord{
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
