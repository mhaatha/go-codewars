package main

import "fmt"

// Display all tasks's detail
func PrintTasks(tasks []Task) {
	for _, task := range tasks {
		task.Display()
		fmt.Println("---")
	}
}

// Count the total duration of all tasks
func CalculateTotalHours(tasks []Task) int {
	total := 0
	for _, task := range tasks {
		total += task.GetHours()
	}
	return total
}
