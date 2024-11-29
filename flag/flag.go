package flag

import (
	"strings"
)

type Set struct {
	body []string
}

func (m *Set) Insert(s string) error {
	if m.body == nil {
		m.body = make([]string, 0)
	}
	m.body = append(m.body, s)
	return nil
}

func (m *Set) String() string {
	set := make(map[string]bool)
	for _, f := range m.body {
		set[f] = true
	}

	str := ""

	for f, _ := range set {
		str += f + " "
	}
	str = strings.TrimSpace(str)
	return str
}
