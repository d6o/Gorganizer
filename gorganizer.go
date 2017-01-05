package main

import (
	"flag"
	"fmt"
	"gopkg.in/ini.v1"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const (
	configFile = ".gorganizer-{lang}.ini"
)

var cfg *ini.File
var cfgFile string
var language string

func main() {
	outputFolder := flag.String("output", ".", "Main directory to put organized folders")
	inputFolder := flag.String("directory", ".", "The directory whose files to classify")

	newRule := flag.String("newrule", "", "Insert a new rule. Format ext:folder Example: mp3:Music")
	delRule := flag.String("delrule", "", "Delete a rule. Format ext Example: mp3")

	printRules := flag.Bool("allrules", false, "Print all rules")

	preview := flag.Bool("preview", false, "Only preview, do not move files")

	flag.StringVar(&language, "language", "en", "Specify language: en|tr|pt")

	flag.Parse()

	initDb()

	defer closeDb()

	if len(*newRule) > 0 {
		fmt.Println("Creating new rule")
		err := insertRule(*newRule)
		if err != nil {
			fmt.Println(err)
			return
		}
		showRules()
		return
	}

	if len(*delRule) > 0 {
		fmt.Println("Deleting rule")
		deleteRule(*delRule)
		showRules()
		return
	}

	if *printRules {
		showRules()
		return
	}

	files, _ := ioutil.ReadDir(*inputFolder)
	fmt.Println("GOrganizing your Files")

	for _, f := range files {
		if f.IsDir() {
			continue
		}

		file := filepath.Join(*inputFolder, f.Name())
		ext := strings.TrimPrefix(path.Ext(file), ".")

		newFolder := iniGet(ext)

		if len(newFolder) > 0 {

			folder := filepath.Join(*outputFolder, iniGet(ext))
			newFile := filepath.Join(folder, f.Name())

			fmt.Println(file + " --> " + newFile)

			if !*preview {
				_ = os.Mkdir(folder, os.ModePerm)
				os.Rename(file, newFile)
			}
		}
	}

	fmt.Println("All files have been gorganized!")

}
