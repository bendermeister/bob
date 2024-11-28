package compiler

type FlagSet struct{}
type Flag struct{}
type Object struct{}
type Target struct{}

type Compiler interface {
	Compile(path string, flag string, object string) error
	Link(t string, objects []string) error
}
