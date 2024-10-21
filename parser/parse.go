package parser

import (
	"github.com/falconandy/strip-literal/types"
)

func Parse(codeFactory types.CodeFactory, source string) []types.Segment {
	return ParseBytes(codeFactory, []byte(source))
}

func ParseBytes(codeFactory types.CodeFactory, source []byte) []types.Segment {
	var segments []types.Segment

	visitors := []types.SegmentVisitor{codeFactory.CreateVisitor(nil, nil)}
	results := []types.VisitResult{{}}
	var position int

	for position < len(source) {
		next := source[position:]
		prev := source[:position]

		currentVisitor := visitors[len(visitors)-1]

		nextVisitor, result := currentVisitor.Visit(next, prev)
		position += result.PrefixLength + result.InnerLength + result.PostfixLength

		if nextVisitor == nil {
			currentResult := results[len(visitors)-1].Add(result)
			if currentResult.HasData() {
				segments = append(segments, types.Segment{
					Type:          currentVisitor.SegmentType(),
					PrefixLength:  int32(currentResult.PrefixLength),
					InnerLength:   int32(currentResult.InnerLength),
					PostfixLength: int32(currentResult.PostfixLength),
				})
			}

			visitors = visitors[:len(visitors)-1]
			results = results[:len(results)-1]
		} else if nextVisitor != currentVisitor {
			currentResult := results[len(visitors)-1]
			if result.PrefixLength == 0 {
				currentResult = currentResult.Add(result)
				result = types.VisitResult{}
			}

			if currentResult.HasData() {
				segments = append(segments, types.Segment{
					Type:          currentVisitor.SegmentType(),
					PrefixLength:  int32(currentResult.PrefixLength),
					InnerLength:   int32(currentResult.InnerLength),
					PostfixLength: int32(currentResult.PostfixLength),
				})
			}

			results[len(visitors)-1] = types.VisitResult{}
			visitors = append(visitors, nextVisitor)
			results = append(results, result)
		} else {
			results[len(visitors)-1] = results[len(visitors)-1].Add(result)
		}
	}

	for visitorIndex := len(visitors) - 1; visitorIndex >= 0; visitorIndex-- {
		currentResult := results[visitorIndex]

		if currentResult.HasData() {
			segments = append(segments, types.Segment{
				Type:          visitors[visitorIndex].SegmentType(),
				PrefixLength:  int32(currentResult.PrefixLength),
				InnerLength:   int32(currentResult.InnerLength),
				PostfixLength: int32(currentResult.PostfixLength),
			})
		}
	}

	return segments
}
