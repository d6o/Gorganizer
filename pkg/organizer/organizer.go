// Package organizer scans directories and organizes files into folders
// based on their extension using an ExtensionResolver.
package organizer

import (
	"os"
	"path/filepath"
	"strings"
)

// ExtensionResolver maps file extensions to destination folder names.
type ExtensionResolver interface {
	Lookup(ext string) string
}

// Config holds configuration for the Organizer.
type Config struct {
	InputFolder       string
	OutputFolder      string
	Preview           bool
	Recursive         bool
	IgnoreHiddenFiles bool
	ExcludeList       ExcludeList
}

// Organizer scans directories and organizes files by their extension.
type Organizer struct {
	resolver ExtensionResolver
	config   Config
}

// NewOrganizer creates an Organizer with the given resolver and config.
func NewOrganizer(resolver ExtensionResolver, config Config) *Organizer {
	return &Organizer{
		resolver: resolver,
		config:   config,
	}
}

// Run scans the input folder and organizes files according to the configured
// rules. It returns a structured result describing what happened to each file.
// In preview mode, files are categorized but not moved.
func (o *Organizer) Run() (*OrganizeResult, error) {
	var actions []FileAction

	if err := o.scan(o.config.InputFolder, &actions); err != nil {
		return nil, err
	}

	return &OrganizeResult{Actions: actions}, nil
}

func (o *Organizer) scan(inputFolder string, actions *[]FileAction) error {
	entries, err := os.ReadDir(inputFolder)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), ".") && !o.config.IgnoreHiddenFiles {
			*actions = append(*actions, FileAction{
				FileName: entry.Name(),
				Reason:   ReasonHidden,
			})
			continue
		}

		if entry.IsDir() && o.config.Recursive {
			if err := o.scan(filepath.Join(inputFolder, entry.Name()), actions); err != nil {
				return err
			}
		}

		file := filepath.Join(inputFolder, entry.Name())
		ext := strings.TrimPrefix(filepath.Ext(file), ".")

		if o.config.ExcludeList.Contains(ext) {
			*actions = append(*actions, FileAction{
				FileName: entry.Name(),
				Reason:   ReasonExcluded,
			})
			continue
		}

		folder := o.resolver.Lookup(ext)

		if folder != "" {
			dest := filepath.Join(o.config.OutputFolder, folder)
			newFile := filepath.Join(dest, entry.Name())

			moved := false
			if !o.config.Preview {
				if err := os.Mkdir(dest, os.ModePerm); err != nil && !os.IsExist(err) {
					return err
				}
				if err := os.Rename(file, newFile); err != nil {
					return err
				}
				moved = true
			}

			*actions = append(*actions, FileAction{
				FileName:    entry.Name(),
				Destination: folder,
				Reason:      ReasonOrganized,
				Moved:       moved,
			})
		} else {
			*actions = append(*actions, FileAction{
				FileName: entry.Name(),
				Reason:   ReasonUnknownExtension,
			})
		}
	}

	return nil
}
