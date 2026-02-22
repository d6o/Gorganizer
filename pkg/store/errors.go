package store

import "errors"

// ErrInvalidRuleFormat is returned when a rule string does not contain
// exactly one colon separator (expected format: "ext:folder").
var ErrInvalidRuleFormat = errors.New("rule must have exactly one colon separator")

// ErrEmptyRuleComponent is returned when either the extension or folder
// part of a rule is empty.
var ErrEmptyRuleComponent = errors.New("rule extension and folder must not be empty")
