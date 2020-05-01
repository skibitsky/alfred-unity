package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

type project struct {
	name          string
	path          string
	editorVersion string
}

const proj = "defaults read com.unity3d.UnityEditor5.x | grep RecentlyUsedProjectPaths"

func getProjects() []project {
	var prs []project
	cmd := exec.Command("sh", "-c", proj)
	out, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	} else {
		outS := strings.Split(string(out), "\n")
		var re = regexp.MustCompile(`\"(.*?)\"`)
		for _, v := range outS {
			matches := re.FindAllStringSubmatch(v, -1)
			if len(matches) < 2 {
				continue
			}

			prPath := matches[1][1]
			prName := filepath.Base(prPath)
			edVersion := getEditorVersion(prPath)
			fmt.Println(project{name: prName, path: prPath, editorVersion: edVersion})
			prs = append(prs, project{name: prName, path: prPath, editorVersion: edVersion})
		}
	}

	return prs
}

func getEditorVersion(projectPath string) string {
	file, err := os.Open(projectPath + "/ProjectSettings/ProjectVersion.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	line, _, err := reader.ReadLine()
	if err != nil {
		log.Fatal(err)
	}

	version := strings.Fields(string(line))[1]

	return version
}
