package main

import (
	"fmt"
	"strings"
)

func insertRule(rule string) error {
	ruleArr := strings.Split(rule, ":")
	if len(ruleArr) != 2 {
		return fmt.Errorf("rule must have exactly one colon ':', found %d in %#v", strings.Count(rule, ":"), rule)
	}
	if len(ruleArr[0]) == 0 || len(ruleArr[1]) == 0 {
		return fmt.Errorf("rule must be in the format 'extension:folder', got %#v", rule)
	}

	return iniSet(ruleArr[0], ruleArr[1])
}

func deleteRule(rule string) {
	iniDelete(rule)
}

func showRules() {
	iniScanExt()
}
