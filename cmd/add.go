package cmd

import (
	"fmt"
	"os"

	"github.com/Kirill-Pinyaev/godo/internal"
)

// add — реализует команду `godo add <taskText>`.
// args     — всё, что пользователь передал после слова "add".
// filePath — путь к файлу с задачами, например ~/.godo/tasks.json.
func add(args []string, filePath string) {
	if len(args) < 1 {
		fmt.Println("usage: godo add <task text>")
		os.Exit(1)
	}

	text := args[0]

	tasks, err := internal.LoadTasks(filePath)
	if err != nil && err != internal.ErrNoTasksFile {
		fmt.Fprintf(os.Stderr, "error loading tasks: %v\n", err)
		os.Exit(1)
	}

	if err == internal.ErrNoTasksFile {
		tasks = []internal.Task{}
	}

	newID := internal.NextID(tasks)

	task := internal.CreateTask(newID, text)

	tasks = append(tasks, task)

	if err := internal.SaveTasks(filePath, tasks); err != nil {
		fmt.Fprintf(os.Stderr, "error saving tasks: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("added task %d: %s\n", task.ID, task.Text)
}
