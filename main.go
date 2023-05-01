// Package main
package main

import (
	//"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"loco-tasks/data"
	"loco-tasks/handlers"
	"loco-tasks/helpers"

	//"loco-tasks/models"
	"loco-tasks/status"
)

func main() {

	data.LoadTasks()

	// removinf finished tasks
	for keyTask, taskRecord := range data.GetTasks() {
		if taskRecord.CurrentStatus == status.Finished {
			data.RemoveTask(keyTask)
		}
	}

	/*
	go func() {
		for {
			data.SaveTasks()
			time.Sleep(5 * time.Second)
		}
	}()
	*/
	
	// Starting tasks
	// https://www.sohamkamani.com/golang/exec-shell-command/#killing-a-child-process
	// https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html
	go func() {
		for {
			for keyTask, taskRecord := range data.GetTasks() {
				// start the waiting task
				if taskRecord.CurrentStatus == status.Waiting {

					// run it / do something
					cmd := &exec.Cmd{
						Path:   taskRecord.Path,
						Args:   taskRecord.Args,
						Stdout: os.Stdout,
						Stderr: os.Stdout,
					}
					newEnv := append(os.Environ(), taskRecord.Env...)
					cmd.Env = newEnv

					log.Println("🚙", cmd.Args)

					err := cmd.Start()

					if err != nil {
						// update the status
						taskRecord.CurrentStatus = status.Failed
						taskRecord.StatusDescription = err.Error()
						taskRecord.FailedAt = time.Now()

						log.Println("🚗", err.Error())
					} else {
						// update the status
						taskRecord.CurrentStatus = status.Started
						taskRecord.StatusDescription = ""
						taskRecord.StartedAt = time.Now()
						taskRecord.Pid = cmd.Process.Pid
						taskRecord.Cmd = cmd
					}
					// update the task
					data.SetTask(keyTask, taskRecord)
				}
			}
			time.Sleep(5 * time.Second)
		}
	}()

	// Monitoring running tasks
	go func() {
		for {
			for keyTask, taskRecord := range data.GetTasks() {
				// start the waiting task
				if taskRecord.CurrentStatus == status.Started {
					//log.Println("🚙", keyTask)

					process, _ := os.FindProcess(taskRecord.Pid)
					err := process.Signal(syscall.Signal(0))

					if err != nil {
						log.Println("🚗[unknown]", process.Pid, err)
						// remove the task, the process does not exist => do it at start
						// data.RemoveTask(keyTask)
						taskRecord.CurrentStatus = status.Finished
						taskRecord.StatusDescription = "🚗[unknown]"

					} else {
						log.Println("🚕[running]", process.Pid, taskRecord.Name, "("+taskRecord.Description+")")
						taskRecord.StatusDescription = "🚕[running]"
					}

					taskRecord.CheckedAt = time.Now()

					// update the task
					data.SetTask(keyTask, taskRecord)
				}
			}
			data.SaveTasks()
			
			time.Sleep(5 * time.Second)
		}
	}()


	httpServer := gin.Default()

	locoTaskEndPoint := helpers.GetEnv("LOCO_TASKS_ENDPOINT", "tasks")

	// 🚧 work in progress
	httpServer.POST("/admin/"+locoTaskEndPoint+"/start", handlers.CreateTask)

	
	httpServer.GET("/admin/"+locoTaskEndPoint+"/list", handlers.GetTasks)
	httpServer.GET("/admin/"+locoTaskEndPoint+"/list/:group", handlers.GetTasksOfGroup)
	
	httpServer.DELETE("/admin/"+locoTaskEndPoint+"/processes/:group", handlers.StopProcessesOfGroup)


	if helpers.GetEnv("LOCO_TASKS_CRT", "") != "" {
		httpServer.RunTLS(
			":"+helpers.GetEnv("LOCO_TASKS_HTTPS_PORT", "4443"),
			helpers.GetEnv("LOCO_TASKS_CRT", "certs/loco-tasks.local.crt"),
			helpers.GetEnv("LOCO_TASKS_KEY", "certs/loco-tasks.local.key"),
		)
	} else {
		httpServer.Run(":" + helpers.GetEnv("LOCO_TASKS_HTTP_PORT", "9090"))
	}

}
