package build

import (
	"bob/cflag"
	"bob/source"
)

type Kind int

const (
	Library    Kind = 0
	Executable Kind = 1
	Module     Kind = 2
)

type Context struct {
	Sources   []source.Context
	CFlags    cflag.Group
	DepCflags cflag.Group
	Name      string
	Kind      Kind
}
