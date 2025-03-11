package command

import (
	"errors"
	"fmt"
	"slices"
	"time"

	"github.com/google/uuid"
	"github.com/y3933y3933/task-tracker-cli/task"
)

func handlerUpdate(args []string) error {
	if len(args) != 2 {
		return errors.New("usage: task-cli update <id> <description>")
	}

	tasks, err := task.LoadTask()

	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Print("No tasks found")
		return nil
	}

	idString := args[0]
	id, err := uuid.Parse(idString)
	if err != nil {
		return err
	}

	newDescription := args[1]

	i := slices.IndexFunc(tasks, func(t task.Task) bool {
		return t.ID.String() == idString
	})

	updatedTask := task.Task{
		ID:          id,
		Description: newDescription,
		UpdatedAt:   time.Now().UTC(),
		CreatedAt:   tasks[i].CreatedAt,
		Status:      tasks[i].Status,
	}

	tasks[i] = updatedTask

	return task.SaveTask(tasks)
}
