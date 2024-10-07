package visitor

import (
	"bytes"

	"github.com/falconandy/strip-literal/types"
)

func NewRegexpFactory() types.VisitorFactory {
	return &regexpFactory{}
}

type regexpFactory struct{}

func (f *regexpFactory) BestPrefix(next, prev []byte) []byte {
	if next[0] != '/' {
		return nil
	}

	prev = bytes.TrimRight(prev, " \t\f")
	if len(prev) == 0 {
		return nil
	}

	switch prev[len(prev)-1] {
	case '=', ',', ';', '(', '{', '}':
		return next[:1]
	}

	return nil
}

func (f *regexpFactory) CreateVisitor(prefix []byte) types.SegmentVisitor {
	return &regexpVisitor{
		baseVisitor:        newBaseVisitor(types.SegmentTypeRegexp, len(prefix)),
		squareBracketLevel: 0,
	}
}
