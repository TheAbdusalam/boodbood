package main

import "strconv"

type DirectoryNode struct {
	path string
	weight int
}

func (d *DirectoryNode) String() string {
	return d.path + " " + strconv.Itoa(d.weight)
}
