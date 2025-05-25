package cmd

import (
	"flag" // 2. Флаги командной строки для разбора опций.
	"fmt"  // 3. Форматированный вывод в stdout/stderr.
	"os"   // 4. Работа с окружением и завершение программы.
	"time" // 5. Работа с датой/временем (фильтрация по today).

	"github.com/Kirill-Pinyaev/godo/internal" // 6. Внутренний пакет с моделями и хранилищем.
)

// list — реализует команду `godo list [--all] [--today]`.
func list(args []string, filePath string) {
	fs := flag.NewFlagSet("list", flag.ExitOnError)

	showAll := fs.Bool("all", false, "показывать все задачи, включая выполненные")
	showToday := fs.Bool("today", false, "показывать только задачи, созданные сегодня")

	fs.Parse(args)

	tasks, err := internal.LoadTasks(filePath)
	if err != nil {
		if err == internal.ErrNoTasksFile {
			fmt.Println("список задач пуст")
			return
		}
		fmt.Fprintf(os.Stderr, "ошибка загрузки задач: %v\n", err)
		os.Exit(1)
	}

	today := time.Now().UTC().Format("2006-01-02")

	for _, t := range tasks {
		if !*showAll && t.Done {
			continue
		}

		if *showToday && t.Created.UTC().Format("2006-01-02") != today {
			continue
		}

		status := " "
		if t.Done {
			status = "x"
		}

		fmt.Printf("[%s] %d: %s (%s)\n", status, t.ID, t.Text, t.Created.UTC().Format(time.RFC3339))
	}
}
