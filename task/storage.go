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
		fmt.Printf("讀取 %s 失敗: %v\n", filename, err)
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
