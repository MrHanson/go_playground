// 给定一个string, 请返回反转后的string

package logic

func ReverseStr(s string) string {
	l := len(s)
	str := []byte(s)
	for i := 0; i < l; i++ {
		str[i] = s[l-1-i]
	}

	return string(str)
}
