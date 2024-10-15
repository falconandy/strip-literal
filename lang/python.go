package lang

import (
	"github.com/falconandy/strip-literal/types"
	"github.com/falconandy/strip-literal/visitor"
)

func NewPythonFactory() types.CodeFactory {
	literalStringFactory := visitor.NewStringFactory(
		types.NewSingleLineString(`"`, `u"`, `U"`, `b"`, `B"`).
			WithPostfix(`"`).WithSkip(`\"`, `\\`, "\\\n\r", "\\\r\n", "\\\n", "\\\r"),
		types.NewSingleLineString(`'`, `u'`, `U'`, `b'`, `B'`).
			WithPostfix(`'`).WithSkip(`\'`, `\\`, "\\\n\r", "\\\r\n", "\\\n", "\\\r"),
		types.NewMultiLineString(`"""`, `u"""`, `U"""`, `b"""`, `B"""`).
			WithPostfix(`"""`),
		types.NewMultiLineString(`'''`, `u'''`, `U'''`, `b'''`, `B'''`).
			WithPostfix(`'''`),
		types.NewSingleLineString(`r"`, `R"`, `br"`, `bR"`, `Br"`, `BR"`, `rb"`, `rB"`, `Rb"`, `RB"`).
			WithPostfix(`"`),
		types.NewSingleLineString(`r'`, `R'`, `br'`, `bR'`, `Br'`, `BR'`, `rb'`, `rB'`, `Rb'`, `RB'`).
			WithPostfix(`'`),
		types.NewMultiLineString(`r"""`, `R"""`, `br"""`, `bR"""`, `Br"""`, `BR"""`, `rb"""`, `rB"""`, `Rb"""`, `RB"""`).
			WithPostfix(`"""`),
		types.NewMultiLineString(`r'''`, `R'''`, `br'''`, `bR'''`, `Br'''`, `BR'''`, `rb'''`, `rB'''`, `Rb'''`, `RB'''`).
			WithPostfix(`'''`),
		types.NewSingleLineString(`f"`, `F"`).
			WithPostfix(`"`).WithSkip(`\"`, `\\`, "\\\n\r", "\\\r\n", "\\\n", "\\\r").
			WithTemplate("{", "}"),
		types.NewSingleLineString(`f'`, `F'`).
			WithPostfix(`'`).WithSkip(`\'`, `\\`, "\\\n\r", "\\\r\n", "\\\n", "\\\r").
			WithTemplate("{", "}"),
		types.NewMultiLineString(`f"""`, `F"""`).
			WithPostfix(`"""`).
			WithTemplate("{", "}"),
		types.NewMultiLineString(`f'''`, `F'''`).
			WithPostfix(`'''`).
			WithTemplate("{", "}"),
		types.NewSingleLineString(`fr"`, `fR"`, `Fr"`, `FR"`, `rf"`, `rF"`, `Rf"`, `RF"`).
			WithPostfix(`"`).
			WithTemplate("{", "}"),
		types.NewSingleLineString(`fr'`, `fR'`, `Fr'`, `FR'`, `rf'`, `rF'`, `Rf'`, `RF'`).
			WithPostfix(`'`).
			WithTemplate("{", "}"),
		types.NewMultiLineString(`fr"""`, `fR"""`, `Fr"""`, `FR"""`, `rf"""`, `rF"""`, `Rf"""`, `RF"""`).
			WithPostfix(`"""`).
			WithTemplate("{", "}"),
		types.NewMultiLineString(`fr'''`, `fR'''`, `Fr'''`, `FR'''`, `rf'''`, `rF'''`, `Rf'''`, `RF'''`).
			WithPostfix(`'''`).
			WithTemplate("{", "}"),
	)
	commentSingleLineFactory := visitor.NewSingleLineCommentFactory("#")

	return visitor.NewCodeFactory(
		literalStringFactory,
		commentSingleLineFactory,
	)
}
