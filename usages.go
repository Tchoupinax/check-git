package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func Usage(args []string) {
	if len(args) == 0 {
		bold := color.New(color.Bold).SprintFunc()
		italic := color.New(color.Italic).SprintFunc()

		fmt.Println()
		fmt.Println(bold("Usage:"))
		fmt.Println()
		fmt.Println(bold("ckg"), italic("foldername [options]"))
		os.Exit(0)
	}
}
