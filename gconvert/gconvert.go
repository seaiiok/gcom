package gconvert

import "strings"

//StringPrefixBom 去掉带有BOM头标识UTF-8编码
func StringPrefixBom(v string) string {
	return strings.TrimPrefix(v, string([]byte{239, 187, 191}))
}
