package utils

// 切片中是否包含
func Includes[T comparable](slice []T, value T) bool {
	isRegister := false
	for _, val := range slice {
		if val == value {
			isRegister = true
			break
		}
	}
	return isRegister
}

// 返回一个新的切片
// fn(v, i) 每一项参与计算返回结果
func Map[T comparable](slice []T, fn func(v T, i int) T) []T {
	var newSlice []T
	for i, val := range slice {
		result := fn(val, i)
		newSlice = append(newSlice, result)
	}
	return newSlice
}
