package visitor

import (
	"bytes"

	"github.com/falconandy/strip-literal/types"
)

func NewStringFactory(definitions ...types.StringDefinition) types.StringFactory {
	return &stringFactory{
		definitions: definitions,
		codeFactory: nil,
	}
}

type stringFactory struct {
	definitions []types.StringDefinition
	codeFactory types.CodeFactory
}

func (f *stringFactory) BestPrefixLen(next, _ []byte) int {
	var bestPrefixLen int

	for _, definition := range f.definitions {
		for _, prefix := range definition.Prefixes {
			if bytes.HasPrefix(next, prefix) && bestPrefixLen < len(prefix) {
				bestPrefixLen = len(prefix)
			}
		}
	}

	return bestPrefixLen
}

func (f *stringFactory) CreateVisitor(prefix []byte) types.SegmentVisitor {
	definition := f.findDefinition(prefix)

	return &stringVisitor{
		baseVisitor:   newBaseVisitor(types.SegmentTypeString, len(prefix)),
		definition:    definition,
		codeFactory:   f.codeFactory,
		pendingPrefix: nil,
	}
}

func (f *stringFactory) SetTemplateCodeFactory(codeFactory types.CodeFactory) {
	f.codeFactory = codeFactory
}

func (f *stringFactory) findDefinition(prefix []byte) types.StringDefinition {
	for _, definition := range f.definitions {
		for _, s := range definition.Prefixes {
			if bytes.Equal(s, prefix) {
				return definition
			}
		}
	}

	return types.StringDefinition{
		Prefixes:        nil,
		Postfix:         nil,
		Skip:            nil,
		Multiline:       false,
		TemplatePrefix:  nil,
		TemplatePostfix: nil,
	}
}
