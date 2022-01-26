package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

func contains(s []fs.FileInfo, str string) bool {
	for _, v := range s {
		if v.Name() == str {
			return true
		}
	}

	return false
}

func searchFolder(folderToSearch string, currentFolder string, level int16) string {
	if level > 2 {
		return ""
	}

	files, _ := ioutil.ReadDir(currentFolder)

	for _, f := range files {
		if strings.HasPrefix(f.Name(), ".") {
			continue
		}

		if f.Name() == folderToSearch {
			return currentFolder
		}

		answer := searchFolder(folderToSearch, fmt.Sprintf("%s/%s", currentFolder, f.Name()), level+1)

		if answer != "" {
			return answer
		}
	}

	return ""
}

func usage(args []string) {
	if len(args) == 0 {
		bold := color.New(color.Bold).SprintFunc()
		italic := color.New(color.Italic).SprintFunc()

		fmt.Println(bold("Usage:"))
		fmt.Println()
		fmt.Println(bold("ckg"), italic("foldername [options]"))
		os.Exit(0)
	}
}

func main() {
	argsWithoutProg := os.Args[1:]

	usage(argsWithoutProg)

	dirname, _ := os.UserHomeDir()
	foundPath := searchFolder(argsWithoutProg[0], dirname, 0)

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

		if !contains(subfiles, ".git") {
			continue
		}

		cmd := exec.Command("git", "status", "--porcelain")
		cmd.Dir = fmt.Sprintf("%s%s%s", path, "/", f.Name())
		porcelainOutput, _ := cmd.Output()

		if len(porcelainOutput) == 0 {
			fmt.Println(green("  "), blue(f.Name()))
		} else {
			fmt.Println(red("  "), yellow(f.Name()))
		}

		if len(argsWithoutProg) > 1 &&
			(argsWithoutProg[1] == "d" || argsWithoutProg[1] == "detail" || argsWithoutProg[1] == "details") &&
			len(porcelainOutput) > 0 {
			fmt.Println("")

			var str = string(porcelainOutput)
			var icon = strings.Replace(str, "\n", "\n    ✏️", -1)
			var text = strings.Replace(icon, "M", "", -1)

			fmt.Println("    ✏️", green(text[1:len(text)-8]))
		}

		if len(argsWithoutProg) > 1 &&
			(argsWithoutProg[1] == "p" || argsWithoutProg[1] == "pull") {
			if contains(subfiles, ".git") {
				cmd := exec.Command("git", "pull")
				cmd.Dir = fmt.Sprintf("%s%s%s", path, "/", f.Name())
				stdout, _ := cmd.Output()

				fmt.Println(string(stdout))
			}
		}
	}

	fmt.Println()
}
