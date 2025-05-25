package internal

import "time"

type Task struct {
	ID      int       `json:"id"`      // уникальный идентификатор
	Text    string    `json:"text"`    // текст задачи
	Created time.Time `json:"created"` // время создания (UTC)
	Done    bool      `json:"done"`    // флаг выполненности
}

func NextID(task []Task) int {
	max := 0
	for _, t := range task {
		if t.ID > max {
			max = t.ID
		}
	}
	return max + 1
}

func CreateTask(id int, text string) Task {
	return Task{
		ID:      id,
		Text:    text,
		Created: time.Now().UTC(),
		Done:    false,
	}
}
