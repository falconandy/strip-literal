package types

type VisitorFactory interface {
	BestPrefix(next, prev []byte) []byte
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
