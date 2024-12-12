package main

import "fmt"

// Implementation for GeneralTask
func (g GeneralTask) Display() {
	fmt.Printf("General task: %s\nDescription: %s\nEstimated Hours: %d\n", g.Title, g.Description, g.Hours)
}

func (g GeneralTask) GetHours() int {
	return g.Hours
}

// Implementation for DevelopmentTask
func (d DevelopmentTask) Display() {
	fmt.Printf("Development task: %s\nCode Lines: %d\nEstimated Hours: %d\n", d.Title, d.CodeLines, d.Hours)
}

func (d DevelopmentTask) GetHours() int {
	return d.Hours
}

// Implementation for TestingTask
func (t TestingTask) Display() {
	fmt.Printf("Testing Task: %s\nTests to Run: %d\nEstimated Hours: %d\n", t.Title, t.TestsToRun, t.Hours)
}

func (t TestingTask) GetHours() int {
	return t.Hours
}
