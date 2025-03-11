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
		"delete": {
			Name:     "delete",
			Callback: handlerDelete,
		},
		"update": {
			Name:     "update",
			Callback: handlerUpdate,
		},
		"mark-in-progress": {
			Name:     "mark-in-progress",
			Callback: handlerMarkInProgress,
		},
		"mark-done": {
			Name:     "mark-done",
			Callback: handlerMarkDone,
		},
	}
}
