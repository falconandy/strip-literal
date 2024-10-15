package lang

import (
	"github.com/falconandy/strip-literal/types"
	"github.com/falconandy/strip-literal/visitor"
)

func NewCPPFactory() types.CodeFactory {
	literalStringFactory := visitor.NewStringFactory(
		types.NewSingleLineString("'", "L'", "u8'", "u'", "U'").WithPostfix("'").WithSkip(`\'`, `\\`),
		types.NewSingleLineString(`"`, `L"`, `u8"`, `u"`, `U"`).WithPostfix(`"`).WithSkip(`\"`, `\\`),
	)
	rawStringFactory := visitor.NewCPPRawStringFactory()
	commentMultiLineFactory := visitor.NewMultiLineCommentFactory("/*", "*/", false)
	commentSingleLineFactory := visitor.NewSingleLineCommentFactory("//")

	return visitor.NewCodeFactory(
		literalStringFactory,
		rawStringFactory,
		commentMultiLineFactory,
		commentSingleLineFactory,
	)
}
