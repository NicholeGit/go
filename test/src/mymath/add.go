package mymath

import "errors"

func Add(a int, b int) (ret int, err error) {
	if a < 0 || b < 0 { // 假设这个函数只支持两个非负数字的加法
		err = errors.New("Should be non-negative numbers!")
		return
	}
	return a + b, nil // 支持多重返回值
}
