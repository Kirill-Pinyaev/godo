package cmd

import (
	"fmt"     // 2. Для форматированного вывода.
	"os"      // 3. Для завершения программы с кодом.
	"strconv" // 4. Для парсинга строки в int (ID задачи).

	"github.com/Kirill-Pinyaev/godo/internal" // 5. Модель Task и функции Load/Save.
)

// done — реализует команду `godo done <taskID>`.
func done(args []string, filePath string) {
	if len(args) < 1 {
		fmt.Println("usage: godo done <task ID>")
		os.Exit(1)
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "некорректный ID '%s': %v\n", args[0], err)
		os.Exit(1)
	}

	tasks, err := internal.LoadTasks(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ошибка загрузки задач: %v\n", err)
		os.Exit(1)
	}

	found := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = true
			found = true
			break
		}
	}

	if !found {
		fmt.Fprintf(os.Stderr, "задача с ID %d не найдена\n", id)
		os.Exit(1)
	}

	if err := internal.SaveTasks(filePath, tasks); err != nil {
		fmt.Fprintf(os.Stderr, "ошибка сохранения задач: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("задача %d отмечена как выполненная\n", id)
}
