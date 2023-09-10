package command

// Command command接口
type Command interface {
	// Name command name
	Name() string
	// Run 运行command
	Run(interface{}) error
}
