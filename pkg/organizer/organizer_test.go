package organizer_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/d6o/Gorganizer/pkg/organizer"
)

type mockResolver struct {
	rules map[string]string
}

func (m *mockResolver) Lookup(ext string) string {
	return m.rules[ext]
}

func newMockResolver() *mockResolver {
	return &mockResolver{
		rules: map[string]string{
			"mp3": "Music",
			"pdf": "Documents",
			"jpg": "Pictures",
			"zip": "Archives",
		},
	}
}

func createTestFile(t *testing.T, dir, name string) {
	t.Helper()
	f, err := os.Create(filepath.Join(dir, name))
	if err != nil {
		t.Fatal(err)
	}
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestOrganizer_Run_EmptyDirectory(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()

	org := organizer.NewOrganizer(newMockResolver(), organizer.OrganizerConfig{
		InputFolder:       dir,
		OutputFolder:      dir,
		Preview:           true,
		IgnoreHiddenFiles: true,
	})

	result, err := org.Run()
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Actions) != 0 {
		t.Errorf("expected 0 actions, got %d", len(result.Actions))
	}
}

func TestOrganizer_Run_CategorizesFiles(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()

	createTestFile(t, dir, "song.mp3")
	createTestFile(t, dir, "doc.pdf")
	createTestFile(t, dir, "photo.jpg")

	org := organizer.NewOrganizer(newMockResolver(), organizer.OrganizerConfig{
		InputFolder:       dir,
		OutputFolder:      dir,
		Preview:           true,
		IgnoreHiddenFiles: true,
	})

	result, err := org.Run()
	if err != nil {
		t.Fatal(err)
	}

	expected := map[string]organizer.ActionReason{
		"song.mp3":  organizer.ReasonOrganized,
		"doc.pdf":   organizer.ReasonOrganized,
		"photo.jpg": organizer.ReasonOrganized,
	}

	for _, a := range result.Actions {
		reason, ok := expected[a.FileName]
		if !ok {
			continue
		}
		if a.Reason != reason {
			t.Errorf("file %q: reason = %d, want %d", a.FileName, a.Reason, reason)
		}
	}
}

func TestOrganizer_Run_PreviewDoesNotMove(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()

	createTestFile(t, dir, "song.mp3")

	org := organizer.NewOrganizer(newMockResolver(), organizer.OrganizerConfig{
		InputFolder:       dir,
		OutputFolder:      dir,
		Preview:           true,
		IgnoreHiddenFiles: true,
	})

	result, err := org.Run()
	if err != nil {
		t.Fatal(err)
	}

	for _, a := range result.Actions {
		if a.Moved {
			t.Errorf("file %q should not be moved in preview mode", a.FileName)
		}
	}

	if _, err := os.Stat(filepath.Join(dir, "song.mp3")); os.IsNotExist(err) {
		t.Error("file should still exist in original location during preview")
	}
}

func TestOrganizer_Run_MovesFiles(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	out := t.TempDir()

	createTestFile(t, dir, "song.mp3")

	org := organizer.NewOrganizer(newMockResolver(), organizer.OrganizerConfig{
		InputFolder:       dir,
		OutputFolder:      out,
		Preview:           false,
		IgnoreHiddenFiles: true,
	})

	result, err := org.Run()
	if err != nil {
		t.Fatal(err)
	}

	var found bool
	for _, a := range result.Actions {
		if a.FileName == "song.mp3" {
			found = true
			if !a.Moved {
				t.Error("expected Moved=true for song.mp3")
			}
			if a.Destination != "Music" {
				t.Errorf("Destination = %q, want %q", a.Destination, "Music")
			}
		}
	}
	if !found {
		t.Fatal("song.mp3 not found in actions")
	}

	if _, err := os.Stat(filepath.Join(out, "Music", "song.mp3")); os.IsNotExist(err) {
		t.Error("file should exist in destination folder")
	}
}

func TestOrganizer_Run_UnknownExtension(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()

	createTestFile(t, dir, "data.xyz")

	org := organizer.NewOrganizer(newMockResolver(), organizer.OrganizerConfig{
		InputFolder:       dir,
		OutputFolder:      dir,
		Preview:           true,
		IgnoreHiddenFiles: true,
	})

	result, err := org.Run()
	if err != nil {
		t.Fatal(err)
	}

	for _, a := range result.Actions {
		if a.FileName == "data.xyz" {
			if a.Reason != organizer.ReasonUnknownExtension {
				t.Errorf("reason = %d, want ReasonUnknownExtension", a.Reason)
			}
			return
		}
	}
	t.Error("data.xyz not found in actions")
}

func TestOrganizer_Run_ExcludeList(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()

	createTestFile(t, dir, "song.mp3")
	createTestFile(t, dir, "doc.pdf")

	org := organizer.NewOrganizer(newMockResolver(), organizer.OrganizerConfig{
		InputFolder:       dir,
		OutputFolder:      dir,
		Preview:           true,
		IgnoreHiddenFiles: true,
		ExcludeList:       organizer.ExcludeList{"mp3"},
	})

	result, err := org.Run()
	if err != nil {
		t.Fatal(err)
	}

	for _, a := range result.Actions {
		if a.FileName == "song.mp3" && a.Reason != organizer.ReasonExcluded {
			t.Errorf("mp3 should be excluded, got reason %d", a.Reason)
		}
		if a.FileName == "doc.pdf" && a.Reason != organizer.ReasonOrganized {
			t.Errorf("pdf should be organized, got reason %d", a.Reason)
		}
	}
}

func TestOrganizer_Run_Recursive(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()

	subdir := filepath.Join(dir, "subdir")
	if err := os.Mkdir(subdir, os.ModePerm); err != nil {
		t.Fatal(err)
	}

	createTestFile(t, dir, "top.mp3")
	createTestFile(t, subdir, "nested.pdf")

	org := organizer.NewOrganizer(newMockResolver(), organizer.OrganizerConfig{
		InputFolder:       dir,
		OutputFolder:      dir,
		Preview:           true,
		Recursive:         true,
		IgnoreHiddenFiles: true,
	})

	result, err := org.Run()
	if err != nil {
		t.Fatal(err)
	}

	fileNames := make(map[string]bool)
	for _, a := range result.Actions {
		fileNames[a.FileName] = true
	}

	if !fileNames["top.mp3"] {
		t.Error("expected top.mp3 in results")
	}
	if !fileNames["nested.pdf"] {
		t.Error("expected nested.pdf in results")
	}
}

func TestOrganizer_Run_HiddenFiles(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()

	createTestFile(t, dir, ".hidden")
	createTestFile(t, dir, "visible.mp3")

	org := organizer.NewOrganizer(newMockResolver(), organizer.OrganizerConfig{
		InputFolder:       dir,
		OutputFolder:      dir,
		Preview:           true,
		IgnoreHiddenFiles: false,
	})

	result, err := org.Run()
	if err != nil {
		t.Fatal(err)
	}

	for _, a := range result.Actions {
		if a.FileName == ".hidden" && a.Reason != organizer.ReasonHidden {
			t.Errorf(".hidden reason = %d, want ReasonHidden", a.Reason)
		}
	}
}
