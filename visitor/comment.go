package visitor

import (
	"bytes"

	"github.com/falconandy/strip-literal/types"
)

type singleLineCommentVisitor struct {
}

func (s *singleLineCommentVisitor) Visit(next, _ []byte) (types.SegmentVisitor, types.VisitResult) {
	newLineIndex := bytes.IndexAny(next, "\n\r")
	if newLineIndex < 0 {
		return nil, types.VisitResult{InnerLength: len(next)}
	}

	return nil, types.VisitResult{InnerLength: newLineIndex}
}

func (s *singleLineCommentVisitor) SegmentType() types.SegmentType {
	return types.SegmentTypeCommentSingleLine
}

type multiLineCommentVisitor struct {
	prefix          []byte
	postfix         []byte
	supportsNesting bool
	nestLevel       int
}

func (s *multiLineCommentVisitor) Visit(next, _ []byte) (types.SegmentVisitor, types.VisitResult) {
	switch {
	case bytes.HasPrefix(next, s.prefix):
		if s.supportsNesting {
			s.nestLevel++
		}

		return s, types.VisitResult{InnerLength: len(s.prefix)}
	case bytes.HasPrefix(next, s.postfix):
		if s.nestLevel > 0 {
			s.nestLevel--

			return s, types.VisitResult{InnerLength: len(s.postfix)}
		}

		return nil, types.VisitResult{PostfixLength: len(s.postfix)}
	default:
		return s, types.VisitResult{InnerLength: 1}
	}
}

func (s *multiLineCommentVisitor) SegmentType() types.SegmentType {
	return types.SegmentTypeCommentMultiLine
}
