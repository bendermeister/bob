package build

type Kind int

const (
	Library    Kind = 0
	Executable Kind = 1
	Module     Kind = 2
)
