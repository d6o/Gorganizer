package store

// Rule represents a mapping from a file extension to a destination folder.
type Rule struct {
	Extension string
	Folder    string
}
