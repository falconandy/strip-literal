package visitor

import (
	"github.com/falconandy/strip-literal/types"
)

type codeVisitor struct {
	*baseVisitor
	f              *codeFactory
	factories      []types.VisitorFactory
	nestedBrackets [][]int
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
	bestFactory, bestPrefix := s.findBestFactory(next, prev)
	if bestFactory != nil {
		return bestFactory.CreateVisitor(bestPrefix), len(bestPrefix)
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

func (s *codeVisitor) findBestFactory(next, prev []byte) (types.VisitorFactory, []byte) {
	var (
		bestFactory types.VisitorFactory
		bestPrefix  []byte
	)

	for _, factory := range s.factories {
		prefix := factory.BestPrefix(next, prev)
		if len(prefix) > len(bestPrefix) {
			bestFactory = factory
			bestPrefix = prefix
		}
	}

	return bestFactory, bestPrefix
}
