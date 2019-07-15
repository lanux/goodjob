package cas

import (
	"regexp"
)

type Matcher interface {
	Match(key string, pattern string) bool
	MatchAny(key string, patterns *[]string) bool
}

type RegexMatch struct {
	Maps map[string]*regexp.Regexp
}

func (m *RegexMatch) MatchAny(key string, patterns *[]string) bool {
	for _, pattern := range *patterns {
		if m.Match(key, pattern) {
			return true
		}
	}
	return false
}

func (m *RegexMatch) Match(key string, pattern string) bool {
	c, ok := m.Maps[pattern]
	if !ok {
		c, err := regexp.Compile(pattern)
		if err != nil {
			panic(err)
		}
		m.Maps[pattern] = c
	}
	return c.MatchString(key)
}
