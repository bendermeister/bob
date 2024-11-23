package source

import (
	"bob/cflag"
)

type Context struct {
	Path   string
	CFlags cflag.Group
}
