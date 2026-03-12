package task

import (
	"bufio"
	"encoding/json"
	. "fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

func InfoTask() {
	menu := []string{
		"Команды:",
		"todo list        показать задачи",
		"    Флаги:",
		"    --done                             завершенные задачи",
		"    --todo                             незавершенные задачи",
		"    --today                            задачи созданные сегодня",
		"    --json                             задачи в формате json",
		`    --search "заголовок задачи"        поиск задачи `,
		"",
		"todo add         добавить задачу",
		`todo add "title" "description"       быстрое создание задачи `,
		"todo done ID     отметить выполненной",
		"todo delete ID   удалить задачу",
		"todo help        показать помощь",
		"todo clear       очистить файл с задачами",
	}

	for _, v := range menu {
		Println(v)
	}
}

func PrintTask(t Task) {
	if t.EndDate != "" {
		Printf("%d. %s\n Создана: %s\n Закончена: %s\n Выполнена: %v\nОписание: %s\n",
			t.Id, t.Title, t.StartDate, t.EndDate, t.IsDone, t.Description)
	} else {
		Printf("%d. %s\n Создана: %s\n Выполнена: %v\nОписание: %s\n",
			t.Id, t.Title, t.StartDate, t.IsDone, t.Description)
	}

}

func GetTasks(tasks []Task, flag, title string) {
	if len(tasks) == 0 {
		Println("Список задач пуст")
		return
	}
	sort.Slice(tasks, func(i, j int) bool { return !tasks[i].IsDone && tasks[j].IsDone })

	shouldShow := func(t Task) bool {

		if flag == "" {
			return true
		}

		switch flag {
		case "--done":
			return t.IsDone == true
		case "--todo":
			return t.IsDone == false
		case "--json":
			return false
		case "--today":
			return len(t.StartDate) >= 10 && t.StartDate[:10] == time.Now().Format("2006-01-02")
		case "--search":
			if title != "" {
				return strings.Contains(strings.ToLower(t.Title), strings.ToLower(title))
			}
		default:
			Println("Неизвестный флаг")
			return false
		}
		return true
	}

	if flag == "--json" {
		jsonTasks, err := json.MarshalIndent(tasks, "", "  ")
		if err != nil {
			Println(err)
		}
		Println(string(jsonTasks))
	}

	for _, t := range tasks {
		if shouldShow(t) {
			PrintTask(t)
		}
	}
}

func AddTask(tasks []Task, title, description string) {
	var tsk Task
	scanner := bufio.NewScanner(os.Stdin)
	//сохранение данных разных
	if title != "" && description != "" {
		tsk.Title = title
		tsk.Description = description
	} else {
		Println("Заголовок:")
		scanner.Scan()
		tsk.Title = scanner.Text()
		Println("Описание:")
		scanner.Scan()
		tsk.Description = scanner.Text()
	}
	tsk.StartDate = time.Now().Format("2006-01-02 15:04:05")

	//Вычисление id
	if len(tasks) == 0 {
		tsk.Id = 1
	} else {
		maxId := 0
		for _, t := range tasks {
			if t.Id > maxId {
				maxId = t.Id
			}
		}
		tsk.Id = maxId + 1
	}

	tasks = append(tasks, tsk)
	SaveTasks(tasks)
	Println("Задача добавлена")
}

func DoneTask(tasks []Task, id string) {
	if len(tasks) == 0 {
		color.Red("Список задач пуст")
		return
	}
	if id == "" {
		color.Red("Ошибка: нужно указать ID задачи")
		return
	}
	idStr, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	for i, t := range tasks {
		if t.Id == idStr {
			tasks[i].EndDate = time.Now().Format("2006-01-02 15:04:05")
			tasks[i].IsDone = true
			SaveTasks(tasks)
			color.Green("Задача отмечена выполненной")
			break
		}
	}
}

func DeleteTask(tasks []Task, id string) {
	if len(tasks) == 0 {
		Println("Список задач пуст")
		return
	}
	if id == "" {
		Println("Ошибка: нужно указать ID задачи")
		return
	}
	idStr, _ := strconv.Atoi(id)
	for i, t := range tasks {
		if t.Id == idStr {
			tasks = append(tasks[:i], tasks[i+1:]...)
			SaveTasks(tasks)
			Println("Задача удалена")
			break
		}
	}
}

func ClearSystem() {
	Println("Вы уверены что хотите очистить файл? (y/n)")
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	if scanner.Text() == "y" || scanner.Text() == "yes" {
		var tasks = []Task{}
		SaveTasks(tasks)

		Println("Документ успешно очищен")
		return
	}

	Println("Удаление файла отменено")
}

//добавить экспорт статистику и редактирование
