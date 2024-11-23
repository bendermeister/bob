package cflag

import (
	"fmt"
	"strings"
)

type Group struct {
	m map[string]bool
}

func (b *Group) Init() {
	if b.m == nil {
		b.m = make(map[string]bool)
	}
}

func (b *Group) Clear() {
	b.m = nil
}

func (b *Group) Set(f string) {
	b.Init()
	b.m[f] = true
}

func (b Group) String() string {
	str := "["
	for flag, _ := range b.m {
		str += flag + ", "
	}
	str += "]"
	return str
}

func (b *Group) UnmarshalText(text []byte) error {
	b.Init()
	argv := strings.Split(string(text), " ")

	// TODO: validate flag against known flags of clang
	// TODO: check if flag needs a parameter and parse it in
	for _, flag := range argv {
		if flag[0] != '-' {
			return fmt.Errorf("Flag: '%s', does not start with '-'.", flag)
		}
		b.Set(flag)
	}
	return nil
}
