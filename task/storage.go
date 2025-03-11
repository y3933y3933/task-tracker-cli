package task

import (
	"encoding/json"
	"fmt"
	"os"
)

const filename = "tasks.json"

func LoadTask() ([]Task, error) {
	dat, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return []Task{}, err
	}

	var tasks []Task

	err = json.Unmarshal(dat, &tasks)
	if err != nil {
		fmt.Printf("解析 json 失敗: %v\n", err)
		return []Task{}, err
	}
	return tasks, nil

}

func SaveTask(tasks []Task) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	return encoder.Encode(tasks)

}
