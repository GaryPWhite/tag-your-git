package pkg

import (
	"encoding/json"
	"fmt"
	"regexp"
)

type rawMatcher struct {
	Patterns []string
	Tags     []string
}

// Unmarshal turns JSON into a TagMatcher array
func Unmarshal(rawJSON []byte) ([]TagMatch, error) {
	tagMatchList := []TagMatch{}
	matchers := []rawMatcher{}
	err := json.Unmarshal(rawJSON, &matchers)
	if err != nil {
		fmt.Printf("err unmarshalling JSON")
		return nil, err
	}
	// for each pattern to match, ensure a corresponding TagMatcher
	for _, rm := range matchers {
		for _, pattern := range rm.Patterns {
			patternRegExp, err := regexp.Compile(pattern)
			if err != nil {
				fmt.Printf("ERROR: Could not parse pattern '%s' as regular expression", pattern)
				return nil, err
			}
			tagMatchList = append(tagMatchList,
				TagMatch{
					pattern: patternRegExp,
					tags:    rm.Tags,
				})
		}
	}
	return tagMatchList, nil
}
