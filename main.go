package main

import (
	"fmt"
	"os"

	"github.com/y3933y3933/task-tracker-cli/command"
)

const cliName = "task-cli"

func main() {
	if len(os.Args) < 3 || os.Args[1] != cliName {
		fmt.Println("Usage: task-cli COMMAND [OPTIONS]")
		return
	}

	commandName := os.Args[2]
	args := os.Args[3:]

	command, exist := command.Commands()[commandName]
	if !exist {
		fmt.Printf("Command \"%s\" Not Found", commandName)
		return
	}

	err := command.Callback(args)

	if err != nil {
		fmt.Printf("Command fail: %v\n", err)
	}

}
