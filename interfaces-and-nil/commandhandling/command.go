package commandhandling

type Command interface {
	CommandType() string
}
