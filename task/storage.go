package task

import (
	"encoding/json"
	. "fmt"
	"log"
	"os"
)

const defaultTaskFile = "json.json"

func LoadTasks() []Task {
	file, err := os.OpenFile(defaultTaskFile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var tasks []Task
	file.Seek(0, 0)
	if err := json.NewDecoder(file).Decode(&tasks); err != nil {
		Println("Ошибка чтения задач:", err)
	}
	return tasks
}

func SaveTasks(tasks []Task) {
	file, err := os.OpenFile(defaultTaskFile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.Truncate(0) // очищаем файл
	file.Seek(0, 0)  // в начало
	if err := json.NewEncoder(file).Encode(tasks); err != nil {
		Println("Ошибка записи:", err)
		return
	}
}
