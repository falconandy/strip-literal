package visitor

import (
	"github.com/falconandy/strip-literal/types"
)

type regexpVisitor struct {
	squareBracketLevel int
}

func (s *regexpVisitor) Visit(next, _ []byte) (types.SegmentVisitor, types.VisitResult) {
	switch {
	case next[0] == '\n' || next[0] == '\r':
		return nil, types.VisitResult{}
	case len(next) >= 2 && next[0] == '\\' && next[1] == '/':
		return s, types.VisitResult{InnerLength: 2}
	case len(next) >= 2 && next[0] == '\\' && next[1] == '[':
		return s, types.VisitResult{InnerLength: 2}
	case next[0] == '[':
		s.squareBracketLevel++

		return s, types.VisitResult{InnerLength: 1}
	case len(next) >= 2 && next[0] == '\\' && next[1] == ']':
		return s, types.VisitResult{InnerLength: 2}
	case next[0] == ']':
		if s.squareBracketLevel > 0 {
			s.squareBracketLevel--
		}

		return s, types.VisitResult{InnerLength: 1}
	case next[0] == '/' && s.squareBracketLevel == 0:
		i := 1
		for i < len(next) && 'a' <= next[i] && next[i] <= 'z' {
			i++
		}

		return nil, types.VisitResult{PostfixLength: i}
	}

	return s, types.VisitResult{InnerLength: 1}
}

func (s *regexpVisitor) SegmentType() types.SegmentType {
	return types.SegmentTypeRegexp
}
