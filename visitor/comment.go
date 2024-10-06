package visitor

import (
	"bytes"

	"github.com/falconandy/strip-literal/types"
)

type singleLineCommentVisitor struct {
	*baseVisitor
}

func (s *singleLineCommentVisitor) Visit(next, _ []byte) (types.SegmentVisitor, int) {
	newLineIndex := bytes.IndexAny(next, "\n\r")
	if newLineIndex < 0 {
		return nil, s.Take(len(next))
	}

	return nil, s.Take(newLineIndex)
}

type multiLineCommentVisitor struct {
	*baseVisitor
	prefix          []byte
	postfix         []byte
	supportsNesting bool
	nestLevel       int
}

func (s *multiLineCommentVisitor) Visit(next, _ []byte) (types.SegmentVisitor, int) {
	switch {
	case bytes.HasPrefix(next, s.prefix):
		if s.supportsNesting {
			s.nestLevel++
		}

		return s, s.Take(len(s.prefix))
	case bytes.HasPrefix(next, s.postfix):
		if s.nestLevel > 0 {
			s.nestLevel--

			return s, s.Take(len(s.postfix))
		}

		return nil, s.TakePostfix(len(s.postfix))
	default:
		return s, s.Take(1)
	}
}
