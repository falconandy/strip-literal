package parser

import (
	"github.com/falconandy/strip-literal/types"
)

func Parse(codeFactory types.CodeFactory, source string) []types.Segment {
	return ParseBytes(codeFactory, []byte(source))
}

func ParseBytes(codeFactory types.CodeFactory, source []byte) []types.Segment {
	var segments []types.Segment

	visitors := []types.SegmentVisitor{codeFactory.CreateVisitor(nil)}

	next := source
	prev := source[:0]

	for len(next) > 0 {
		var lastSegment types.Segment

		currentVisitor := visitors[len(visitors)-1]

		if len(segments) > 0 {
			lastSegment = segments[len(segments)-1]
		}

		nextVisitor, used := currentVisitor.Visit(next, prev)
		if nextVisitor == nil {
			length, prefixLength, postfixLength := currentVisitor.PopVisited()
			if length > 0 {
				segments = append(segments, types.Segment{
					Type:          currentVisitor.SegmentType(),
					Position:      lastSegment.Position + lastSegment.Length,
					Length:        int32(length),
					PrefixLength:  int32(prefixLength),
					PostfixLength: int32(postfixLength),
				})
			}

			visitors = visitors[:len(visitors)-1]
		} else if nextVisitor != currentVisitor {
			length, prefixLength, postfixLength := currentVisitor.PopVisited()
			if length > 0 {
				segments = append(segments, types.Segment{
					Type:          currentVisitor.SegmentType(),
					Position:      lastSegment.Position + lastSegment.Length,
					Length:        int32(length),
					PrefixLength:  int32(prefixLength),
					PostfixLength: int32(postfixLength),
				})
			}

			visitors = append(visitors, nextVisitor)
		}

		next = next[used:]
		prev = prev[:len(prev)+used]
	}

	for visitorIndex := len(visitors) - 1; visitorIndex >= 0; visitorIndex-- {
		var lastSegment types.Segment
		if len(segments) > 0 {
			lastSegment = segments[len(segments)-1]
		}

		length, prefixLength, postfixLength := visitors[visitorIndex].PopVisited()
		if length > 0 {
			segments = append(segments, types.Segment{
				Type:          visitors[visitorIndex].SegmentType(),
				Position:      lastSegment.Position + lastSegment.Length,
				Length:        int32(length),
				PrefixLength:  int32(prefixLength),
				PostfixLength: int32(postfixLength),
			})
		}
	}

	return segments
}
