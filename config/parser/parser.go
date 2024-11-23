package parser

import (
	"bob/build"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Kind struct {
	Kind build.Kind
}

func (k *Kind) UnmarshalText(text []byte) error {
	str := string(text)
	if str == "exe" || str == "executable" {
		k.Kind = build.Executable
		return nil
	}
	if str == "mod" || str == "module" {
		k.Kind = build.Module
		return nil
	}
	if str == "lib" || str == "library" {
		k.Kind = build.Library
		return nil
	}
	return fmt.Errorf("Kind: '%s' does not match any of: ['exe', 'executable', 'mod', 'module', 'lib', 'library']", str)
}

type FlagGroup struct {
	Name       string   `toml:"Name"`
	Flags      []string `toml:"Flags"`
	FlagGroup  string   `toml:"FlagGroup"`
	FlagGroups []string `toml:"FlagGroups"`
}

type Build struct {
	Name       string   `toml:"Name"`
	Kind       Kind     `toml:"Type"`
	FlagGroups []string `toml:"FlagGroups"`
	Target     string   `toml:"Target"`
	FlagGroup  string   `toml:"FlagGroup"`
	Build      string   `toml:"Build"`
}

type Source struct {
	Path  string   `toml:"Path"`
	Flags []string `toml:"Flags"`
}

type Config struct {
	ProjectName string      `toml:"ProjectName"`
	MainBuild   string      `toml:"MainBuild"`
	Builds      []Build     `toml:"Build"`
	FlagGroups  []FlagGroup `toml:"FlagGroup"`
	Sources     []Source    `toml:"Source"`
}

func Run(path string) (Config, error) {
	var config Config

	text, err := os.ReadFile(path)
	if err != nil {
		return config, err
	}

	_, err = toml.Decode(string(text), &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
