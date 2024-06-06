package main

import "fmt"

type Project struct {
	Name        string
	Description string
	Directory   string
	Tasks       []Task
}

type Task struct {
	Name        string
	Description string
	Status      string
	Urgency     string
}

func main() {
	fmt.Println("GO Project Manager")
}
