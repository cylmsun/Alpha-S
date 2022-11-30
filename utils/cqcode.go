package utils

import (
	"fmt"
	"strings"
)

// GenCQCode 生成cq码
// 1:  2:图片
func GenCQCode(msg string, sType int8) (cqCode string) {
	fmtStr := "[cq:%s,%s]"
	switch sType {
	case 2:
		tempS := fmt.Sprintf("file=%s", msg)
		cqCode = fmt.Sprintf(fmtStr, "image", tempS)
	default:
		cqCode = msg
	}
	return
}

// 特殊字符转义
func escapeCQText(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "[", "&#91;")
	s = strings.ReplaceAll(s, "]", "&#93;")
	return s
}
