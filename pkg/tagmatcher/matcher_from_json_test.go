package tagmatcher

import (
	regexp "regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalOnEmpty(t *testing.T) {
	// create json struct example
	rawJSON := []byte(`
		[{
			"patterns": [],
			"tags": []
		}]
	`)
	matchers, err := Unmarshal(rawJSON)
	if err != nil {
		t.Errorf("Could not Unmarshal JSON %s", err)
	}
	// matchers should be empty.
	assert.Empty(t, matchers)
}

func TestUnmarshalOnePatternNoTag(t *testing.T) {
	// create json struct
	rawJSON := []byte(`
		[{
			"patterns": ["anything/in/the/project*"],
			"tags": []
		}]
	`)
	matchers, err := Unmarshal(rawJSON)
	if err != nil {
		t.Errorf("Could not Unmarshal JSON %s", err)
	}
	// matchers should have one, with a pattern and no tags
	assert.Equal(t, 1, len(matchers))
	assert.IsType(t, &regexp.Regexp{}, matchers[0].pattern)
	assert.Empty(t, len(matchers[0].tags))
}

func TestUnmarshalOnePatternOneTag(t *testing.T) {
	// create json struct
	rawJSON := []byte(`
		[{
			"patterns": ["scibidipishibiboombadomdoop"],
			"tags": ["scat man"]
		}]
	`)
	matchers, err := Unmarshal(rawJSON)
	if err != nil {
		t.Errorf("Could not Unmarshal JSON %s", err)
	}
	// matchers should have one, with a pattern and one tag
	assert.Equal(t, 1, len(matchers))
	assert.IsType(t, &regexp.Regexp{}, matchers[0].pattern)
	assert.Equal(t, len(matchers[0].tags), 1)
}

func TestUnmarshalOnePatternManyTags(t *testing.T) {
	// create json struct
	rawJSON := []byte(`
		[{
			"patterns": ["im david pumpkins, man!"],
			"tags": ["any", "questions?"]
		}]
	`)
	matchers, err := Unmarshal(rawJSON)
	if err != nil {
		t.Errorf("Could not Unmarshal JSON %s", err)
	}
	// matchers should have one, with a pattern and no tags
	assert.Equal(t, 1, len(matchers))
	assert.IsType(t, &regexp.Regexp{}, matchers[0].pattern)
	assert.Equal(t, len(matchers[0].tags), 2)
}

func TestUnmarshalManyPatternsNoTags(t *testing.T) {
	// create json struct
	rawJSON := []byte(`
		[{
			"patterns": ["one fish", "two fish"],
			"tags": []
		}]
	`)
	matchers, err := Unmarshal(rawJSON)
	if err != nil {
		t.Errorf("Could not Unmarshal JSON %s", err)
	}
	// matchers should have one, with a pattern and no tags
	assert.Equal(t, 2, len(matchers))
	assert.IsType(t, &regexp.Regexp{}, matchers[0].pattern)
	assert.IsType(t, &regexp.Regexp{}, matchers[1].pattern)
	assert.Empty(t, len(matchers[0].tags))
	assert.Empty(t, len(matchers[1].tags))
}

func TestUnmarshalManyPatternsOneTag(t *testing.T) {
	// create json struct
	rawJSON := []byte(`
		[{
			"patterns": ["comanche", "roads"],
			"tags": ["cities"]
		}]
	`)
	matchers, err := Unmarshal(rawJSON)
	if err != nil {
		t.Errorf("Could not Unmarshal JSON %s", err)
	}
	// matchers should have one, with a pattern and no tags
	assert.Equal(t, 2, len(matchers))
	assert.IsType(t, &regexp.Regexp{}, matchers[0].pattern)
	assert.IsType(t, &regexp.Regexp{}, matchers[1].pattern)
	assert.Equal(t, len(matchers[0].tags), 1)
	assert.Equal(t, len(matchers[1].tags), 1)
}

func TestUnmarshalManyPatternsManyTags(t *testing.T) {
	// create json struct
	rawJSON := []byte(`
		[{
			"patterns": ["baby", "yoda"],
			"tags": ["save", "him"]
		}]
	`)
	matchers, err := Unmarshal(rawJSON)
	if err != nil {
		t.Errorf("Could not Unmarshal JSON %s", err)
	}
	// matchers should have one, with a pattern and no tags
	assert.Equal(t, 2, len(matchers))
	assert.IsType(t, &regexp.Regexp{}, matchers[0].pattern)
	assert.IsType(t, &regexp.Regexp{}, matchers[1].pattern)
	assert.Equal(t, len(matchers[0].tags), 2)
	assert.Equal(t, len(matchers[1].tags), 2)
}
