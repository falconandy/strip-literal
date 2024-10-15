package lang

import (
	"github.com/falconandy/strip-literal/types"
	"github.com/falconandy/strip-literal/visitor"
)

func NewHTMLFactory() types.CodeFactory {
	commentMultiLineFactory := visitor.NewMultiLineCommentFactory("<!--", "-->", false)

	return visitor.NewCodeFactory(
		commentMultiLineFactory,
	)
}
