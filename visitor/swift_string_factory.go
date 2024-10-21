package visitor

import (
	"bytes"

	"github.com/falconandy/strip-literal/types"
)

var (
	swiftTemplatePrefixPrefix  = []byte{'\\'}
	swiftTemplatePrefixPostfix = []byte{'('}
	swiftTemplatePostfix       = []byte(")")
	swiftSingleLineStringSkip  = [][]byte{[]byte("\\\""), []byte("\\\\")}
	swiftMultiLineStringSkip   = [][]byte{[]byte("\\\""), []byte("\\\\"), []byte("\\\n\r"), []byte("\\\r\n"), []byte("\\\n"), []byte("\\\r")}
)

func NewSwiftStringFactory() types.StringFactory {
	return &swiftStringFactory{}
}

type swiftStringFactory struct {
	codeFactory types.CodeFactory
}

func (f *swiftStringFactory) BestPrefixLen(next, _ []byte) int {
	if next[0] == '"' {
		if len(next) >= 3 && next[1] == '"' && next[2] == '"' {
			return 3
		}
		return 1
	}

	if next[0] != '#' {
		return 0
	}

	index := bytes.IndexByte(next, '"')
	if index < 0 {
		return 0
	}

	for i := range index - 1 {
		if next[i+1] != '#' {
			return 0
		}
	}

	if index+2 < len(next) && next[index+1] == '"' && next[index+2] == '"' {
		return index + 3
	}

	return index + 1
}

func (f *swiftStringFactory) CreateVisitor(prefix, next []byte) types.SegmentVisitor {
	multiline := len(prefix) >= 3 && prefix[len(prefix)-1] == '"' && prefix[len(prefix)-2] == '"' && prefix[len(prefix)-3] == '"'

	skip := swiftSingleLineStringSkip
	if multiline {
		skip = swiftMultiLineStringSkip
	}

	var postfix []byte
	var templatePrefix []byte
	if multiline {
		postfixIndex := FindSubData(next, prefix[len(prefix)-3:], nil, prefix[:len(prefix)-3])
		if postfixIndex >= 0 {
			postfix = next[postfixIndex : postfixIndex+len(prefix)]
		}
		templatePrefixIndex := FindSubData(next, swiftTemplatePrefixPrefix, prefix[:len(prefix)-3], swiftTemplatePrefixPostfix)
		if templatePrefixIndex >= 0 {
			templatePrefix = next[templatePrefixIndex : templatePrefixIndex+len(prefix)-1]
		}
	} else {
		postfixIndex := FindSubData(next, prefix[len(prefix)-1:], nil, prefix[:len(prefix)-1])
		if postfixIndex >= 0 {
			postfix = next[postfixIndex : postfixIndex+len(prefix)]
		}
		templatePrefixIndex := FindSubData(next, swiftTemplatePrefixPrefix, prefix[:len(prefix)-1], swiftTemplatePrefixPostfix)
		if templatePrefixIndex >= 0 {
			templatePrefix = next[templatePrefixIndex : templatePrefixIndex+len(prefix)+1]
		}
	}

	return &stringVisitor{
		definition: types.StringDefinition{
			Prefixes:        [][]byte{prefix},
			Postfix:         postfix,
			Skip:            skip,
			Multiline:       multiline,
			TemplatePrefix:  templatePrefix,
			TemplatePostfix: swiftTemplatePostfix,
		},
		codeFactory:   f.codeFactory,
		pendingPrefix: nil,
	}
}

func (f *swiftStringFactory) SetTemplateCodeFactory(codeFactory types.CodeFactory) {
	f.codeFactory = codeFactory
}
