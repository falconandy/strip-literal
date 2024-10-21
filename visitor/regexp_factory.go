package visitor

import (
	"bytes"

	"github.com/falconandy/strip-literal/types"
)

func NewRegexpFactory() types.VisitorFactory {
	return &regexpFactory{}
}

type regexpFactory struct{}

func (f *regexpFactory) BestPrefixLen(next, prev []byte) int {
	if next[0] != '/' {
		return 0
	}

	prev = bytes.TrimRight(prev, " \t\f")
	if len(prev) == 0 {
		return 0
	}

	switch prev[len(prev)-1] {
	case '=', ',', ';', '(', '{', '}':
		return 1
	}

	return 0
}

func (f *regexpFactory) CreateVisitor([]byte, []byte) types.SegmentVisitor {
	return &regexpVisitor{
		squareBracketLevel: 0,
	}
}
