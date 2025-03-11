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

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
	}

	for _, task := range tasks {

		fmt.Printf("%v %s %s\n", task.ID, task.Description, task.Status)

	}
	return err
}
