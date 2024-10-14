package lang

import (
	"github.com/falconandy/strip-literal/types"
	"github.com/falconandy/strip-literal/visitor"
)

func NewCSharpFactory() types.CodeFactory {
	literalStringFactory := visitor.NewStringFactory(
		types.NewSingleLineString("'").WithPostfix("'").WithSkip(`\'`, `\\`),
		types.NewSingleLineString(`"`).WithPostfix(`"`).WithSkip(`\"`, `\\`),
		types.NewMultiLineString(`@"`).WithPostfix(`"`).WithSkip(`""`),
		types.NewSingleLineString(`$"`).WithPostfix(`"`).WithSkip(`\"`, `\\`, "{{").WithTemplate("{", "}"),
		types.NewMultiLineString(`$@"`, `@$"`).WithPostfix(`"`).WithSkip(`""`, "{{").WithTemplate("{", "}"),
	)
	commentMultiLineFactory := visitor.NewMultiLineCommentFactory("/*", "*/", false)
	commentSingleLineFactory := visitor.NewSingleLineCommentFactory("//")

	return visitor.NewCodeFactory(
		literalStringFactory,
		commentMultiLineFactory,
		commentSingleLineFactory,
	)
}
