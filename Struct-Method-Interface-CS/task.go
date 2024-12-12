package main

// Struct for general task
type GeneralTask struct {
	Title       string
	Description string
	Hours       int
}

// Struct for development task
type DevelopmentTask struct {
	Title     string
	CodeLines int // Lines of code that must be written
	Hours     int
}

// Struct for testing task
type TestingTask struct {
	Title      string
	TestsToRun int // Amount of test that must be run
	Hours      int
}

// Interface for all the task type
type Task interface {
	Display()
	GetHours() int
}
