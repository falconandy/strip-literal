package lang

import (
	"github.com/falconandy/strip-literal/types"
	"github.com/falconandy/strip-literal/visitor"
)

func NewCFactory() types.CodeFactory {
	literalStringFactory := visitor.NewStringFactory(
		types.NewSingleLineString("'").WithPostfix("'").WithSkip(`\'`, `\\`),
		types.NewSingleLineString(`"`).WithPostfix(`"`).WithSkip(`\"`, `\\`, "\\\n\r", "\\\r\n", "\\\n", "\\\r"),
	)
	commentMultiLineFactory := visitor.NewMultiLineCommentFactory("/*", "*/", false)
	commentSingleLineFactory := visitor.NewSingleLineCommentFactory("//")

	return visitor.NewCodeFactory(
		literalStringFactory,
		commentMultiLineFactory,
		commentSingleLineFactory,
	)
}
