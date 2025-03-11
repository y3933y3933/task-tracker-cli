package command

import (
	"errors"
	"fmt"
	"slices"
	"time"

	"github.com/google/uuid"
	"github.com/y3933y3933/task-tracker-cli/task"
)

func handlerMarkInProgress(args []string) error {
	if len(args) != 1 {
		return errors.New("usage: task-cli mark-in-progress <id>")
	}

	tasks, err := task.LoadTask()
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return nil
	}

	idString := args[0]
	id, err := uuid.Parse(idString)
	if err != nil {
		return err
	}

	i := slices.IndexFunc(tasks, func(t task.Task) bool {
		return t.ID.String() == idString
	})

	updatedTask := task.Task{
		ID:          id,
		Description: tasks[i].Description,
		UpdatedAt:   time.Now().UTC(),
		CreatedAt:   tasks[i].CreatedAt,
		Status:      "in-progress",
	}
	tasks[i] = updatedTask

	return task.SaveTask(tasks)
}

func handlerMarkDone(args []string) error {
	if len(args) != 1 {
		return errors.New("usage: task-cli mark-in-progress <id>")
	}

	tasks, err := task.LoadTask()
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return nil
	}

	idString := args[0]
	id, err := uuid.Parse(idString)
	if err != nil {
		return err
	}

	i := slices.IndexFunc(tasks, func(t task.Task) bool {
		return t.ID.String() == idString
	})

	updatedTask := task.Task{
		ID:          id,
		Description: tasks[i].Description,
		UpdatedAt:   time.Now().UTC(),
		CreatedAt:   tasks[i].CreatedAt,
		Status:      "done",
	}
	tasks[i] = updatedTask

	return task.SaveTask(tasks)
}
