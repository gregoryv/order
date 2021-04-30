package order

import "regexp"

func LinesByPattern(lines, patterns []string) *ByPattern {
	return &ByPattern{
		lines:    lines,
		patterns: patterns,
	}
}

type ByPattern struct {
	lines    []string
	patterns []string
}

func (b ByPattern) Less(i, j int) bool {
	lineI := b.patternIndex(b.lines[i])
	lineJ := b.patternIndex(b.lines[j])
	if lineI == lineJ {
		return b.lines[i] < b.lines[j] // normal sorting
	}
	return lineI < lineJ
}

// patternIndex returns the line index of the first matching pattern
func (b ByPattern) patternIndex(v string) int {
	for i, pattern := range b.patterns {
		if match, _ := regexp.MatchString(pattern, v); match {
			return i
		}
	}
	return len(b.patterns)
}

func (b ByPattern) Len() int { return len(b.lines) }

func (b ByPattern) Swap(i, j int) {
	b.lines[i], b.lines[j] = b.lines[j], b.lines[i]
}
