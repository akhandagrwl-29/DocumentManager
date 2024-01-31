package model

type CMD interface {
	Execute(command *Command) error
}

type Command struct {
	CommandName string
	Arguments   interface{}
	Connection  CMD
}
