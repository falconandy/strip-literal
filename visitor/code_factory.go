package visitor

import (
	"github.com/falconandy/strip-literal/types"
)

func NewCodeFactory(factories ...types.VisitorFactory) types.CodeFactory {
	factory := &codeFactory{
		factories: factories,
		brackets:  nil,
	}

	return factory
}

type codeFactory struct {
	factories []types.VisitorFactory
	brackets  []bracketPair
}

func (f *codeFactory) BestPrefix([]byte, []byte) []byte {
	return nil
}

func (f *codeFactory) CreateVisitor(prefix []byte) types.SegmentVisitor {
	return &codeVisitor{
		baseVisitor:    newBaseVisitor(types.SegmentTypeCode, len(prefix)),
		f:              f,
		factories:      f.factories,
		nestedBrackets: make([][]int, curlyBracketIndex+1),
	}
}

func (f *codeFactory) Brackets() []bracketPair {
	return f.brackets
}
