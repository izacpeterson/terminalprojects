package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

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
	scanner := bufio.NewScanner(os.Stdin)
	var input string

	fmt.Println("GO Project Manager")
	view := "projects"
	var selectedProject int
	var projects []Project
	projects = append(projects, demoProject())

	for {
		clear()

		switch view {
		case "projects":
			fmt.Println("Enter Project Number or Type 'new' to create a new project")

			listProjects(projects)

			fmt.Print("Action: ")
			scanner.Scan()
			input = scanner.Text()

			projectNum, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}
			selectedProject = projectNum
			view = "tasks"
		case "tasks":
			fmt.Println(projects[selectedProject].Name)
			fmt.Println(projects[selectedProject].Description)
			fmt.Println("Tasks:")
			listTasks(projects[selectedProject].Tasks)

			var input string
			fmt.Print("Action (i: new task): ")
			scanner.Scan()
			input = scanner.Text()

			switch input {
			case "i":
				projects[selectedProject].Tasks = append(projects[selectedProject].Tasks, newTask())
			}
		}
	}
}

func listProjects(projects []Project) {
	for i := 0; i < len(projects); i++ {
		fmt.Printf("%v: %v\n", i, projects[i].Name)
	}
}
func listTasks(tasks []Task) {
	for i := 0; i < len(tasks); i++ {
		fmt.Printf("%v: %v | %v | %v\n", i, tasks[i].Name, tasks[i].Urgency, tasks[i].Status)
	}
}
func newTask() Task {
	scanner := bufio.NewScanner(os.Stdin)
	var input string

	var newTask Task

	fmt.Print("Task Name: ")
	scanner.Scan()
	input = scanner.Text()

	newTask.Name = input

	var urgency string
	fmt.Print("Enter Urgency (1: Low, 2: Med, 3: High): ")
	fmt.Scanln(&urgency)

	switch urgency {
	case "1":
		newTask.Urgency = "Low"
	case "2":
		newTask.Urgency = "Medium"
	case "3":
		newTask.Urgency = "High"
	}

	newTask.Status = "New"

	return newTask
}

func demoProject() Project {
	project := Project{
		Name:        "Demo Project",
		Description: "This is a demo project",
		Directory:   "/home/izac/Documents/Dev/goPM",
		Tasks: []Task{
			{
				Name:        "Task 1",
				Description: "This is the first task",
				Status:      "Not Started",
				Urgency:     "High",
			},
			{
				Name:        "Task 2",
				Description: "This is the second task",
				Status:      "In Progress",
				Urgency:     "Medium",
			},
			{
				Name:        "Task 3",
				Description: "This is the third task",
				Status:      "Completed",
				Urgency:     "Low",
			},
		},
	}

	return project
}

func clear() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
