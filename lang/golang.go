package lang

import (
	"github.com/falconandy/strip-literal/types"
	"github.com/falconandy/strip-literal/visitor"
)

func NewGoFactory() types.CodeFactory {
	literalStringFactory := visitor.NewStringFactory(
		types.NewSingleLineString(`'`).WithPostfix(`'`).WithSkip(`\'`, `\\`),
		types.NewSingleLineString(`"`).WithPostfix(`"`).WithSkip(`\"`, `\\`),
		types.NewMultiLineString("`").WithPostfix("`"),
	)
	commentMultiLineFactory := visitor.NewMultiLineCommentFactory("/*", "*/", false)
	commentSingleLineFactory := visitor.NewSingleLineCommentFactory("//")

	return visitor.NewCodeFactory(
		literalStringFactory,
		commentMultiLineFactory,
		commentSingleLineFactory,
	)
}
