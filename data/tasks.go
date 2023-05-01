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


// SaveTasks ...
func SaveTasks() {
	// save map to json file
	jsonString, _ := json.MarshalIndent(tasksMap, "", "  ")
	_ = ioutil.WriteFile(storage, jsonString, 0644)
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