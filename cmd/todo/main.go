package main

import (
	"fmt"
	"go_cli/interacting/todo"
	"os"
	"strings"
)

const todoFileName = ".todo.json"

func main() {
	l := &todo.List{}

	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		// User STDERR instead of STDOUT to display errors
		os.Exit(1)
	}

	switch {
	case len(os.Args) == 1:
		for _, item := range *l {
			fmt.Println(item.Task)
		}
	default:
		item := strings.Join(os.Args[1:], " ")
	}

	l.Add(item)

	if err := l.Save(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
