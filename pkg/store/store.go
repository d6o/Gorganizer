// Package store manages file extension to folder mapping rules backed by an INI config file.
package store

import (
	"os/user"
	"path/filepath"
	"strings"
	"unicode"

	"gopkg.in/ini.v1"
)

const configFile = ".gorganizer-{lang}.ini"

// StoreEvent represents a lifecycle event emitted during Store operations.
type StoreEvent int

const (
	// StoreEventDatabaseNotFound is emitted when no existing config file is found.
	StoreEventDatabaseNotFound StoreEvent = iota
	// StoreEventCreatingDefaults is emitted when default rules are being created.
	StoreEventCreatingDefaults
	// StoreEventDefaultsInitialized is emitted after default rules have been written.
	StoreEventDefaultsInitialized
)

// StoreOption configures a Store during construction.
type StoreOption func(*Store)

// WithEventHandler sets a callback for store lifecycle events.
// The callback is invoked synchronously during NewStore when creating defaults.
func WithEventHandler(fn func(StoreEvent)) StoreOption {
	return func(s *Store) {
		s.onEvent = fn
	}
}

// WithConfigDir overrides the directory where the config file is stored.
// When set, the store only looks in this directory instead of the current
// directory and user home directory.
func WithConfigDir(dir string) StoreOption {
	return func(s *Store) {
		s.configDir = dir
	}
}

// Store manages file extension to folder mapping rules backed by an INI config file.
type Store struct {
	cfg       *ini.File
	cfgFile   string
	onEvent   func(StoreEvent)
	configDir string
}

// NewStore creates a Store for the given language. It searches for a config file
// in the current directory and the user's home directory. If none is found, it
// creates a new config with default rules for 60+ file types.
func NewStore(lang string, opts ...StoreOption) (*Store, error) {
	if lang == "" {
		lang = "en"
	}

	s := &Store{}
	for _, opt := range opts {
		opt(s)
	}

	cfgFileName := strings.Replace(configFile, "{lang}", lang, 1)

	if s.configDir != "" {
		file := filepath.Join(s.configDir, cfgFileName)
		if s.tryLoad(file) {
			return s, nil
		}
		s.cfg = ini.Empty()
		s.cfgFile = file
	} else {
		if s.tryLoad(cfgFileName) {
			return s, nil
		}

		currentUser, err := user.Current()
		if err != nil {
			return nil, err
		}

		homeFile := filepath.Join(currentUser.HomeDir, cfgFileName)
		if s.tryLoad(homeFile) {
			return s, nil
		}

		s.cfg = ini.Empty()
		s.cfgFile = homeFile
	}

	if err := s.populateDefaults(lang); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Store) tryLoad(file string) bool {
	cfg, err := ini.LoadSources(ini.LoadOptions{AllowBooleanKeys: true, Loose: false}, file)
	if err != nil {
		return false
	}
	s.cfg = cfg
	s.cfgFile = file
	return true
}

func (s *Store) emitEvent(evt StoreEvent) {
	if s.onEvent != nil {
		s.onEvent(evt)
	}
}

// Close saves the current rules to the config file on disk.
func (s *Store) Close() error {
	return s.cfg.SaveTo(s.cfgFile)
}

// Lookup returns the destination folder name for the given file extension,
// or an empty string if no rule matches. The lookup is case-insensitive.
func (s *Store) Lookup(ext string) string {
	ext = strings.ToLower(ext)
	sections := s.cfg.SectionStrings()

	for _, section := range sections[1:] {
		if s.cfg.Section(section).HasKey(ext) {
			return section
		}
	}

	return ""
}

// InsertRule adds a new extension-to-folder mapping. The rule must be in
// "ext:folder" format (e.g., "mp3:Music"). Returns ErrInvalidRuleFormat
// or ErrEmptyRuleComponent on invalid input.
func (s *Store) InsertRule(rule string) error {
	parts := strings.Split(rule, ":")
	if len(parts) != 2 {
		return ErrInvalidRuleFormat
	}
	if parts[0] == "" || parts[1] == "" {
		return ErrEmptyRuleComponent
	}

	return s.set(parts[0], parts[1])
}

// DeleteRule removes the rule for the given file extension. If no rule
// exists for the extension, it is a no-op.
func (s *Store) DeleteRule(ext string) {
	sections := s.cfg.SectionStrings()

	for _, section := range sections {
		if s.cfg.Section(section).HasKey(ext) {
			s.cfg.Section(section).DeleteKey(ext)
			return
		}
	}
}

// Rules returns all configured rules ordered by folder, then by extension.
func (s *Store) Rules() []Rule {
	var rules []Rule
	sections := s.cfg.SectionStrings()

	for _, section := range sections[1:] {
		keys := s.cfg.Section(section).KeyStrings()
		for _, key := range keys {
			rules = append(rules, Rule{Extension: key, Folder: section})
		}
	}

	return rules
}

func (s *Store) set(key, value string) error {
	prev := rune(' ')
	runes := []rune(value)
	for i, r := range runes {
		if s.isTitleSeparator(prev) {
			runes[i] = unicode.ToTitle(r)
		}
		prev = r
	}

	key = strings.ToLower(key)
	_, err := s.cfg.Section(string(runes)).NewKey(key, "")
	return err
}

func (s *Store) isTitleSeparator(r rune) bool {
	if r <= 0x7F {
		switch {
		case '0' <= r && r <= '9':
			return false
		case 'a' <= r && r <= 'z':
			return false
		case 'A' <= r && r <= 'Z':
			return false
		case r == '_':
			return false
		}
		return true
	}
	if unicode.IsLetter(r) || unicode.IsDigit(r) {
		return false
	}
	return unicode.IsSpace(r)
}
