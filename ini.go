package main

import (
	"strings"
	"fmt"
)


func iniGet(key string) string {

	names := cfg.SectionStrings()

	for _, section := range names[1:] {

		if cfg.Section(section).HasKey(key) {
			return  section
		}
	}

	return ""
}

func iniSet(key, value string) error {

	title := strings.Title(value)

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
	fmt.Println("Extension	|	Folder")

	for _, section := range names[1:] {
		keys := cfg.Section(section).KeyStrings()

		for _, key := range keys {
			fmt.Printf("%s		|	%s\n", key, section)
		}

	}

}
