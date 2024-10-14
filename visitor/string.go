package visitor

import (
	"bytes"

	"github.com/falconandy/strip-literal/types"
)

type stringVisitor struct {
	*baseVisitor
	definition    types.StringDefinition
	codeFactory   types.CodeFactory
	pendingPrefix []byte
}

func (s *stringVisitor) Visit(next, _ []byte) (types.SegmentVisitor, int) {
	if len(s.pendingPrefix) > 0 {
		pendingPrefix := s.pendingPrefix
		s.pendingPrefix = nil

		if bytes.HasPrefix(next, pendingPrefix) {
			return s, s.TakePrefix(len(pendingPrefix))
		}
	}

	if bytes.HasPrefix(next, s.definition.Postfix) {
		return nil, s.TakePostfix(len(s.definition.Postfix))
	}

	for _, skip := range s.definition.Skip {
		if bytes.HasPrefix(next, skip) {
			return s, s.Take(len(skip))
		}
	}

	if !s.definition.Multiline {
		if next[0] == '\n' || next[0] == '\r' {
			return nil, 0
		}
	}

	if len(s.definition.TemplatePrefix) > 0 {
		if bytes.HasPrefix(next, s.definition.TemplatePrefix) {
			s.pendingPrefix = s.definition.TemplatePostfix

			return s.codeFactory.CreateStringTemplateVisitor(
					s.definition.TemplatePrefix,
					s.definition.TemplatePostfix),
				s.TakePostfix(len(s.definition.TemplatePrefix))
		}
	}

	return s, s.Take(1)
}
