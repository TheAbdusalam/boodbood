package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


type Store struct {
	store *os.File
	nodes []DirectoryNode
}

func NewStore() *Store {
	file, err := os.OpenFile(storeFilePath, os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		_ = fmt.Errorf("could not open file %s: %v", storeFilePath, err)
		return nil
	}

	s := &Store{
		store: file,
	}
 
	s.nodes = parseStore(s.store)

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

	return
}

func (s *Store) save() {
	_ = s.store.Truncate(0)
	_, _ = s.store.Seek(0, 0)

	for _, node := range s.nodes {
		_, _ = s.store.WriteString(node.String() + "\n")
	}
}

func (s *Store) Close() {
	_ = s.store.Close()
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
		path: path,
		weight: 0,
	}

	s.nodes = append(s.nodes, node)
	s.save()

	return nil
}
