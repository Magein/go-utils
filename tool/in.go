package tools

import (
	"strings"
)

// IndexOfInt Get array index of integer value
func IndexOfInt(arr []int, value int) int {
	for key, val := range arr {
		if val == value {
			return key
		}
	}
	return -1
}

// IndexOfString Get array index of string value
func IndexOfString(arr []string, value string) int {
	for key, val := range arr {
		if val == value {
			return key
		}
	}
	return -1
}

// IndexOf Get the array index of the value
// Return the key for needle if it is found in the array,-1 otherwise
func IndexOf(arr interface{}, value interface{}) int {

	s1, s1ok := arr.([]string)
	s2, s2ok := value.(string)

	if s1ok && s2ok {
		return IndexOfString(s1, s2)
	}

	i1, i1ok := arr.([]int)
	i2, i2ok := value.(int)
	if i1ok && i2ok {
		return IndexOfInt(i1, i2)
	}

	return -1
}

func InArray(haystack interface{}, needle interface{}) bool {

	if str, ok := haystack.(string); ok {
		if IndexOf(strings.Split(str, ","), needle) == -1 {
			return false
		}
	} else {
		if IndexOf(haystack, needle) == -1 {
			return false
		}
	}

	return true
}
