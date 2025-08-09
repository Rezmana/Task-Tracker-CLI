package main

import (
	"fmt"
	"os"

)

func main(){
	if len(os.Args)	< 2 {
		fmt.Println("Usage: go run main.go <command>")
		return
	}

	command := os.Args[1]
	// Handle the command
	// tasks, err := loadTasks()
	tasks, _ := loadTasks()

	switch command{
	case "add":
        if len(os.Args) < 3 {
            fmt.Println("Please provide a task title.")
            return
        }
        title := os.Args[2]
        id := getNextID(tasks)
        tasks = append(tasks, Task{ID: id, Title: title, Status: "incomplete"})
        saveTasks(tasks)
        fmt.Println("Task added.")

	case "deleteTask":
		
	}

}

func getNextID(tasks []Task) int {
	maxID := 0
	for _, t := range tasks{
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}
// func main(){
// 	if len(os.Args) < 2 {
// 		fmt.Println("Usage: go run main.go <task_number>")
// 		return
// 	}

// 	command := os.Args[1]

// 	tasks, err := loadTasks()

// 	switch command {
// 	case "saveTasks":
// 		if err := saveTasks(tasks); err != nil {
// 			fmt.Printf("Error saving tasks: %v\n", err)
// 		} else {
// 			fmt.Println("Tasks saved successfully.")
// 		}
// 	case "loadTasks":
// 		if err != nil {
// 			fmt.Printf("Error loading tasks: %v\n", err)
// 		} else {
// 			fmt.Println("Tasks loaded successfully.")
// 			for i, task := range tasks {
// 				fmt.Printf("%d: %s - %s\n", i+1, task.Title, task.Status)
// 			}
// 		}
// 	case "deleteTasks":
// 		if err := deleteTasks(); err != nil {
// 			fmt.Printf("Error deleting tasks: %v\n", err)
// 		} else {
// 			fmt.Println("Tasks deleted successfully.")
// 		}
// 	case "updateTask":
// 		if len(os.Args) < 4 {
// 			fmt.Println("Usage: go run main.go updateTask <task_title> <new_status>")
// 			return
// 		}
// 		title := os.Args[2]
// 		newStatus := os.Args[3]
// 		tasks, err = updateTask(tasks, title, newStatus)
// 		if err != nil {
// 			fmt.Printf("Error updating task: %v\n", err)
// 		} else {
// 			if err := saveTasks(tasks); err != nil {
// 				fmt.Printf("Error saving updated tasks: %v\n", err)
// 			} else {
// 				fmt.Println("Task updated successfully.")
// 			}
// 		}
// 	case "addTask":
// 		if len(os.Args) < 4 {
// 			fmt.Println("Usage: go run main.go addTask <task_title> <task_status>")
// 			return
// 		}
// 		title := os.Args[2]
// 		status := os.Args[3]
// 		newTask := Task{Title: title, Status: status}
// 		tasks, err = addTask(tasks, newTask)
// 		if err != nil {
// 			fmt.Printf("Error adding task: %v\n", err)
// 		} else {
// 			if err := saveTasks(tasks); err != nil {
// 				fmt.Printf("Error saving tasks after adding: %v\n", err)
// 			} else {
// 				fmt.Println("Task added successfully.")
// 			}
// 		}
// 	default:
// 		fmt.Println("Unknown command:", command)
// 		fmt.Println("Available commands: saveTasks, loadTasks, deleteTasks, updateTask, addTask")
// 	}
// }

