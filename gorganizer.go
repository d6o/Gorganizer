package main

import (
	"flag"
	"fmt"
	"github.com/boltdb/bolt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

const (
	dbName        = "gorganizer.db"
	dbPermissions = 0600
	bucketName    = "Gorganizer"
)

var db, err = bolt.Open(dbName, dbPermissions, nil)

func main() {

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	outputFolder := flag.String("output", "./", "Main directory to put organized folders")
	inputFolder := flag.String("directory", "./", "The directory whose files to classify")

	newRule := flag.String("newrule", "", "Insert a new rule. Format ext:folder Example: mp3:Music")
	delRule := flag.String("delrule", "", "Delete a rule. Format ext Example: mp3")

	printRules := flag.Bool("allrules", false, "Print all rules")

	preview := flag.Bool("preview", false, "Only preview, do not move files")

	flag.Parse()

	initDb()

	if len(*newRule) > 0 {
		insertRule(*newRule)
		showRules()
		return
	}

	if len(*delRule) > 0 {
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

		file := *inputFolder + "/" + f.Name()
		ext := strings.TrimPrefix(path.Ext(file), ".")
		folder := *outputFolder + "/" + boltGet("ext:"+ext)
		newFile := folder + "/" + f.Name()

		fmt.Println(file + " --> " + newFile)

		if !*preview {
			_ = os.Mkdir(folder, os.ModePerm)
			os.Rename(file, newFile)
		}
	}
	fmt.Println("All files have been gorganized!")
}
