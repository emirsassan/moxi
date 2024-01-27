package parser

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

type Task struct {
	Name    string
	Command string
}

type Group struct {
	Name  string
	Tasks []Task
}

func ParseTaskSyntax(input string) (Task, error) {
	// Task structure validation using regex.
	regex := regexp.MustCompile(`^task\s*:\s*(\w+)\s*=>\s*(.+)$`)
	matches := regex.FindStringSubmatch(input)

	if len(matches) != 3 {
		return Task{}, fmt.Errorf("Invalid task syntax")
	}

	taskName := strings.TrimSpace(matches[1])
	command := strings.TrimSpace(matches[2])

	return Task{Name: taskName, Command: command}, nil
}

func RunTask(task Task) error {
	fmt.Printf("Running task: %s\n", task.Name)

	cmd := exec.Command("cmd", "/C", task.Command)

	// Check command output.
	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed: %v\n", err)
		return err
	}

	// Task success message
	fmt.Println("Success: Task completed")
	return nil
}

func ParseGroupSyntax(input string) (Group, error) {
	// Group structure validation using regex.
	regex := regexp.MustCompile(`^group\s*([^\s]+)\s*\{\s*$`)
	matches := regex.FindStringSubmatch(input)

	if len(matches) != 2 {
		return Group{}, fmt.Errorf("Invalid group syntax")
	}

	groupName := strings.TrimSpace(matches[1])
	tasks, err := parseGroupTasks(input)
	if err != nil {
		return Group{}, err
	}

	group := Group{
		Name:  groupName,
		Tasks: tasks,
	}

	return group, nil
}

// parseGroupTasks parses the tasks within a group
func parseGroupTasks(input string) ([]Task, error) {
	var tasks []Task

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		// Extract task name and command using a regex or custom logic
		taskRegex := regexp.MustCompile(`^\s*task\s*:\s*([^\s]+)\s*=>\s*(.+)\s*$`)
		taskMatches := taskRegex.FindStringSubmatch(line)

		if len(taskMatches) == 3 {
			taskName := strings.TrimSpace(taskMatches[1])
			taskCommand := strings.TrimSpace(taskMatches[2])

			tasks = append(tasks, Task{
				Name:    taskName,
				Command: taskCommand,
			})
		}
	}

	return tasks, nil
}

// RunGroup runs the tasks within a group
func RunGroup(group Group) {
	fmt.Printf("Running tasks for group '%s':\n", group.Name)
	for _, task := range group.Tasks {
		err := RunTask(task)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
