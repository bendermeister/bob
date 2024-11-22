package cflag

import (
	"fmt"
	"slices"
	"strings"
)

type Flag struct {
	body map[string]bool
}

func (f *Flag) Set(flag string) {
	f.body[flag] = true
}

func normalise(f string) string {
	return strings.TrimSpace(f)
}

func check(f string) error {
	if f[0] != '-' {
		return fmt.Errorf("Flag: '%s' has wrong format: expected '-' as first character.", f)
	}
	if strings.Contains(f, " ") {
		return fmt.Errorf("Flag: '%s' contains whitespace.", f)
	}
	return nil
}

func New(f []string) (Flag, error) {
	flag := Flag{}
	flag.body = make(map[string]bool)
	for _, g := range f {
		g = normalise(g)
		err := check(g)
		if err != nil {
			return flag, err
		}
		flag.Set(g)
	}

	return flag, nil
}

func (f Flag) Format() string {
	if f.body == nil {
		return ""
	}
	s := make([]string, len(f.body))
	slices.Sort(s)

	str := ""
	for _, f := range s {
		str += f + " "
	}
	return str
}
