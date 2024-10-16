package visitor

import (
	"bytes"

	"github.com/falconandy/strip-literal/types"
)

var (
	swiftTemplatePostfix      = []byte(")")
	swiftSingleLineStringSkip = [][]byte{[]byte(`\"`), []byte(`\\`)}
	swiftMultiLineStringSkip  = [][]byte{[]byte(`\"`), []byte(`\\`), []byte("\\\n\r"), []byte("\\\r\n"), []byte("\\\n"), []byte("\\\r")}
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

	if index+1 < len(next) && next[index+1] == '"' {
		index++
	}

	if index+1 < len(next) && next[index+1] == '"' {
		index++
	}

	return index + 1
}

func (f *swiftStringFactory) CreateVisitor(prefix []byte) types.SegmentVisitor {
	multiline := len(prefix) >= 3 && prefix[len(prefix)-1] == '"' && prefix[len(prefix)-2] == '"' && prefix[len(prefix)-3] == '"'

	skip := swiftSingleLineStringSkip
	if multiline {
		skip = swiftMultiLineStringSkip
	}

	postfix := make([]byte, len(prefix))
	var templatePrefix []byte
	if multiline {
		copy(postfix, prefix[len(prefix)-3:])
		copy(postfix[3:], prefix[:len(prefix)-3])

		templatePrefix = make([]byte, 2+len(prefix)-3)
		templatePrefix[0] = '\\'
		copy(templatePrefix[1:], prefix[:len(prefix)-3])
		templatePrefix[len(templatePrefix)-1] = '('
	} else {
		copy(postfix, prefix[len(prefix)-1:])
		copy(postfix[1:], prefix[:len(prefix)-1])

		templatePrefix = make([]byte, 2+len(prefix)-1)
		templatePrefix[0] = '\\'
		copy(templatePrefix[1:], prefix[:len(prefix)-1])
		templatePrefix[len(templatePrefix)-1] = '('
	}

	return &stringVisitor{
		baseVisitor: newBaseVisitor(types.SegmentTypeString, len(prefix)),
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
