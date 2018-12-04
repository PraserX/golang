package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

var (
	source      = flag.String("s", "", "Required, Path to source directory")
	destination = flag.String("d", "", "Required, Path to destination diretory")
	pattern     = flag.String("p", "", "Required, Rename pattern")
	fileType    = flag.String("f", ".txt", "File type.")
	newLines    = flag.String("nl", "linux", "New lines: windows, linux")
)

func main() {
	var err error
	var files []os.FileInfo
	var r *regexp.Regexp

	flag.Parse()

	if *source == "" || *destination == "" || *pattern == "" {
		fmt.Fprintf(os.Stderr, "Empty string!\n")
	}

	if r, err = regexp.Compile(*pattern); err != nil {
		fmt.Fprintf(os.Stderr, "Can't compile regular expression: %s\n", err.Error())
	}

	if _, err := os.Stat(*destination); os.IsNotExist(err) {
		os.MkdirAll(*destination, os.ModePerm)
	}

	if files, err = ioutil.ReadDir(*source); err != nil {
		fmt.Fprintf(os.Stderr, "Can't read directory %s content: %s\n", *source, err.Error())
		os.Exit(1)
	}

	for _, file := range files {
		if file.IsDir() == false {
			newFileName := fmt.Sprintf("%s/%s%s", *destination, r.FindString(file.Name()), *fileType)
			fileContent, _ := ioutil.ReadFile(fmt.Sprintf("%s/%s", *source, file.Name()))
			ioutil.WriteFile(newFileName, fileContent, 0755)
		}
	}
}
