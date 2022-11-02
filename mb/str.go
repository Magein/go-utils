package mb

// Len Get string length, compatible with Chinese characters
func Len(str string) int {
	values := []rune(str)
	return len(values)
}

// Sub Intercept a string from the specified location
// Intercepts a string of the specified length from the specified location
// sub(str,2) Intercept from the second to the end
// sub(str,2,4) Cut four from the second
func Sub(str string, offset ...int) string {

	// 兼容字符串中包含汉字
	values := []rune(str)

	// 字符串长度
	strLen := len(values)

	// 截取的起始位置
	start := 0
	// 截取的长度
	length := strLen

	// 参数的长度
	pLen := len(offset)

	if pLen == 1 {
		start = offset[0]
	} else if pLen == 2 {
		start = offset[0]
		length = offset[1]
	}

	if length >= strLen {
		length = strLen
	}

	if start > strLen {
		start = 0
		length = strLen
	}

	// 将字符串先从指定位置截取
	values = values[start : start+length]

	res := ""
	for _, val := range values {
		res += string(val)
	}

	return res
}
