package types

type VisitorFactory interface {
	BestPrefixLen(next, prev []byte) int
	CreateVisitor(prefix []byte) SegmentVisitor
}

type CodeFactory interface {
	VisitorFactory
	CreateStringTemplateVisitor(templatePrefix, templatePostfix []byte) SegmentVisitor
}

type StringFactory interface {
	VisitorFactory
	SetTemplateCodeFactory(codeFactory CodeFactory)
}
