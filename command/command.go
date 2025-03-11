package command

type Command struct {
	Name     string
	Callback func(args []string) error
}

func Commands() map[string]Command {
	return map[string]Command{
		"add": {
			Name:     "add",
			Callback: handlerAdd,
		},
		"list": {
			Name:     "list",
			Callback: handlerList,
		},
	}
}
