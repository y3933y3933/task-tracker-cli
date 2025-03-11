package command

type Command struct {
	Name     string
	Callback func(args []string) error
}

func GetCommands() map[string]Command {
	return map[string]Command{
		"add": {
			Name:     "add",
			Callback: handlerAdd,
		},
	}
}
