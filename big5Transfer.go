package nhicard

import (
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

func Big5ToUTF8(src []byte) string {
	decoder := traditionalchinese.Big5.NewDecoder()
	result, _, _ := transform.String(decoder, string(src))
	return result
}
