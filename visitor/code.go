package visitor

import (
	"bytes"

	"github.com/falconandy/strip-literal/types"
)

type codeVisitor struct {
	f               *codeFactory
	factories       []types.VisitorFactory
	templatePostfix []byte
	nestedBrackets  [3]int
}

type BracketIndex int

const (
	squareBracketIndex BracketIndex = 0
	parenthesesIndex   BracketIndex = 1
	curlyBracketIndex  BracketIndex = 2
)

func (s *codeVisitor) Visit(next, prev []byte) (types.SegmentVisitor, types.VisitResult) {
	bestFactory, bestPrefixLen := s.findBestFactory(next, prev)
	if bestFactory != nil {
		return bestFactory.CreateVisitor(next[:bestPrefixLen], next[bestPrefixLen:]), types.VisitResult{PrefixLength: bestPrefixLen}
	}

	if len(s.templatePostfix) > 0 && bytes.HasPrefix(next, s.templatePostfix) {
		isClosingBracket := false

		if len(s.templatePostfix) == 1 &&
			(s.templatePostfix[0] == ']' || s.templatePostfix[0] == ')' || s.templatePostfix[0] == '}') {
			switch {
			case s.templatePostfix[0] == ']' && s.nestedBrackets[squareBracketIndex] > 0:
				isClosingBracket = true
			case s.templatePostfix[0] == ')' && s.nestedBrackets[parenthesesIndex] > 0:
				isClosingBracket = true
			case s.templatePostfix[0] == '}' && s.nestedBrackets[curlyBracketIndex] > 0:
				isClosingBracket = true
			}
		}

		if !isClosingBracket {
			return nil, types.VisitResult{}
		}
	}

	switch next[0] {
	case '[':
		s.openBracket(squareBracketIndex)
	case ']':
		s.closeBracket(squareBracketIndex)
	case '(':
		s.openBracket(parenthesesIndex)
	case ')':
		s.closeBracket(parenthesesIndex)
	case '{':
		s.openBracket(curlyBracketIndex)
	case '}':
		s.closeBracket(curlyBracketIndex)
	}

	return s, types.VisitResult{InnerLength: 1}
}

func (s *codeVisitor) SegmentType() types.SegmentType {
	return types.SegmentTypeCode
}

func (s *codeVisitor) openBracket(bracketIndex BracketIndex) {
	s.nestedBrackets[bracketIndex]++
}

func (s *codeVisitor) closeBracket(bracketIndex BracketIndex) {
	if s.nestedBrackets[bracketIndex] > 0 {
		s.nestedBrackets[bracketIndex]--
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
