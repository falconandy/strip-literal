package visitor

import (
	"bytes"
)

func FindSubData(data, prefix, inner, postfix []byte) int {
	position := 0
	for position <= len(data)-len(prefix)-len(inner)-len(postfix) {
		prefixIndex := bytes.Index(data[position:], prefix)
		if prefixIndex < 0 {
			return -1
		}

		if bytes.HasPrefix(data[position+prefixIndex+len(prefix):], inner) && bytes.HasPrefix(data[position+prefixIndex+len(prefix)+len(inner):], postfix) {
			return position + prefixIndex
		}

		position += prefixIndex + 1
	}

	return -1
}
