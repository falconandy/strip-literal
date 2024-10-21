package visitor

import (
	"github.com/falconandy/strip-literal/types"
)

var (
	cppRawStringPostfixPrefix  = []byte{')'}
	cppRawStringPostfixPostfix = []byte{'"'}
)

type cppRawStringVisitor struct {
	postfix []byte
}

func (s *cppRawStringVisitor) Visit(next, _ []byte) (types.SegmentVisitor, types.VisitResult) {
	index := FindSubData(next, cppRawStringPostfixPrefix, s.postfix, cppRawStringPostfixPostfix)
	if index >= 0 {
		return nil, types.VisitResult{InnerLength: index, PostfixLength: len(s.postfix) + 2}
	}

	return nil, types.VisitResult{InnerLength: len(next)}
}

func (s *cppRawStringVisitor) SegmentType() types.SegmentType {
	return types.SegmentTypeString
}
