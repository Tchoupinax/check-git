package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

var (
	version   string
	buildDate string
	commit    string
)

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func cliCommandDisplayVersion(args []string) {
	displayVersion := StringInSlice("-v", args[1:]) || StringInSlice("--version", args[1:])

	if displayVersion {
		bold := color.New(color.Bold).SprintFunc()
		italic := color.New(color.Italic).SprintFunc()

		fmt.Println()
		fmt.Println(bold("⚡️ Check git"))
		fmt.Println()
		fmt.Println("build date: ", bold(buildDate))
		fmt.Println("version:    ", bold(version))
		fmt.Println("commit:     ", bold(commit))
		fmt.Println()
		fmt.Println(italic("Need help?"))
		fmt.Println(italic("https://github.com/Tchoupinax/check-git/issues"))
		os.Exit(0)
	}
}
