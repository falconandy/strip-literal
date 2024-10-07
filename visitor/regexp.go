package visitor

import (
	"github.com/falconandy/strip-literal/types"
)

type regexpVisitor struct {
	*baseVisitor
	squareBracketLevel int
}

func (s *regexpVisitor) Visit(next, _ []byte) (types.SegmentVisitor, int) {
	switch {
	case len(next) >= 2 && next[0] == '\\' && next[1] == '/':
		return s, s.Take(2)
	case len(next) >= 2 && next[0] == '\\' && next[1] == '[':
		return s, s.Take(2)
	case next[0] == '[':
		s.squareBracketLevel++

		return s, s.Take(1)
	case len(next) >= 2 && next[0] == '\\' && next[1] == ']':
		return s, s.Take(2)
	case next[0] == ']':
		if s.squareBracketLevel > 0 {
			s.squareBracketLevel--
		}

		return s, s.Take(1)
	case next[0] == '/' && s.squareBracketLevel == 0:
		i := 1
		for i < len(next) && 'a' <= next[i] && next[i] <= 'z' {
			i++
		}

		return nil, s.TakePostfix(i)
	}

	return s, s.Take(1)
}
