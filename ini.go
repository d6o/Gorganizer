package main

import (
	"strings"

	"github.com/disiqueira/gotree"
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

	var tree gotree.GTStructure

	tree.Name = "Rules"

	for _, section := range names[1:] {

		var treeFolder gotree.GTStructure

		treeFolder.Name = section

		keys := cfg.Section(section).KeyStrings()

		for _, key := range keys {

			var treeItem gotree.GTStructure
			treeItem.Name = key

			treeFolder.Items = append(treeFolder.Items, &treeItem)
		}

		tree.Items = append(tree.Items, &treeFolder)

	}

	gotree.PrintTree(&tree)

}
