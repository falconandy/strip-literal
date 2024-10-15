package strip

import (
	"unicode/utf8"

	"github.com/falconandy/strip-literal/lang"
	"github.com/falconandy/strip-literal/parser"
	"github.com/falconandy/strip-literal/types"
)

type Language string

const (
	LangGo         Language = "go"
	LangJavaScript Language = "javascript"
	LangTypeScript Language = "typescript"
	LangJava       Language = "java"
	LangKotlin     Language = "kotlin"
	LangPython     Language = "python"
	LangCSharp     Language = "csharp"
	LangHTML       Language = "html"
	LangCSS        Language = "css"
)

type Mode uint8

const (
	None = iota
	Remove
	ByteToSpace
	RuneToSpace
)

type Options struct {
	Comments Mode
	Strings  Mode
}

func StripLiterals(code []byte, language Language, options Options) int32 {
	var codeFactory types.CodeFactory

	switch language {
	case LangGo:
		codeFactory = lang.NewGoFactory()
	case LangJavaScript, LangTypeScript:
		codeFactory = lang.NewJavaScriptFactory()
	case LangJava:
		codeFactory = lang.NewJavaFactory()
	case LangKotlin:
		codeFactory = lang.NewKotlinFactory()
	case LangPython:
		codeFactory = lang.NewPythonFactory()
	case LangCSharp:
		codeFactory = lang.NewCSharpFactory()
	case LangHTML:
		codeFactory = lang.NewHTMLFactory()
	case LangCSS:
		codeFactory = lang.NewCSSFactory()
	default:
		return int32(len(code))
	}

	if codeFactory == nil || (options.Strings == None && options.Comments == None) {
		return int32(len(code))
	}

	return stripLiterals(code, codeFactory, options)
}

func stripLiterals(code []byte, codeFactory types.CodeFactory, options Options) int32 {
	var processed int32

	segments := parser.ParseBytes(codeFactory, code)
	for _, segment := range segments {
		switch {
		case options.Comments != None && segment.IsComment():
			processed = moveBytes(code, processed, segment.Position, segment.Length, options.Comments)
		case options.Strings != None && segment.IsString():
			processed = copyBytes(code, processed, segment.Position, segment.PrefixLength)
			processed = moveBytes(code, processed,
				segment.Position+segment.PrefixLength, segment.Length-segment.PostfixLength-segment.PrefixLength,
				options.Strings)
			processed = copyBytes(code, processed, segment.Position+segment.Length-segment.PostfixLength, segment.PostfixLength)
		case options.Strings != None && segment.IsRegexp():
			processed = copyBytes(code, processed, segment.Position, segment.PrefixLength)
			prevProcessed := processed

			processed = moveBytes(code, processed,
				segment.Position+segment.PrefixLength, segment.Length-segment.PostfixLength-segment.PrefixLength,
				options.Strings)
			if prevProcessed == processed && segment.Length-segment.PrefixLength-segment.PostfixLength > 0 {
				code[processed] = ' '
				processed++
			}

			processed = copyBytes(code, processed, segment.Position+segment.Length-segment.PostfixLength, segment.PostfixLength)
		default:
			processed = copyBytes(code, processed, segment.Position, segment.Length)
		}
	}

	return processed
}

func copyBytes(code []byte, processed, from, count int32) int32 {
	copy(code[processed:], code[from:from+count])

	return processed + count
}

func moveBytes(code []byte, processed, from, count int32, stripMode Mode) int32 {
	position, till := from, from+count

	for position < till {
		size := 1

		if code[position] == '\n' || code[position] == '\r' {
			code[processed] = code[position]
			processed++
		} else if stripMode != Remove {
			if stripMode == RuneToSpace {
				_, size = utf8.DecodeRune(code[position:])
			}

			code[processed] = ' '
			processed++
		}

		position += int32(size)
	}

	return processed
}
