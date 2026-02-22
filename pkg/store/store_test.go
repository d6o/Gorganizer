package store_test

import (
	"testing"

	"github.com/d6o/Gorganizer/pkg/store"
)

func newTestStore(t *testing.T, lang string, opts ...store.StoreOption) *store.Store {
	t.Helper()
	dir := t.TempDir()
	allOpts := append([]store.StoreOption{store.WithConfigDir(dir)}, opts...)
	s, err := store.NewStore(lang, allOpts...)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := s.Close(); err != nil {
			t.Fatal(err)
		}
	})
	return s
}

func TestNewStore(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		lang string
	}{
		{"english", "en"},
		{"portuguese", "pt"},
		{"turkish", "tr"},
		{"empty defaults to english", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := newTestStore(t, tt.lang)
			rules := s.Rules()
			if len(rules) == 0 {
				t.Error("expected default rules to be populated")
			}
		})
	}
}

func TestNewStoreEventHandler(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()

	var events []store.StoreEvent
	s, err := store.NewStore("en",
		store.WithConfigDir(dir),
		store.WithEventHandler(func(evt store.StoreEvent) {
			events = append(events, evt)
		}),
	)
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()

	expected := []store.StoreEvent{
		store.StoreEventDatabaseNotFound,
		store.StoreEventCreatingDefaults,
		store.StoreEventDefaultsInitialized,
	}

	if len(events) != len(expected) {
		t.Fatalf("got %d events, want %d", len(events), len(expected))
	}
	for i, evt := range events {
		if evt != expected[i] {
			t.Errorf("event[%d] = %d, want %d", i, evt, expected[i])
		}
	}
}

func TestNewStoreLoadsExisting(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()

	s1, err := store.NewStore("en", store.WithConfigDir(dir))
	if err != nil {
		t.Fatal(err)
	}
	if err := s1.InsertRule("xyz:TestFolder"); err != nil {
		t.Fatal(err)
	}
	if err := s1.Close(); err != nil {
		t.Fatal(err)
	}

	var events []store.StoreEvent
	s2, err := store.NewStore("en",
		store.WithConfigDir(dir),
		store.WithEventHandler(func(evt store.StoreEvent) {
			events = append(events, evt)
		}),
	)
	if err != nil {
		t.Fatal(err)
	}
	defer s2.Close()

	if len(events) != 0 {
		t.Errorf("expected no events when loading existing db, got %d", len(events))
	}

	if folder := s2.Lookup("xyz"); folder != "TestFolder" {
		t.Errorf("Lookup(xyz) = %q, want %q", folder, "TestFolder")
	}
}

func TestInsertRule(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		rule    string
		wantErr error
	}{
		{"valid rule", "py:Python", nil},
		{"missing colon", "pyPython", store.ErrInvalidRuleFormat},
		{"too many colons", "py:Py:thon", store.ErrInvalidRuleFormat},
		{"empty extension", ":Python", store.ErrEmptyRuleComponent},
		{"empty folder", "py:", store.ErrEmptyRuleComponent},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := newTestStore(t, "en")

			err := s.InsertRule(tt.rule)
			if err != tt.wantErr {
				t.Errorf("InsertRule(%q) error = %v, want %v", tt.rule, err, tt.wantErr)
			}
		})
	}
}

func TestLookup(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		ext  string
		want string
	}{
		{"known extension mp3", "mp3", "Music"},
		{"known extension pdf", "pdf", "Documents"},
		{"known extension png", "png", "Pictures"},
		{"known extension zip", "zip", "Archives"},
		{"case insensitive", "MP3", "Music"},
		{"unknown extension", "xyz123", ""},
		{"empty extension", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := newTestStore(t, "en")

			got := s.Lookup(tt.ext)
			if got != tt.want {
				t.Errorf("Lookup(%q) = %q, want %q", tt.ext, got, tt.want)
			}
		})
	}
}

func TestDeleteRule(t *testing.T) {
	t.Parallel()
	s := newTestStore(t, "en")

	if folder := s.Lookup("mp3"); folder == "" {
		t.Fatal("mp3 should exist before deletion")
	}

	s.DeleteRule("mp3")

	if folder := s.Lookup("mp3"); folder != "" {
		t.Errorf("Lookup(mp3) after delete = %q, want empty", folder)
	}
}

func TestRules(t *testing.T) {
	t.Parallel()
	s := newTestStore(t, "en")

	rules := s.Rules()
	if len(rules) == 0 {
		t.Fatal("expected non-empty rules")
	}

	found := false
	for _, r := range rules {
		if r.Extension == "mp3" && r.Folder == "Music" {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected to find mp3:Music rule")
	}
}

func TestRulesAfterInsert(t *testing.T) {
	t.Parallel()
	s := newTestStore(t, "en")

	before := len(s.Rules())
	if err := s.InsertRule("xyz:CustomFolder"); err != nil {
		t.Fatal(err)
	}
	after := len(s.Rules())

	if after != before+1 {
		t.Errorf("rules count after insert = %d, want %d", after, before+1)
	}
}

func TestPortugueseTranslations(t *testing.T) {
	t.Parallel()
	s := newTestStore(t, "pt")

	if folder := s.Lookup("mp3"); folder != "Musicas" {
		t.Errorf("Lookup(mp3) with pt = %q, want %q", folder, "Musicas")
	}
	if folder := s.Lookup("pdf"); folder != "Documentos" {
		t.Errorf("Lookup(pdf) with pt = %q, want %q", folder, "Documentos")
	}
}

func TestTurkishTranslations(t *testing.T) {
	t.Parallel()
	s := newTestStore(t, "tr")

	if folder := s.Lookup("mp3"); folder != "Müzikler" {
		t.Errorf("Lookup(mp3) with tr = %q, want %q", folder, "Müzikler")
	}
	if folder := s.Lookup("pdf"); folder != "Dokümanlar" {
		t.Errorf("Lookup(pdf) with tr = %q, want %q", folder, "Dokümanlar")
	}
}
