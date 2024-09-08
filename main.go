package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

type Todo struct {
	task   string
	completed bool
	day time.Time
}

func main() {
	fmt.Println("Hello Gofers.. 👽👽")

	fmt.Println("Welcome to CLI Task Manager\nPlease select an option:\n\t1. Add todo\n\t2. View todo\n\t3. Delete todo\n\t4. Update todo")

	var option int

	fmt.Scanln(&option)

	var todo= []Todo{
		{
			task:"Complete CLI_TODO",
			completed: false,
			day: time.Now(),
		},
	}

	for {
		if option == 1 {
			color.Green("Please enter the task to do")
			break
		} else if option == 2 {
			color.Green("Here's your task list")
			fmt.Println(todo)
			break
		} else if option == 3 {
			color.Green("Select task to delete")
			break
		} else if option == 4 {
			color.Green("Select task to update")
			break
		} else {
			break
		}
	}

}
