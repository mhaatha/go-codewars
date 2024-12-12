package main

import "fmt"

func main() {
	// Create task list
	tasks := []Task{
		GeneralTask{Title: "Team Meeting", Description: "Discuss project requirements", Hours: 2},
		DevelopmentTask{Title: "Feature Implementation", CodeLines: 500, Hours: 10},
		TestingTask{Title: "Unit Testing", TestsToRun: 20, Hours: 5},
	}

	fmt.Println("Task Lists:")
	PrintTasks(tasks)

	totalHours := CalculateTotalHours(tasks)
	fmt.Printf("Total Estimated Hours: %d\n", totalHours)
}
