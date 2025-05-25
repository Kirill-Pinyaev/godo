package internal

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

var ErrNoTasksFile = errors.New("tasks file not found")

func LoadTasks(path string) ([]Task, error) {
	data, err := ioutil.ReadFile(path)

	if os.IsNotExist(err) {
		return nil, ErrNoTasksFile
	}

	if err != nil {
		return nil, err
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func SaveTasks(path string, tasks []Task) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, data, 0o644)
}
