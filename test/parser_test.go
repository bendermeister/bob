package parser

import (
	"bob/build"
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
	path := "example.toml"
	config, err := parser.Run(path)
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

	sources := make(map[string]parser.Source)
	for _, s := range config.Sources {
		sources[s.Path] = s
	}

	var fg parser.FlagGroup

	v := new(Validator)

	v.Init()
	v.Add(config.ProjectName == "hello")
	v.Add(config.MainBuild == "debug")
	if !v.Bool {
		t.Fatal("general inforation incorrectly parsed")
	}

	v.Init()
	fg = flagGroups["general"]

	v.Add(fg.Name == "general")
	v.Add(slices.Equal(fg.Flags, []string{"-Wall", "-Wextra", "-std=c99"}))
	v.Add(fg.FlagGroup == "")
	v.Add(fg.FlagGroups == nil)
	if !v.Bool {
		t.Fatal("FlagGroup general incorecctly parsed")
	}

	v.Init()
	fg = flagGroups["debug"]

	v.Add(fg.Name == "debug")
	v.Add(slices.Equal(fg.FlagGroups, []string{"general"}))
	v.Add(slices.Equal(fg.Flags, []string{"-g"}))
	v.Add(fg.FlagGroup == "")
	if !v.Bool {
		t.Fatal("FlagGroup debug incorrectly parsed")
	}

	var b parser.Build

	v.Init()
	b = builds["general"]
	v.Add(b.Name == "general")
	v.Add(b.Kind.Kind == build.Executable)
	v.Add(b.Target == "hello.out")
	v.Add(b.FlagGroup == "general")
	v.Add(b.FlagGroups == nil)
	v.Add(b.Build == "")
	if !v.Bool {
		t.Fatal("Build general incorrectly parsed")
	}

	v.Init()
	b = builds["debug"]
	v.Add(b.Name == "debug")
	v.Add(b.Build == "general")
	v.Add(b.FlagGroup == "debug")
	v.Add(b.FlagGroups == nil)
	if !v.Bool {
		t.Fatal("Build debug incorrectly parsed")
	}

	var s parser.Source
	v.Init()
	s = sources["main.c"]

	v.Add(s.Path == "main.c")
	v.Add(slices.Equal(s.Flags, []string{"-DEEZNUTS=1"}))
	if !v.Bool {
		t.Fatal("Source main.c incorrectly parsed")
	}

	v.Init()
	s = sources["*.c"]
	v.Add(s.Path == "*.c")
	v.Add(s.Flags == nil)
	if !v.Bool {
		t.Fatal("Source *.c incorrectly parsed")
	}

}
