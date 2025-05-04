package cmd

import (
	"flag"
	"fmt"
	"os"
)

func Execute() {
	// глобальные флаги, если понадобятся
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("usage: godo <command> [options]")
		return
	}

	switch os.Args[1] {
	case "add":
		// временный вывод-заглушка
		fmt.Println("TODO: implement add")
	default:
		fmt.Println("unknown command:", os.Args[1])
	}
}
