package main

import (
	"regexp"
	"strings"
)

func testRule(rule string) bool {
	r, _ := regexp.Compile("([a-zA-Z]+):([a-zA-Z]+)")

	return r.MatchString(rule)
}

func insertRule(rule string) {

	if testRule(rule) == false {
		return
	}

	ruleArr := strings.Split(rule, ":")

	boltSet("ext:"+ruleArr[0],ruleArr[1])
}

func deleteRule(rule string) {
	boltDelete("ext:"+rule)
}

func showRules() {
	boltScanExt()
}