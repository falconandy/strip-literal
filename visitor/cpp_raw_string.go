package visitor

import (
	"bytes"

	"github.com/falconandy/strip-literal/types"
)

type cppRawStringVisitor struct {
	*baseVisitor
	postfix []byte
}

func (s *cppRawStringVisitor) Visit(next, _ []byte) (types.SegmentVisitor, int) {
	postfixIndex := bytes.Index(next, s.postfix)
	if postfixIndex < 0 {
		return nil, s.Take(len(next))
	}

	return nil, s.Take(postfixIndex) + s.TakePostfix(len(s.postfix))
}
