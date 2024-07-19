package main

import (
	"strconv"
	"strings"
)

type DirectoryNode struct {
	name   string
	path   string
	weight int
}

func (d *DirectoryNode) Name() string {
	return strings.Split(d.path, "/")[len(strings.Split(d.path, "/"))-1]
}

func (d *DirectoryNode) String() string {
	return d.path + " " + strconv.Itoa(d.weight)
}
