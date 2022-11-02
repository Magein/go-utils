package tools

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

// ConvertByte Supported Types include int,int8,int32,int64,float,float32,float64,string
// 1B(byte 字节)=8bit
// 1KB(Kilobyte 千字节)=1024B
// 1MB(Megabyte 兆字节 简称“兆”)=1024KB
// 1GB(Gigabyte 吉字节 又称“千兆”)=1024MB
func ConvertByte(parameter interface{}) string {

	size := ConvertNumeralToFloat(parameter)

	unit := ""
	// byte
	var val, bytes, kilobytes, megabyte, gigabyte float64
	bytes = 1024
	// Kilobyte
	kilobytes = 1048576
	// Megabyte
	megabyte = 1073741824
	// Gigabyte
	gigabyte = 1099511627776

	val = 0
	if size >= 0 && size <= bytes {
		unit = "B"
		val = size
	} else if size > bytes && size < kilobytes {
		unit = "KB"
		val = size / bytes
	} else if size > kilobytes && size < megabyte {
		unit = "MB"
		val = size / kilobytes
	} else if size > gigabyte {
		unit = "GB"
		val = size / megabyte
	}

	return strconv.FormatFloat(math.Ceil(val), 'f', -1, 64) + unit
}

// ConvertNumeral Convert int int8 int32 int64 float float32 float64 string to a string
func ConvertNumeral(number interface{}) string {
	value := ""
	switch number.(type) {
	case int8:
		if v, ok := number.(int8); ok {
			if vint := int(v); vint != 0 {
				value = strconv.Itoa(vint)
			}
		}
		break
	case int16:
		if v, ok := number.(int16); ok {
			if vint := int(v); vint != 0 {
				value = strconv.Itoa(vint)
			}
		}
		break
	case int32:
		if v, ok := number.(int32); ok {
			if vint := int(v); vint != 0 {
				value = strconv.Itoa(vint)
			}
		}
		break
	case int64:
		if v, ok := number.(int64); ok {
			if vint := int(v); vint != 0 {
				value = strconv.Itoa(vint)
			}
		}
		break
	case int:
		if v, ok := number.(int); ok && v != 0 {
			value = strconv.Itoa(v)
		}
		break
	case float32:
		if v, ok := number.(float32); ok {
			value = strconv.FormatFloat(float64(v), 'f', 2, 64)
		}
		break
	case float64:
		value = strconv.FormatFloat(number.(float64), 'f', 2, 64)
		break
	case string:
		value = number.(string)
		if v, _ := regexp.MatchString("^0+\\.0+$", value); v {
			value = ""
		} else {
			if vv, _ := regexp.MatchString("^[1-9]?[0-9]+\\.?[0-9]+$", value); !vv {
				value = ""
			}
		}
		break
	}
	return value
}

// ConvertNumeralToFloat Convert  type to float
// Supported types int float string
func ConvertNumeralToFloat(number interface{}) float64 {

	numeral := ConvertNumeral(number)

	if v, e := strconv.ParseFloat(numeral, 64); e == nil {
		return v
	}

	return 0
}

// ConvertNumeralToInt Convert  type to integer
// Supported types int float string
func ConvertNumeralToInt(number interface{}) int {

	numeral := ConvertNumeral(number)

	index := strings.Index(numeral, ".")

	if index != -1 {
		numeral = numeral[0:index]
	}

	if v, e := strconv.Atoi(numeral); e == nil {
		return v
	}

	return 0
}
