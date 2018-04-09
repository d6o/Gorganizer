package main

import (
	"strings"

	"github.com/disiqueira/gotree"
	"fmt"
)

func iniGet(key string) string {
	key = strings.ToLower(key)
	names := cfg.SectionStrings()

	for _, section := range names[1:] {
		if cfg.Section(section).HasKey(key) {
			return section
		}
	}

	return ""
}

func iniSet(key, value string) error {
	title := strings.Title(value)
	key = strings.ToLower(key)

	_, err := cfg.Section(title).NewKey(key, "")

	return err
}

func iniDelete(key string) bool {
	names := cfg.SectionStrings()

	for _, section := range names {
		if cfg.Section(section).HasKey(key) {
			cfg.Section(section).DeleteKey(key)
			return true
		}
	}

	return false
}

func iniScanExt() {
	names := cfg.SectionStrings()
	tree := gotree.New("Rules")

	for _, section := range names[1:] {
		folder := tree.Add(section)
		keys := cfg.Section(section).KeyStrings()
		for _, key := range keys {
			folder.Add(key)
		}
	}

	fmt.Println(tree.Print())
}
