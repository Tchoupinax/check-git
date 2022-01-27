package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"sync"

	File "ckg/utils"

	"github.com/fatih/color"
)

var wg sync.WaitGroup

func main() {
	argsWithoutProg := os.Args[1:]

	Usage(argsWithoutProg)

	dirname, _ := os.UserHomeDir()
	foundPath := File.LocalizeFolder(argsWithoutProg[0], dirname, 0)

	if foundPath == "" {
		fmt.Println("No path found")
		os.Exit(0)
	}

	path := fmt.Sprintf("%s%s%s", foundPath, "/", argsWithoutProg[0])

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

		if len(argsWithoutProg) == 1 || len(argsWithoutProg) > 1 && (argsWithoutProg[1] != "p" && argsWithoutProg[1] != "pull") {
			if len(porcelainOutput) == 0 {
				fmt.Println(green("  "), blue(f.Name()))
			} else {
				fmt.Println(red("  "), yellow(f.Name()))
			}
		}

		if len(argsWithoutProg) > 1 &&
			(argsWithoutProg[1] == "d" || argsWithoutProg[1] == "detail" || argsWithoutProg[1] == "details") &&
			len(porcelainOutput) > 0 {
			fmt.Println("")

			var str = string(porcelainOutput)
			var icon = strings.Replace(str, "\n", "\n    ✏️", -1)
			var text = strings.Replace(icon, "M", " ", -1)
			var text2 = strings.Replace(text, "?", " ", -1)

			fmt.Println("    ✏️", green(text2[1:len(text2)-8]))
		}

		if len(argsWithoutProg) > 1 &&
			(argsWithoutProg[1] == "p" || argsWithoutProg[1] == "pull") {
			if File.Contains(subfiles, ".git") {
				go Pull(fmt.Sprintf("%s%s%s", path, "/", f.Name()), f.Name(), 21)
			}
		}
	}

	wg.Wait()
	fmt.Println()
}
