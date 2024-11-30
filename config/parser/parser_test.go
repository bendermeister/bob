package parser_test

import (
	"bob/config/parser"
	"slices"
	"testing"
)

type Validator struct {
	Bool bool
}

func (v *Validator) Add(b bool) *Validator {
	v.Bool = v.Bool && b
	return v
}

func (v *Validator) Check() bool {
	return v.Bool
}

func (v *Validator) Init() {
	v.Bool = true
}

func TestRun(t *testing.T) {
	path := "../bob.toml"
	config, err := parser.ParseFile(path)
	if err != nil {
		t.Fatal(err)
	}

	flagGroups := make(map[string]parser.FlagGroup)
	for _, fg := range config.FlagGroups {
		flagGroups[fg.Name] = fg
	}

	builds := make(map[string]parser.Build)
	for _, b := range config.Builds {
		builds[b.Name] = b
	}

	var fg parser.FlagGroup

	v := new(Validator)

	v.Init()
	fg = flagGroups["general"]

	v.Add(fg.Name == "general")
	v.Add(fg.Flags[0].String() == "-Wall")
	v.Add(fg.Flags[1].String() == "-Wextra")
	v.Add(fg.Flags[2].String() == "-std=c99")
	v.Add(fg.FlagGroup == "")
	v.Add(fg.FlagGroups == nil)
	if !v.Bool {
		t.Fatal("FlagGroup general incorecctly parsed")
	}

	v.Init()
	fg = flagGroups["debug"]

	v.Add(fg.Name == "debug")
	v.Add(slices.Equal(fg.FlagGroups, []string{"general"}))
	v.Add(fg.Flags[0].String() == "-g")
	v.Add(fg.FlagGroup == "")
	if !v.Bool {
		t.Fatal("FlagGroup debug incorrectly parsed")
	}

	v.Init()
	fg = flagGroups["release"]

	v.Add(fg.Name == "release")
	v.Add(fg.FlagGroup == "general")
	v.Add(fg.Flags[0].String() == "-O3")
	v.Add(fg.Flags[1].String() == "-march=native")
	v.Add(fg.Flags[2].String() == "-mtune=native")
	v.Add(slices.Equal(fg.FlagGroups, []string{}))
	if !v.Bool {
		t.Fatal("FlagGroup release incorrectly parsed")
	}
}
