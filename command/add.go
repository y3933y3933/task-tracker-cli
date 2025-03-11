package command

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/y3933y3933/task-tracker-cli/task"
)

func handlerAdd(args []string) error {
	if len(args) != 1 {
		return errors.New("參數錯誤")
	}

	tasks, err := task.LoadTask()

	if err != nil {
		return err
	}

	description := args[0]
	newTask := task.Task{
		ID:          uuid.New(),
		Description: description,
		Status:      "todo",
		UpdatedAt:   time.Now().UTC(),
		CreatedAt:   time.Now().UTC(),
	}

	tasks = append(tasks, newTask)
	return task.SaveTask(tasks)
}
