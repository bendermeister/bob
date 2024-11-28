package flag

import "strings"

// TODO: how to deal with arguments
type Single struct {
	body string
}

func (s *Single) UnmarshalText(text []byte) error {
	s.body = string(text)
	s.body = strings.TrimSpace(s.body)
	// TODO check if this is a real flag
	return nil
}

func (s Single) String() string {
	return s.body
}

type Set struct {
	body []Single
}

func (m *Set) Insert(s Single) error {
	if m.body == nil {
		m.body = make([]Single, 0)
	}
	m.body = append(m.body, s)
	return nil
}

func (m *Set) String() string {
	set := make(map[string]bool)
	for _, f := range m.body {
		set[f.String()] = true
	}

	str := ""

	for f, _ := range set {
		str += f + " "
	}
	str = strings.TrimSpace(str)
	return str
}
