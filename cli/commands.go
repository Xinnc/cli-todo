package cli

import (
	"WorkWithFiles/task"
	"os"
	"strings"
)

func RunCLI(command string) {
	tasks := task.LoadTasks()
	var flag string
	var title string
	var description string
	var id string

	switch command {
	case "list":
		if len(os.Args) > 2 {
			flag = os.Args[2]
		}
		if len(os.Args) > 3 {
			title = strings.ToLower(os.Args[3])
		}
		task.GetTasks(tasks, flag, title)
	case "add":
		if len(os.Args) > 3 {
			title = os.Args[2]
			description = os.Args[3]
		}
		task.AddTask(tasks, title, description)
	case "done":
		if len(os.Args) > 3 {
			id = os.Args[2]
		}
		task.DoneTask(tasks, id)
	case "delete":
		if len(os.Args) > 3 {
			id = os.Args[2]
		}
		task.DeleteTask(tasks, id)
	case "clear":
		task.ClearSystem()
	default:
		task.InfoTask()
	}
}
