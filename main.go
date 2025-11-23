package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/pkdiv/taskman/data"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: taskman command")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: taskman add <task>")
			os.Exit(1)
		}
		task := os.Args[2]
		data.AddRecord(task)

	case "rm":
		if len(os.Args) < 3 {
			fmt.Println("Usage: taskman rm <id>")
			os.Exit(1)
		}

		id, err := strconv.Atoi(os.Args[2])

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		data.RemoveRecord(id)

	case "list":

		data.ListRecords()

	case "toggle":
		if len(os.Args) < 3 {
			fmt.Println("Usage: taskman toggle <id>")
			os.Exit(1)
		}

		id, err := strconv.Atoi(os.Args[2])

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		data.ToggleRecord(id)

	case "update":

		if len(os.Args) < 4 {
			fmt.Println("Usage: taskman update <id> <task title>")
			os.Exit(1)
		}

		id, err := strconv.Atoi(os.Args[2])

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		data.UpdateRecord(id, os.Args[3])
	default:
		fmt.Println("Command:", command, "Not Present")
	}

}
