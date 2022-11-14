package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"sync"

	File "ckg/utils"
	String "ckg/utils"

	"github.com/fatih/color"
)

var wg sync.WaitGroup

func main() {
	// Check if the version is asked by flag
	cliCommandDisplayVersion(os.Args)

	commandLinesArgumentsWithoutProgram := os.Args[1:]
	Usage(commandLinesArgumentsWithoutProgram)

	dirname, _ := os.UserHomeDir()
	foundPath := File.LocalizeFolder(commandLinesArgumentsWithoutProgram[0], dirname, 0)

	if foundPath == "" {
		fmt.Println("No path found")
		os.Exit(0)
	}

	path := fmt.Sprintf("%s%s%s", foundPath, "/", commandLinesArgumentsWithoutProgram[0])

	title := color.New(color.Bold, color.FgHiMagenta).SprintFunc()
	green := color.New(color.Bold, color.FgHiGreen).SprintFunc()
	red := color.New(color.Bold, color.FgHiRed).SprintFunc()
	blue := color.New(color.Bold, color.FgHiBlue).SprintFunc()
	yellow := color.New(color.Bold, color.FgHiYellow).SprintFunc()

	fmt.Println()
	fmt.Printf("%s%s", "  🔎 ", title(path))
	fmt.Println()
	fmt.Println()

	children, _ := ioutil.ReadDir(path)

	for _, f := range children {
		subfiles, err := ioutil.ReadDir(fmt.Sprintf("%s%s%s", path, "/", f.Name()))

		if err != nil {
			continue
		}

		if !File.Contains(subfiles, ".git") {
			continue
		}

		cmd := exec.Command("git", "status", "--porcelain")
		cmd.Dir = fmt.Sprintf("%s%s%s", path, "/", f.Name())
		porcelainOutput, _ := cmd.Output()

		if len(commandLinesArgumentsWithoutProgram) == 1 || len(commandLinesArgumentsWithoutProgram) > 1 && (commandLinesArgumentsWithoutProgram[1] != "p" && commandLinesArgumentsWithoutProgram[1] != "pull") {
			if len(porcelainOutput) == 0 {
				fmt.Println(green("  "), blue(f.Name()))
			} else {
				fmt.Println(red("  "), yellow(f.Name()))
			}
		}

		if len(commandLinesArgumentsWithoutProgram) > 1 &&
			String.ContainsOneOfThese(commandLinesArgumentsWithoutProgram[1], []string{"d", "detail", "details"}) &&
			len(porcelainOutput) > 0 {
			var str = string(porcelainOutput)
			var icon = strings.Replace(str, "\n", "\n    ✏️", -1)
			var text = strings.Replace(icon, " M", "  ", -1)
			var text2 = strings.Replace(text, "?", " ", -1)

			fmt.Println("    ✏️", green(text2[1:len(text2)-8]))
		}

		if len(commandLinesArgumentsWithoutProgram) > 1 && String.ContainsOneOfThese(commandLinesArgumentsWithoutProgram[1], []string{"p", "pull"}) {
			if File.Contains(subfiles, ".git") {
				go Pull(fmt.Sprintf("%s%s%s", path, "/", f.Name()), f.Name(), 21)
			}
		}
	}

	wg.Wait()
	fmt.Println()
}
