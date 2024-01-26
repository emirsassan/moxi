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

func RunGroup(group Group) {
	fmt.Printf("Running Group: %s\n", group.Name)

	// Run each task in the group.
	for _, task := range group.Tasks {
		err := RunTask(task)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}

	// Group completion message.
	fmt.Printf("Success: Group completed: %s\n", group.Name)
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
	regex := regexp.MustCompile(`^\[group\s*([^\]]+)\]\s*$`)
	matches := regex.FindStringSubmatch(input)

	if len(matches) != 2 {
		return Group{}, fmt.Errorf("Invalid group start")
	}

	groupName := strings.TrimSpace(matches[1])

	group := Group{Name: groupName}
	return group, nil
}
