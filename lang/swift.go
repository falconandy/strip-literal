package lang

import (
	"github.com/falconandy/strip-literal/types"
	"github.com/falconandy/strip-literal/visitor"
)

func NewSwiftFactory() types.CodeFactory {
	literalStringFactory := visitor.NewSwiftStringFactory()
	literalRegexpFactory := visitor.NewRegexpFactory()
	commentMultiLineFactory := visitor.NewMultiLineCommentFactory("/*", "*/", true)
	commentSingleLineFactory := visitor.NewSingleLineCommentFactory("//")

	return visitor.NewCodeFactory(
		literalStringFactory,
		literalRegexpFactory,
		commentMultiLineFactory,
		commentSingleLineFactory,
	)
}
