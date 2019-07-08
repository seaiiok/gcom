package gconvert

import "strings"

type I struct{}

func (g *I) StringPrefixBom(v string) string {
	return strings.TrimPrefix(v, string([]byte{239, 187, 191}))
}
