package visitor

import (
	"bytes"

	"github.com/falconandy/strip-literal/types"
)

var cppRawStringPrefixes = [][]byte{
	[]byte(`R"`),
	[]byte(`LR"`),
	[]byte(`u8R"`),
	[]byte(`uR"`),
	[]byte(`UR"`),
}

func NewCPPRawStringFactory() types.VisitorFactory {
	return &cppRawStringFactory{}
}

type cppRawStringFactory struct {
}

func (f *cppRawStringFactory) BestPrefix(next, _ []byte) []byte {
	var bestPrefix []byte
	for _, prefix := range cppRawStringPrefixes {
		if bytes.HasPrefix(next, prefix) {
			bestPrefix = prefix
			break
		}
	}

	if len(bestPrefix) == 0 {
		return nil
	}

	index := bytes.IndexByte(next, '(')
	if index < 0 {
		return nil
	}

	return next[:index+1]
}

func (f *cppRawStringFactory) CreateVisitor(prefix []byte) types.SegmentVisitor {
	index := bytes.IndexByte(prefix, '"')
	postfix := make([]byte, len(prefix)-index)
	postfix[0] = ')'
	copy(postfix[1:], prefix[index+1:len(prefix)-1])
	postfix[len(postfix)-1] = '"'

	return &cppRawStringVisitor{
		baseVisitor: newBaseVisitor(types.SegmentTypeString, len(prefix)),
		postfix:     postfix,
	}
}
