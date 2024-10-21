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

func (f *cppRawStringFactory) BestPrefixLen(next, _ []byte) int {
	var bestPrefix []byte
	for _, prefix := range cppRawStringPrefixes {
		if bytes.HasPrefix(next, prefix) {
			bestPrefix = prefix
			break
		}
	}

	if len(bestPrefix) == 0 {
		return 0
	}

	index := bytes.IndexByte(next, '(')
	if index < 0 {
		return 0
	}

	return index + 1
}

func (f *cppRawStringFactory) CreateVisitor(prefix []byte, _ []byte) types.SegmentVisitor {
	index := bytes.IndexByte(prefix, '"')
	postfix := prefix[index+1 : len(prefix)-1]

	return &cppRawStringVisitor{
		postfix: postfix,
	}
}
