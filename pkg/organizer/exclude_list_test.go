package organizer_test

import (
	"testing"

	"github.com/d6o/Gorganizer/pkg/organizer"
)

func TestExcludeList_Contains(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		list organizer.ExcludeList
		ext  string
		want bool
	}{
		{"match found", organizer.ExcludeList{"pdf", "mp3"}, "pdf", true},
		{"no match", organizer.ExcludeList{"pdf", "mp3"}, "jpg", false},
		{"empty extension", organizer.ExcludeList{"pdf"}, "", false},
		{"empty list", organizer.ExcludeList{}, "pdf", false},
		{"nil list", nil, "pdf", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.list.Contains(tt.ext)
			if got != tt.want {
				t.Errorf("Contains(%q) = %v, want %v", tt.ext, got, tt.want)
			}
		})
	}
}
