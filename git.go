package main

import (
	"fmt"
	"os/exec"
	"strings"

	String "ckg/utils"

	"github.com/fatih/color"
)

func Pull(path string, name string, nameSize int) {
	defer wg.Done()

	cmd := exec.Command("git", "pull")
	cmd.Dir = path

	answer, err := cmd.CombinedOutput()
	output := strings.TrimSpace(string(answer))

	red := color.New(color.Bold, color.FgHiRed).SprintFunc()
	green := color.New(color.Bold, color.FgHiGreen).SprintFunc()
	blue := color.New(color.Bold, color.FgHiBlue).SprintFunc()

	var message string
	if err != nil {
		if output != "" {
			message = red(output)
		} else {
			message = red(err.Error())
		}
	} else {
		if output == "" {
			message = green("Already up to date.")
		} else {
			message = green(output)
		}
	}

	fmt.Println(blue(String.AddSpaceToEnd(name, nameSize)), message)
}
