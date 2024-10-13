//go:build js

package main

import (
	"unsafe"

	"github.com/falconandy/strip-literal"
)

var languages = []strip.Language{
	strip.LangGo,
	strip.LangJavaScript,
	strip.LangTypeScript,
	strip.LangJava,
	strip.LangKotlin,
	strip.LangPython,
}

//export stripLiterals
func _stripLiterals(codePtr, codeSize, languageIndex, stripModes uint32) (count uint32) {
	if languageIndex >= uint32(len(languages)) {
		return codeSize
	}

	commentsMode := stripModes % 4
	stringsMode := (stripModes / 4) % 4

	code := ptrToBytes(codePtr, codeSize)
	count = uint32(stripLiterals(code, languages[languageIndex], strip.Mode(commentsMode), strip.Mode(stringsMode)))
	return count
}

func stripLiterals(code []byte, language strip.Language, commentsMode, stringsMode strip.Mode) int32 {
	return strip.StripLiterals(code, language, strip.Options{
		Comments: commentsMode,
		Strings:  stringsMode,
	})
}

func ptrToBytes(ptr uint32, size uint32) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(uintptr(ptr))), size)
}

func main() {}
