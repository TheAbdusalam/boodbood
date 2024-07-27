package main

import (
	"fmt"
	"os"
	"strings"
)

type boodbood struct {
	actionFile string
	store      *Store
}

func NewBoodBood(actionFile string) *boodbood {
	bd := &boodbood{
		actionFile: actionFile,
		store:      NewStore(),
	}
	return bd
}

func (b *boodbood) Close() {
	_ = b.store.storeFile.Close()
}

func (b *boodbood) Navigate(args []string) {
	if len(args) > 0 && args[0] == "list" {
		for i := 0; i < 10; i++ {
			if i >= len(b.store.nodes) {
				break
			}

			fmt.Println(b.store.nodes[i].Name(), "\t\t", b.store.nodes[i].path, "=>[ weight: ", b.store.nodes[i].weight, "]")
		}
		return
	}

	switch len(args) {
	case 0:
		help()
		os.Exit(0)

	case 1:
		stat, err := os.Stat(args[0])
		if os.IsNotExist(err) {
			b.getMatch(args[0])

			return
		}

		if !stat.IsDir() {
			fmt.Println("boodbood:", args[0], "is not a directory")
			os.Exit(0)
		} else {
			b.navigateTo(args[0])
		}

		if strings.Contains(args[0], "/") {
			currentDir, _ := os.Getwd()
			b.navigateTo(currentDir + "/" + args[0])
			return
		}

		b.getMatch(args[0])
		return
	}
}

func (b *boodbood) navigateTo(identifier string) {
	runContent := "#!/bin/zsh\ncd " + identifier + "/\n"
	_ = os.WriteFile(b.actionFile, []byte(runContent), 0644)
	_ = b.store.Visit(identifier)
}

func (b *boodbood) getMatch(identifier string) {
	identifier = strings.ToLower(identifier)
	for _, node := range b.store.nodes {
		if strings.Contains(strings.ToLower(node.Name()), identifier) {
			b.navigateTo(node.path)
			return
		}
	}

	_, _ = fmt.Println("no matches found for:", identifier)
}

func help() {
	fmt.Println("Isticmaalka: \tbd <dir_name>")
	fmt.Println("\t\tbd <dir> <dir> si aad u gashid child dir")
}
