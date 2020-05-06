package main

import (
	"io/ioutil"
	"log"
	"regexp"
)

const editorsPath string = "/Applications/Unity/Hub/Editor/"

type install struct {
	version string
	path    string
}

func getInstalls() []install {
	versions, err := ioutil.ReadDir(editorsPath)
	if err != nil {
		log.Fatal(err)
	}

	var i []install
	var re = regexp.MustCompile(`\d+(?:\.\d+)+\w\d$`)
	for _, v := range versions {
		matches := re.FindAllString(v.Name(), -1)
		if len(matches) == 0 {
			continue
		}

		ni := install{
			version: matches[0],
			path:    editorsPath + matches[0],
		}

		i = append(i, ni)
	}
	return i
}
