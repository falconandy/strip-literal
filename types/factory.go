package types

type VisitorFactory interface {
	BestPrefix(next, prev []byte) []byte
	CreateVisitor(prefix []byte) SegmentVisitor
}

type CodeFactory interface {
	VisitorFactory
}

type StringFactory interface {
	VisitorFactory
}
