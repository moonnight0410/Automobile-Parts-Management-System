package utils

import (
	"strconv"
)

// ParseInt 解析字符串为整数
func ParseInt(s string, result *int) error {
	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*result = i
	return nil
}
