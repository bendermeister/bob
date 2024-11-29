package main

import "flag"

type Source struct {
	path string
	flag flag.FlagSet
}

func (s Source) Path() string {
	return s.path
}

func New(path string) Source {
	return Source{
		path: path,
	}
}
