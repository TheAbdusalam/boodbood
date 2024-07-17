package main

import (
	"fmt"
	"os"
)

type boodbood struct {
	actionFile string
	store      *Store
}

func NewBoodBood(actionFile string) *boodbood {
	return &boodbood{
		actionFile: actionFile,
		store:      NewStore(),
	}
}

func (b *boodbood) Navigate(args []string) {
	switch len(args) {
	case 0:
		help()
		os.Exit(0)

	case 1:
		stat, err := os.Stat(args[0])
		if os.IsNotExist(err) {
			fmt.Println("boodbood:", args[0], "does not exist")
			os.Exit(0)
		}

		if !stat.IsDir() {
			fmt.Println("boodbood:", args[0], "is not a directory")
			os.Exit(0)
		}

		runContent := "#!/bin/zsh\ncd " + args[0] + "\n"
		_ = os.WriteFile(b.actionFile, []byte(runContent), 0644)

		wd, _ := os.Getwd()
		if err := b.store.Visit(wd + "/" + args[0]); err != nil {
			_ = fmt.Errorf("Error visiting node: %v", err)
			os.Exit(1)
		}
	}
}

func help() {
	fmt.Println("Isticmaalka: \tbd <dir_name>")
	fmt.Println("\t\tbd <dir> <dir> si aad u gashid child dir")
}
