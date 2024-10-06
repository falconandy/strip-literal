package visitor

import (
	"github.com/falconandy/strip-literal/types"
)

func newBaseVisitor(segmentType types.SegmentType, prefixLength int) *baseVisitor {
	visitor := &baseVisitor{
		segmentType:   segmentType,
		prefixLength:  0,
		postfixLength: 0,
		length:        0,
	}
	visitor.TakePrefix(prefixLength)

	return visitor
}

type baseVisitor struct {
	segmentType   types.SegmentType
	prefixLength  int
	postfixLength int
	length        int
}

func (s *baseVisitor) SegmentType() types.SegmentType {
	return s.segmentType
}

func (s *baseVisitor) PopVisited() (length, prefixLength, postfixLength int) {
	length, prefixLength, postfixLength = s.length, s.prefixLength, s.postfixLength
	s.length, s.prefixLength, s.postfixLength = 0, 0, 0

	return length, prefixLength, postfixLength
}

func (s *baseVisitor) Take(count int) int {
	s.length += count

	return count
}

func (s *baseVisitor) TakePostfix(count int) int {
	s.length += count
	s.postfixLength += count

	return count
}

func (s *baseVisitor) TakePrefix(count int) int {
	s.length += count
	s.prefixLength += count

	return count
}
