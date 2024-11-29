package compiler

type Build struct{}

type Compiler interface {
	Compile(path string, flag string, object string) error
	Link(t string, objects []string) error
}
