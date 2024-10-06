package types

type StringDefinition struct {
	Prefixes  [][]byte
	Postfix   []byte
	Skip      [][]byte
	Multiline bool
}

func NewSingleLineString(prefix string, prefixes ...string) StringDefinition {
	return newStringDefinition(prefix, prefixes, false)
}

func NewMultiLineString(prefix string, prefixes ...string) StringDefinition {
	return newStringDefinition(prefix, prefixes, true)
}

func newStringDefinition(prefix string, prefixes []string, multiline bool) StringDefinition {
	bytePrefixes := make([][]byte, 1+len(prefixes))
	bytePrefixes[0] = []byte(prefix)

	for i, p := range prefixes {
		bytePrefixes[i+1] = []byte(p)
	}

	return StringDefinition{
		Prefixes:  bytePrefixes,
		Multiline: multiline,
	}
}

func (d StringDefinition) WithPostfix(postfix string) StringDefinition {
	d.Postfix = []byte(postfix)

	return d
}

func (d StringDefinition) WithSkip(skip ...string) StringDefinition {
	if len(skip) == 1 {
		d.Skip = [][]byte{[]byte(skip[0])}

		return d
	}

	byteSkips := make([][]byte, len(skip))
	for i, s := range skip {
		byteSkips[i] = []byte(s)
	}

	d.Skip = byteSkips

	return d
}
