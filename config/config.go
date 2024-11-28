package config

import (
	"bob/build"
	"bob/flag"
)

type Config struct {
	Name      string
	MainBuild string
	FlagGroup map[string]flag.Set
	Source    map[string]build.Source
}

func New(path string) (Config, error) {
}
