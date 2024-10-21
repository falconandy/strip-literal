package types

type VisitorFactory interface {
	BestPrefixLen(next, prev []byte) int
	CreateVisitor(prefix, next []byte) SegmentVisitor
}

type CodeFactory interface {
	VisitorFactory
	CreateStringTemplateVisitor(templatePostfix []byte) SegmentVisitor
}

type StringFactory interface {
	VisitorFactory
	SetTemplateCodeFactory(codeFactory CodeFactory)
}
