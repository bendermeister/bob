package build

import "bob/flag"

type Ctx struct {
	Name   string
	Flag   flag.Set
	Kind   Kind
	Target string
}
