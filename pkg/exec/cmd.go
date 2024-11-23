package exec

import (
	"fmt"
)

type Command struct {
	Name string
	Args map[string]string
}

func NewCommand(name string) *Command {
	return &Command{Name: name, Args: make(map[string]string)}
}

func (c *Command) Exec(eCtx ExecCtx) (ExecCtx, error) {
	switch c.Name {
	case "fetch":
		return fetch(eCtx, c.Args)
	default:
		return eCtx, fmt.Errorf("%q is an invalid command", c.Name)
	}
}
