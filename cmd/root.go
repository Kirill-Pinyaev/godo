package cmd

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func Execute() {
	// глобальные флаги, если понадобятся
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("usage: godo <command> [options]")
		return
	}

	home := os.Getenv("HOME")
	filePath := filepath.Join(home, ".godo", "tasks.json")

	switch args[0] {
	case "add":
		add(args[1:], filePath)
	case "list":
		list(args[1:], filePath)
	case "done":
		done(args[1:], filePath)
	default:
		fmt.Println("unknown command:", os.Args[1])
		os.Exit(1)
	}
}
