package types

type SegmentVisitor interface {
	SegmentType() SegmentType
	Visit(next, prev []byte) (SegmentVisitor, int)
	PopVisited() (length, prefixLength, postfixLength int)
}
