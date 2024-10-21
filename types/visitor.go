package types

type VisitResult struct {
	PrefixLength  int
	InnerLength   int
	PostfixLength int
}

type SegmentVisitor interface {
	SegmentType() SegmentType
	Visit(next, prev []byte) (SegmentVisitor, VisitResult)
}

func (r VisitResult) HasData() bool {
	return r.PrefixLength > 0 || r.InnerLength > 0 || r.PostfixLength > 0
}

func (r VisitResult) Add(other VisitResult) VisitResult {
	return VisitResult{
		PrefixLength:  r.PrefixLength + other.PrefixLength,
		InnerLength:   r.InnerLength + other.InnerLength,
		PostfixLength: r.PostfixLength + other.PostfixLength,
	}
}
