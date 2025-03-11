package command

import (
	"errors"
	"fmt"
	"slices"

	"github.com/y3933y3933/task-tracker-cli/task"
)

func handlerDelete(args []string) error {
	if len(args) != 1 {
		return errors.New("usage: task-cli delete <id>")
	}

	id := args[0]
	tasks, err := task.LoadTask()
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Print("No tasks found")
	}

	tasks = slices.DeleteFunc(tasks, func(t task.Task) bool {
		return t.ID.String() == id
	})

	return task.SaveTask(tasks)

}
