package types

type SegmentType string

const (
	SegmentTypeCode              SegmentType = "code"
	SegmentTypeString            SegmentType = "literal/string"
	SegmentTypeRegexp            SegmentType = "literal/regexp"
	SegmentTypeCommentSingleLine SegmentType = "comment/single-line"
	SegmentTypeCommentMultiLine  SegmentType = "comment/multi-line"
)

type Segment struct {
	Type          SegmentType
	Position      int32
	Length        int32
	PrefixLength  int32
	PostfixLength int32
}

func (s Segment) IsComment() bool {
	return s.Type == SegmentTypeCommentSingleLine || s.Type == SegmentTypeCommentMultiLine
}

func (s Segment) IsString() bool {
	return s.Type == SegmentTypeString
}

func (s Segment) IsRegexp() bool {
	return s.Type == SegmentTypeRegexp
}
