package visitor

import (
	"github.com/falconandy/strip-literal/types"
)

func NewCodeFactory(factories ...types.VisitorFactory) types.CodeFactory {
	factory := &codeFactory{
		factories: factories,
	}

	for _, f := range factories {
		if stringFactory, ok := f.(types.StringFactory); ok {
			stringFactory.SetTemplateCodeFactory(factory)
		}
	}

	return factory
}

type codeFactory struct {
	factories []types.VisitorFactory
}

func (f *codeFactory) BestPrefixLen([]byte, []byte) int {
	return 0
}

func (f *codeFactory) CreateVisitor([]byte, []byte) types.SegmentVisitor {
	return &codeVisitor{
		f:               f,
		factories:       f.factories,
		nestedBrackets:  [curlyBracketIndex + 1]int{},
		templatePostfix: nil,
	}
}

func (f *codeFactory) CreateStringTemplateVisitor(templatePostfix []byte) types.SegmentVisitor {
	return &codeVisitor{
		f:               f,
		factories:       f.factories,
		nestedBrackets:  [curlyBracketIndex + 1]int{},
		templatePostfix: templatePostfix,
	}
}
