package command

import (
	"fmt"

	"github.com/wychl/gocli/command"
	"github.com/wychl/gocli/command/cron"
	"github.com/wychl/gocli/command/middleware"
)

var commandDict = map[string]command.Command{}

// Register 注册command
func Register(list ...command.Command) {
	for _, cmd := range list {
		cmd = middleware.NewLogger(cmd)
		name := cmd.Name()
		commandDict[name] = cmd
	}
}

// Get 获取 command
func Get(name string) command.Command {
	if cmd := commandDict[name]; cmd == nil {
		panic(fmt.Sprintf("%s not found", name))
	}
	return commandDict[name]
}

func init() {
	Register(&cron.Cron{}) // cron表达式
}
