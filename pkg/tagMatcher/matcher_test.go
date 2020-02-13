package pkg

import (
	regexp "regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatch(t *testing.T) {
	exp, err := regexp.Compile("^match")
	matchAgainst := "match"
	if err != nil {
		t.Fatalf("bad test code bro")
	}
	tagMatcher := TagMatch{
		pattern: exp,
	}
	assert.True(t, tagMatcher.Match(matchAgainst))
	assert.False(t, tagMatcher.Match("wrong string"))
}

func TestGlobPatternMatch(t *testing.T) {
	exp, err := regexp.Compile("/path/*")
	matchAgainst := "/path/that/should/match"
	if err != nil {
		t.Fatalf("bad test code bro")
	}
	tagMatcher := TagMatch{
		pattern: exp,
	}
	assert.True(t, tagMatcher.Match(matchAgainst))
	assert.False(t, tagMatcher.Match("/dont/match/this"))
}
