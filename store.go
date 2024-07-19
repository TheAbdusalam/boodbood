package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Store struct {
	storeFile *os.File
	nodes     []DirectoryNode
}

func NewStore() *Store {
	file, err := os.OpenFile(storeFilePath, os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		_ = fmt.Errorf("could not open file %s: %v", storeFilePath, err)
		return nil
	}

	s := &Store{
		storeFile: file,
	}

	s.nodes = parseStore(s.storeFile)

	return s
}

func parseStore(file *os.File) (nodes []DirectoryNode) {
	reader := bufio.NewScanner(file)

	for reader.Scan() {
		line := reader.Text()
		splitted := strings.Split(line, " ")
		weight, _ := strconv.Atoi(splitted[1])

		nodes = append(nodes, DirectoryNode{
			path:   splitted[0],
			weight: weight,
		})
	}

	slices.SortFunc(nodes, func(a, b DirectoryNode) int {
		if a.weight > b.weight {
			return -1
		} else if a.weight < b.weight {
			return 1
		}

		return 0
	})

	return
}

func (s *Store) save() {
	_ = s.storeFile.Truncate(0)
	_, _ = s.storeFile.Seek(0, 0)

	for _, node := range s.nodes {
		_, err := s.storeFile.WriteString(node.String() + "\n")
		if err != nil {
			_ = fmt.Errorf("could not write to file: %v", err)
		}
	}
}

func (s *Store) Visit(path string) error {

	for i, node := range s.nodes {
		if node.path == path {
			s.nodes[i].weight += 1
			s.save()
			return nil
		}
	}

	node := DirectoryNode{
		path:   path,
		weight: 0,
	}

	s.nodes = append(s.nodes, node)
	s.save()

	return nil
}
