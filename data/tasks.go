// Package data ...
package data

import (
	"encoding/json"
	"io/ioutil"
	"loco-tasks/helpers"
	"loco-tasks/models"
	"log"
	"os"
)

var tasksMap = make(map[string] models.TaskRecord)
var storage = helpers.GetEnv("LOCO_TASKS_STORAGE", "tasks.json")

// SetTask ...
func SetTask(keyTask string, task models.TaskRecord) {
	tasksMap[keyTask] = task
}

// GetTask ...
func GetTask(keyTask string) models.TaskRecord {
	return tasksMap[keyTask]
}

// RemoveTask ...
func RemoveTask(keyTask string) {
	delete(tasksMap, keyTask)
}

// GetTasks ...
func GetTasks() map[string] models.TaskRecord {
	return tasksMap
}

// GetTasksOfGroup → get tasks of a group
func GetTasksOfGroup(group string) map[string] models.TaskRecord {
	filteredList := make(map[string] models.TaskRecord)
	for keyTask, taskRecord := range tasksMap {
		if taskRecord.Group == group {
			filteredList[keyTask] = taskRecord
		}
	}
	return filteredList
}

// DeleteTasksOfGroup ...
func DeleteTasksOfGroup(group string) map[string] models.TaskRecord {
	for keyTask, taskRecord := range tasksMap {
		if taskRecord.Group == group {
			delete(tasksMap, keyTask)
		}
	}
	return tasksMap
}




// SaveTasks ...
func SaveTasks() {
	// save map to json file
	jsonString, err := json.MarshalIndent(tasksMap, "", "  ")
	if err != nil {
		log.Println(err)
	}
	err = ioutil.WriteFile(storage, jsonString, 0644)
	if err != nil {
		log.Println(err)
	}
}
// TODO: handle errors correctly

// LoadTasks ...
func LoadTasks() {
	// load map from json file
	jsonFile, err := os.Open(storage)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Println(err)
	}
	log.Println("Successfully Opened", storage)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array
	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &tasksMap)

}
// TODO: handle errors correctly