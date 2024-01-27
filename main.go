package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/emirsassan/moxi/src/file"
	"github.com/emirsassan/moxi/src/parser"
	"os"
	"strings"
)

var moxiFilePath string

func init() {
	flag.StringVar(&moxiFilePath, "file", "Moxifile", "Path to the Moxifile")
	flag.Parse()
}

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("  moxi [options] [task_name]")
	fmt.Println("\nOptions:")
	flag.PrintDefaults()
}

func checkMoxiFileExists() error {
	if _, err := os.Stat(moxiFilePath); os.IsNotExist(err) {
		return errors.New("Moxifile not found. Please make sure it exists in the specified path or the working directory")
	}
	return nil
}

func main() {
	helpFlag := flag.Bool("help", false, "Print this help message")
	flag.Parse()

	if *helpFlag {
		printHelp()
		return
	}

	if err := checkMoxiFileExists(); err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Read input from the file
	fileContent, err := file.ReadFile(moxiFilePath)
	if err != nil {
		fmt.Println("File reading error:", err)
		return
	}

	// Convert file content to string
	lines := strings.Split(string(fileContent), "\n")

	// Parse syntax using Group and Task structures
	var groups []parser.Group
	var tasks []parser.Task
	var currentGroup *parser.Group

	for _, line := range lines {
		// Check for group start
		if strings.HasPrefix(line, "group") {
			group, err := parser.ParseGroupSyntax(line)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			currentGroup = &group
			groups = append(groups, group)

			// Print group header
			fmt.Printf("Group Name: %s\n", group.Name)
		} else if strings.HasPrefix(line, "task") {
			task, err := parser.ParseTaskSyntax(line)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			// Add task to the current group
			if currentGroup != nil {
				currentGroup.Tasks = append(currentGroup.Tasks, task)
			} else {
				// If there is no group, add tasks to the basic tasks list
				tasks = append(tasks, task)
			}
		}
	}

	// CLI handling
	args := os.Args[1:] // Exclude the program name

	if len(args) == 0 {
		// Run all tasks if no task name is provided
		fmt.Println("Running all tasks:")
		for _, task := range tasks {
			err := parser.RunTask(task)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		}

		// Run group tasks
		for _, group := range groups {
			// Skip if the group is empty
			if len(group.Tasks) == 0 {
				continue
			}

			parser.RunGroup(group)
		}
	} else if len(args) == 1 {
		// Run a specific task if a task name is provided
		taskName := args[0]
		foundTask := false

		// Check basic tasks
		for _, task := range tasks {
			if task.Name == taskName {
				err := parser.RunTask(task)
				if err != nil {
					return
				}
				foundTask = true
				break
			}
		}

		// Check group tasks
		for _, group := range groups {
			for _, task := range group.Tasks {
				if task.Name == taskName {
					err := parser.RunTask(task)
					if err != nil {
						return
					}
					foundTask = true
					break
				}
			}
		}

		if !foundTask {
			fmt.Printf("Error: Task '%s' not found\n", taskName)
		}
	} else {
		// Display usage information if more than one argument is provided
		fmt.Println("Usage:")
		fmt.Println("  Run all tasks: moxi")
		fmt.Println("  Run a specific task: moxi <task_name>")
	}
}
