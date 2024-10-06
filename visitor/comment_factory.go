package visitor

import (
	"bytes"

	"github.com/falconandy/strip-literal/types"
)

func NewSingleLineCommentFactory(prefix string) types.VisitorFactory {
	return &singleLineFactory{
		prefix: []byte(prefix),
	}
}

type singleLineFactory struct {
	prefix []byte
}

func (f *singleLineFactory) BestPrefix(next, _ []byte) []byte {
	if bytes.HasPrefix(next, f.prefix) {
		return f.prefix
	}

	return nil
}

func (f *singleLineFactory) CreateVisitor(prefix []byte) types.SegmentVisitor {
	return &singleLineCommentVisitor{
		baseVisitor: newBaseVisitor(types.SegmentTypeCommentSingleLine, len(prefix)),
	}
}

func NewMultiLineCommentFactory(prefix, postfix string, supportsNesting bool) types.VisitorFactory {
	return &multiLineCommentFactory{
		prefix:          []byte(prefix),
		postfix:         []byte(postfix),
		supportsNesting: supportsNesting,
	}
}

type multiLineCommentFactory struct {
	prefix          []byte
	postfix         []byte
	supportsNesting bool
}

func (f *multiLineCommentFactory) BestPrefix(next, _ []byte) []byte {
	if bytes.HasPrefix(next, f.prefix) {
		return f.prefix
	}

	return nil
}

func (f *multiLineCommentFactory) CreateVisitor(prefix []byte) types.SegmentVisitor {
	return &multiLineCommentVisitor{
		baseVisitor:     newBaseVisitor(types.SegmentTypeCommentMultiLine, len(prefix)),
		prefix:          f.prefix,
		postfix:         f.postfix,
		supportsNesting: f.supportsNesting,
		nestLevel:       0,
	}
}
