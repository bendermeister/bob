package config

import (
	"bob/flag"
)

type Config struct {
	Name      string
	MainBuild string
	FlagGroup map[string]flag.Set
}
