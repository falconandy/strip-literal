package visitor

import (
	"github.com/falconandy/strip-literal/types"
)

func NewCodeFactory(factories ...types.VisitorFactory) types.CodeFactory {
	factory := &codeFactory{
		factories: factories,
		brackets:  nil,
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
	brackets  []bracketPair
}

func (f *codeFactory) BestPrefixLen([]byte, []byte) int {
	return 0
}

func (f *codeFactory) CreateVisitor(prefix []byte) types.SegmentVisitor {
	return &codeVisitor{
		baseVisitor:     newBaseVisitor(types.SegmentTypeCode, len(prefix)),
		f:               f,
		factories:       f.factories,
		nestedBrackets:  make([][]int, curlyBracketIndex+1),
		templatePrefix:  nil,
		templatePostfix: nil,
	}
}

func (f *codeFactory) CreateStringTemplateVisitor(templatePrefix, templatePostfix []byte) types.SegmentVisitor {
	return &codeVisitor{
		baseVisitor:     newBaseVisitor(types.SegmentTypeCode, 0),
		f:               f,
		factories:       f.factories,
		nestedBrackets:  make([][]int, curlyBracketIndex+1),
		templatePrefix:  templatePrefix,
		templatePostfix: templatePostfix,
	}
}

func (f *codeFactory) Brackets() []bracketPair {
	return f.brackets
}
