package visitor

import (
	"bytes"

	"github.com/falconandy/strip-literal/types"
)

type codeVisitor struct {
	*baseVisitor
	f               *codeFactory
	factories       []types.VisitorFactory
	templatePrefix  []byte
	templatePostfix []byte
	nestedBrackets  [][]int
}

type bracketPair struct {
	OpenedAt int32
	ClosedAt int32
}

type BracketIndex int

const (
	squareBracketIndex BracketIndex = 0
	parenthesesIndex   BracketIndex = 1
	curlyBracketIndex  BracketIndex = 2
)

func (s *codeVisitor) Visit(next, prev []byte) (types.SegmentVisitor, int) {
	bestFactory, bestPrefixLen := s.findBestFactory(next, prev)
	if bestFactory != nil {
		return bestFactory.CreateVisitor(next[:bestPrefixLen]), bestPrefixLen
	}

	if len(s.templatePostfix) > 0 && bytes.HasPrefix(next, s.templatePostfix) {
		isClosingBracket := false

		if len(s.templatePostfix) == 1 &&
			(s.templatePostfix[0] == ']' || s.templatePostfix[0] == ')' || s.templatePostfix[0] == '}') {
			switch {
			case s.templatePostfix[0] == ']' && len(s.nestedBrackets[squareBracketIndex]) > 0:
				isClosingBracket = true
			case s.templatePostfix[0] == ')' && len(s.nestedBrackets[parenthesesIndex]) > 0:
				isClosingBracket = true
			case s.templatePostfix[0] == '}' && len(s.nestedBrackets[curlyBracketIndex]) > 0:
				isClosingBracket = true
			}
		}

		if !isClosingBracket {
			return nil, 0
		}
	}

	switch next[0] {
	case '[':
		s.openBracket(squareBracketIndex, prev)
	case ']':
		s.closeBracket(squareBracketIndex, prev)
	case '(':
		s.openBracket(parenthesesIndex, prev)
	case ')':
		s.closeBracket(parenthesesIndex, prev)
	case '{':
		s.openBracket(curlyBracketIndex, prev)
	case '}':
		s.closeBracket(curlyBracketIndex, prev)
	}

	return s, s.Take(1)
}

func (s *codeVisitor) openBracket(bracketIndex BracketIndex, prev []byte) {
	s.nestedBrackets[bracketIndex] = append(s.nestedBrackets[bracketIndex], len(prev))
}

func (s *codeVisitor) closeBracket(bracketIndex BracketIndex, prev []byte) {
	if len(s.nestedBrackets[bracketIndex]) > 0 {
		openIndex := int32(s.nestedBrackets[bracketIndex][len(s.nestedBrackets[bracketIndex])-1])
		closeIndex := int32(len(prev))

		s.f.brackets = append(s.f.brackets, bracketPair{
			OpenedAt: openIndex,
			ClosedAt: closeIndex,
		})

		s.nestedBrackets[bracketIndex] = s.nestedBrackets[bracketIndex][:len(s.nestedBrackets[bracketIndex])-1]
	}
}

func (s *codeVisitor) findBestFactory(next, prev []byte) (types.VisitorFactory, int) {
	var (
		bestFactory   types.VisitorFactory
		bestPrefixLen int
	)

	for _, factory := range s.factories {
		prefixLen := factory.BestPrefixLen(next, prev)
		if prefixLen > bestPrefixLen {
			bestFactory = factory
			bestPrefixLen = prefixLen
		}
	}

	return bestFactory, bestPrefixLen
}
