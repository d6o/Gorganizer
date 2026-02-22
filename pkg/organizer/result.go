package organizer

// ActionReason describes why a file was categorized in a particular way.
type ActionReason int

const (
	// ReasonOrganized means the file matched a rule and was (or would be) moved.
	ReasonOrganized ActionReason = iota
	// ReasonExcluded means the file's extension is in the exclude list.
	ReasonExcluded
	// ReasonHidden means the file is a hidden file (name starts with ".").
	ReasonHidden
	// ReasonUnknownExtension means no rule matched the file's extension.
	ReasonUnknownExtension
)

// FileAction describes what happened (or would happen) to a single file
// during an organize operation.
type FileAction struct {
	FileName    string
	Destination string
	Reason      ActionReason
	Moved       bool
}

// OrganizeResult is the structured output of an organize operation,
// containing one FileAction per file encountered.
type OrganizeResult struct {
	Actions []FileAction
}
