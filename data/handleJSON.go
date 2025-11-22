package data

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/renderer"
	"github.com/olekukonko/tablewriter/tw"
)

type Task struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

const dataFile = "data.json"

func readJSON() []Task {
	var tasks []Task

	file, err := os.ReadFile(dataFile)

	if err != nil {
		fmt.Println(err)
	}

	_ = json.Unmarshal(file, &tasks)

	return tasks
}

func writeJSON(tasks []Task) {

	data, err := json.MarshalIndent(tasks, "", " ")

	if err != nil {
		fmt.Println(err)
	}

	writeErr := os.WriteFile(dataFile, data, 0644)

	if writeErr != nil {
		fmt.Println(writeErr)
		os.Exit(1)
	}

}

func AddRecord(task string) {

	tasks := readJSON()

	tasks = append(tasks, Task{Title: task, Completed: false})

	writeJSON(tasks)

}

func RemoveRecord(id int) {
	tasks := readJSON()

	id = id - 1

	if id < 0 || id > len(tasks) {
		fmt.Println("Record ID out of range")
		os.Exit(1)
	}

	records := append(tasks[:id], tasks[id+1:]...)

	writeJSON(records)
}

func ListRecords() {

	tasks := readJSON()

	tableColorCfg := renderer.ColorizedConfig{
		Header: renderer.Tint{
			FG: renderer.Colors{color.FgYellow, color.Bold},
		},
		Column: renderer.Tint{
			Columns: []renderer.Tint{
				{FG: renderer.Colors{color.FgMagenta}},
				{FG: renderer.Colors{color.FgBlue}},
				{FG: renderer.Colors{color.FgRed}},
			},
		},
	}

	table := tablewriter.NewTable(os.Stdout,
		tablewriter.WithRenderer(renderer.NewColorized(tableColorCfg)),
		tablewriter.WithConfig(tablewriter.Config{
			Row: tw.CellConfig{
				Formatting:   tw.CellFormatting{AutoWrap: tw.WrapNormal},
				Alignment:    tw.CellAlignment{Global: tw.AlignLeft},
				ColMaxWidths: tw.CellWidth{Global: 30},
			},
		}),
	)

	table.Header(
		"ID",
		"Tasks",
		"Complete")

	for i, task := range tasks {
		status := "PENDING"
		if task.Completed {
			status = color.New(color.FgGreen).Sprint("DONE")
		}
		table.Append([]string{strconv.Itoa(i + 1), task.Title, status})
	}

	table.Render()

}

func ToggleRecord(id int) {

	tasks := readJSON()

	id = id - 1

	if id < 0 || id > len(tasks) {
		fmt.Println("Record ID out of range")
		os.Exit(1)
	}

	tasks[id].Completed = !tasks[id].Completed

	records := tasks

	writeJSON(records)
}

func UpdateRecord(id int, taskTitle string) {

	tasks := readJSON()

	id = id - 1

	if id < 0 || id > len(tasks) {
		fmt.Println("Record ID out of range")
		os.Exit(1)
	}

	tasks[id].Title = taskTitle

	records := tasks

	writeJSON(records)

}
