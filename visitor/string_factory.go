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

func (f *stringFactory) BestPrefix(next, _ []byte) []byte {
	var bestPrefix []byte

	for _, definition := range f.definitions {
		for _, prefix := range definition.Prefixes {
			if bytes.HasPrefix(next, prefix) && len(bestPrefix) < len(prefix) {
				bestPrefix = prefix
			}
		}
	}

	return bestPrefix
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

func (f *stringFactory) findDefinition(prefix []byte) types.StringDefinition {
	for _, definition := range f.definitions {
		for _, s := range definition.Prefixes {
			if bytes.Equal(s, prefix) {
				return definition
			}
		}
	}

	return types.StringDefinition{
		Prefixes:  nil,
		Postfix:   nil,
		Skip:      nil,
		Multiline: false,
	}
}
