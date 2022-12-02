package Utilities

import "regexp"

type ExtensionString = string

func SplitByEmptyNewline(input string) []string {
	StringWithOnlyNewline := regexp.
		MustCompile("\r\n").
		ReplaceAllString(input, "\n")

	return regexp.
		MustCompile(`\n\s*\n`).
		Split(StringWithOnlyNewline, -1)
}

func SplitByNewline(input string) []string {
	StringWithOnlyNewline := regexp.
		MustCompile("\r\n").
		ReplaceAllString(input, "\n")

	return regexp.
		MustCompile(`\n`).
		Split(StringWithOnlyNewline, -1)
}
