package lang

import (
	"github.com/falconandy/strip-literal/types"
	"github.com/falconandy/strip-literal/visitor"
)

func NewKotlinFactory() types.CodeFactory {
	literalStringFactory := visitor.NewStringFactory(
		types.NewSingleLineString(`'`).WithPostfix(`'`).WithSkip(`\'`, `\\`),
		types.NewSingleLineString(`"`).WithPostfix(`"`).WithSkip(`\"`, `\\`).WithTemplate("${", "}"),
		types.NewMultiLineString(`"""`).WithPostfix(`"""`).WithTemplate("${", "}"),
	)
	commentMultiLineFactory := visitor.NewMultiLineCommentFactory("/*", "*/", true)
	commentSingleLineFactory := visitor.NewSingleLineCommentFactory("//")

	return visitor.NewCodeFactory(
		literalStringFactory,
		commentMultiLineFactory,
		commentSingleLineFactory,
	)
}
