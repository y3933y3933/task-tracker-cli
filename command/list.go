package command

import (
	"errors"
	"fmt"

	"github.com/y3933y3933/task-tracker-cli/task"
)

func handlerList(args []string) error {
	if len(args) > 1 {
		return errors.New("usage: task-cli list <todo | done | in-progress>")
	}

	tasks, err := task.LoadTask()
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return nil
	}
	if len(args) == 0 {
		printTask(tasks)
		return nil
	}

	status := args[0]

	if isValidStatus(status) {
		filteredTasks := filterByStatus(tasks, status)
		printTask(filteredTasks)

	} else {
		return errors.New("usage: task-cli list <todo | done | in-progress>")
	}

	return nil
}

func printTask(tasks []task.Task) {

	for _, task := range tasks {
		fmt.Printf("%v %s %s\n", task.ID, task.Description, task.Status)

	}
}

func filterByStatus(tasks []task.Task, status string) []task.Task {

	result := []task.Task{}
	for _, task := range tasks {
		if task.Status == status {
			result = append(result, task)
		}
	}

	return result

}

func isValidStatus(status string) bool {
	validStatuses := map[string]bool{
		"todo":        true,
		"in-progress": true,
		"done":        true,
	}
	_, exist := validStatuses[status]
	return exist
}
