package main

import (
	Utils "ckg/utils"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/fatih/color"
)

// nolint go-golangci-lint
var wg sync.WaitGroup

// nolint go-golangci-lint
func main() {
	// Check if the version is asked by flag
	cliCommandDisplayVersion(os.Args)

	commandLinesArgumentsWithoutProgram := os.Args[1:]
	Usage(commandLinesArgumentsWithoutProgram)

	dirname, _ := os.UserHomeDir()
	foundPath := Utils.LocalizeFolder(commandLinesArgumentsWithoutProgram[0], dirname, 0)
	fmt.Println(foundPath)

	if foundPath == "" {
		fmt.Println("No path found")
		os.Exit(0)
	}

	var path string
	if (foundPath == commandLinesArgumentsWithoutProgram[0]) {
		path = foundPath
	} else {
		path = fmt.Sprintf("%s%s%s", foundPath, "/", commandLinesArgumentsWithoutProgram[0])
	}

	title := color.New(color.Bold, color.FgHiMagenta).SprintFunc()
	green := color.New(color.Bold, color.FgHiGreen).SprintFunc()
	red := color.New(color.Bold, color.FgHiRed).SprintFunc()
	blue := color.New(color.Bold, color.FgHiBlue).SprintFunc()
	yellow := color.New(color.Bold, color.FgHiYellow).SprintFunc()

	fmt.Println()
	fmt.Printf("%s%s", "  🔎 ", title(path))
	fmt.Println()
	fmt.Println()

	children, _ := os.ReadDir(path)

	for _, f := range children {
		subfiles, err := os.ReadDir(fmt.Sprintf("%s%s%s", path, "/", f.Name()))

		if err != nil {
			continue
		}

		if !Utils.Contains(subfiles, ".git") {
			continue
		}

		cmd := exec.Command("git", "status", "--porcelain")
		cmd.Dir = fmt.Sprintf("%s%s%s", path, "/", f.Name())
		porcelainOutput, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(red("  \uf00d"), yellow(f.Name()), red(string(porcelainOutput)))
			continue
		}

		if len(commandLinesArgumentsWithoutProgram) == 1 ||
			len(commandLinesArgumentsWithoutProgram) > 1 &&
				(commandLinesArgumentsWithoutProgram[1] != "p" && commandLinesArgumentsWithoutProgram[1] != "pull") {
			if len(porcelainOutput) == 0 {
				fmt.Println(green("  "), blue(f.Name()))
			} else {
				fmt.Println(red("  "), yellow(f.Name()))
			}
		}

		if len(commandLinesArgumentsWithoutProgram) > 1 &&
			Utils.ContainsOneOfThese(commandLinesArgumentsWithoutProgram[1], []string{"d", "detail", "details"}) &&
			len(porcelainOutput) > 0 {
			for _, line := range strings.Split(string(porcelainOutput), "\n") {
				line = strings.TrimSpace(line)
				if line == "" {
					continue
				}
				fmt.Printf("    ✏️  %s\n", green(line))
			}
		}

		if len(commandLinesArgumentsWithoutProgram) > 1 && Utils.ContainsOneOfThese(commandLinesArgumentsWithoutProgram[1], []string{"p", "pull"}) {
			if Utils.Contains(subfiles, ".git") {
				wg.Add(1)
				go Pull(fmt.Sprintf("%s%s%s", path, "/", f.Name()), f.Name(), 30)
			}
		}
	}

	wg.Wait()
	fmt.Println()
}
