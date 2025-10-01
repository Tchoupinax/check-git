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

	wg.Add(1)

	cmd := exec.Command("git", "pull")
	cmd.Dir = path

	answer, _ := cmd.Output()

	message := strings.ReplaceAll(string(answer), "\n", "")

	red := color.New(color.Bold, color.FgHiRed).SprintFunc()
	green := color.New(color.Bold, color.FgHiGreen).SprintFunc()
	blue := color.New(color.Bold, color.FgHiBlue).SprintFunc()

	if len(message) == 0 {
		message = red("Failed")
	}

	fmt.Println(blue(String.AddSpaceToEnd(name, nameSize)), green(message))
}
