package lang

import (
	"github.com/falconandy/strip-literal/types"
	"github.com/falconandy/strip-literal/visitor"
)

func NewJavaScriptFactory() types.CodeFactory {
	literalStringFactory := visitor.NewStringFactory(
		types.NewSingleLineString(`"`).WithPostfix(`"`).WithSkip(`\"`, `\\`, "\\\n\r", "\\\r\n", "\\\n", "\\\r"),
		types.NewSingleLineString("'").WithPostfix("'").WithSkip(`\'`, `\\`, "\\\n\r", "\\\r\n", "\\\n", "\\\r"),
		types.NewMultiLineString("`").WithPostfix("`").WithSkip("\\`", `\u{`, `\$`).WithTemplate("${", "}"),
	)
	literalRegexpFactory := visitor.NewRegexpFactory()
	commentMultiLineFactory := visitor.NewMultiLineCommentFactory("/*", "*/", false)
	commentSingleLineFactory := visitor.NewSingleLineCommentFactory("//")

	return visitor.NewCodeFactory(
		literalStringFactory,
		literalRegexpFactory,
		commentMultiLineFactory,
		commentSingleLineFactory,
	)
}
