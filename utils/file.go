package utils

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func Contains(s []fs.DirEntry, str string) bool {
	for _, v := range s {
		if v.Name() == str {
			return true
		}
	}

	return false
}

func LocalizeFolder(folderToSearch string, currentFolder string, level int16) string {
	if level > 2 {
		return ""
	}

	if _, err := os.Stat(folderToSearch); !os.IsNotExist(err) {
		return folderToSearch
	}

	files, _ := os.ReadDir(currentFolder)

	for _, f := range files {
		if strings.HasPrefix(f.Name(), ".") {
			continue
		}

		if f.Name() == folderToSearch {
			return currentFolder
		}

		answer := LocalizeFolder(folderToSearch, fmt.Sprintf("%s/%s", currentFolder, f.Name()), level+1)

		if answer != "" {
			return answer
		}
	}

	return ""
}
