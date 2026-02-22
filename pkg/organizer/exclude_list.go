package organizer

// ExcludeList is a list of file extensions to exclude from organizing.
type ExcludeList []string

// Contains reports whether ext is in the exclude list.
// Returns false for empty extensions.
func (e ExcludeList) Contains(ext string) bool {
	if ext == "" {
		return false
	}
	for _, item := range e {
		if item == ext {
			return true
		}
	}
	return false
}
