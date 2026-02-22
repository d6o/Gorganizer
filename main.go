package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/disiqueira/gotree"

	"github.com/d6o/Gorganizer/pkg/organizer"
	"github.com/d6o/Gorganizer/pkg/store"
)

var version = "dev"

func main() {
	showVersion := flag.Bool("version", false, "Print version and exit")
	outputFolder := flag.String("output", ".", "Main directory to put organized folders")
	inputFolder := flag.String("directory", ".", "The directory whose files to classify")
	newRule := flag.String("newrule", "", "Insert a new rule. Format ext:folder Example: mp3:Music")
	delRule := flag.String("delrule", "", "Delete a rule. Format ext Example: mp3")
	printRules := flag.Bool("allrules", false, "Print all rules")
	preview := flag.Bool("preview", false, "Only preview, do not move files")
	recursive := flag.Bool("recursive", false, "Search over all directories.")
	ignoreHiddenFiles := flag.Bool("hidden", true, "Ignore hidden files")
	excludeExtensions := flag.String("exclude", "", "Exclude files will ignore files for organizer. Format pdf,odt")
	lang := flag.String("language", "en", "Specify language: en|tr|pt")

	flag.Parse()

	if *showVersion {
		fmt.Println(version)
		return
	}

	s, err := store.NewStore(*lang, store.WithEventHandler(func(evt store.StoreEvent) {
		switch evt {
		case store.StoreEventDatabaseNotFound:
			fmt.Println("No database found")
		case store.StoreEventCreatingDefaults:
			fmt.Println("Creating default database")
		case store.StoreEventDefaultsInitialized:
			fmt.Println("Default database initialized")
		}
	}))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func() {
		if err := s.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	if *newRule != "" {
		fmt.Println("Creating new rule")
		if err := s.InsertRule(*newRule); err != nil {
			fmt.Println(err)
			return
		}
		printRulesTree(s)
		return
	}

	if *delRule != "" {
		fmt.Println("Deleting rule")
		s.DeleteRule(*delRule)
		printRulesTree(s)
		return
	}

	if *printRules {
		printRulesTree(s)
		return
	}

	org := organizer.NewOrganizer(s, organizer.OrganizerConfig{
		InputFolder:       *inputFolder,
		OutputFolder:      *outputFolder,
		Preview:           *preview,
		Recursive:         *recursive,
		IgnoreHiddenFiles: *ignoreHiddenFiles,
		ExcludeList:       organizer.ExcludeList(strings.Split(*excludeExtensions, ",")),
	})

	fmt.Println("GOrganizing your Files")

	result, err := org.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	printResultTree(result)

	fmt.Println("All files have been GOrganized!")
}

func printRulesTree(s *store.Store) {
	rules := s.Rules()
	tree := gotree.New("Rules")
	folders := make(map[string]gotree.Tree)

	for _, r := range rules {
		ft, ok := folders[r.Folder]
		if !ok {
			ft = tree.Add(r.Folder)
			folders[r.Folder] = ft
		}
		ft.Add(r.Extension)
	}

	fmt.Println(tree.Print())
}

func printResultTree(result *organizer.OrganizeResult) {
	tree := gotree.New("Files")

	for _, a := range result.Actions {
		var label string
		switch a.Reason {
		case organizer.ReasonHidden:
			label = "Hidden Files"
		case organizer.ReasonExcluded:
			label = "Excluded Files"
		case organizer.ReasonUnknownExtension:
			label = "Unknown extension (will not be moved)"
		case organizer.ReasonOrganized:
			label = a.Destination
		}
		addToTree(tree, label, a.FileName)
	}

	fmt.Println(tree.Print())
}

func addToTree(tree gotree.Tree, folder, file string) {
	for _, item := range tree.Items() {
		if item.Text() == folder {
			item.Add(file)
			return
		}
	}
	tree.Add(folder).Add(file)
}
