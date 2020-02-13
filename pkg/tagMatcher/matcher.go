package pkg

import (
	regexp "regexp"
)

// TagMatch will hold a regular expression and a set of tags applying to that expression
type TagMatch struct {
	pattern *regexp.Regexp
	tags    []string
}

// Match will check if "path" supplied to tagmatcher will match
// @param path to match against
func (t *TagMatch) Match(path string) bool {
	return t.pattern.Match([]byte(path))
}

// GetTags will return tags contained for a TagMatch struct
func (t *TagMatch) GetTags() []string {
	return t.tags
}
