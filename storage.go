package main

import (
	"fmt"
	"encoding/json"
	"os"
)

const fileName = "tasks.json"

func saveTasks(tasks []Task) error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(tasks); err != nil {
		return fmt.Errorf("error encoding tasks: %w", err)
	}

	return nil
}

func loadTasks() ([]Task, error) {
	file, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil // Return empty slice if file does not exist
		}
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var tasks []Task
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		return nil, fmt.Errorf("error decoding tasks: %w", err)
	}

	return tasks, nil
}

func deleteTasks() error {
	if err := os.Remove(fileName); err != nil {
		if os.IsNotExist(err) {
			return nil // File does not exist, nothing to delete
		}
		return fmt.Errorf("error deleting file: %w", err)
	}
	return nil
}

func updateTask(tasks []Task, title string, newStatus string) ([]Task, error) {
	for i, task := range tasks {
		if task.Title == title {
			tasks[i].Status = newStatus
			return tasks, nil
		}
	}
	return nil, fmt.Errorf("task with title '%s' not found", title)
}

func addTask(tasks []Task, newTask Task) ([]Task, error) {
	for _, task := range tasks {
		if task.Title == newTask.Title {
			return nil, fmt.Errorf("task with title '%s' already exists", newTask.Title)
		}
	}
	tasks = append(tasks, newTask)
	return tasks, nil
}

