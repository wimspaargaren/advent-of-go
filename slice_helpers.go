package aoc

func RemoveElementWithCopy(slice []int, index int) []int {
	res := []int{}
	for i, v := range slice {
		if i == index {
			continue
		}
		res = append(res, v)
	}
	return res
}

func CopySlice[T any](slice []T) []T {
	res := make([]T, len(slice))
	copy(res, slice)
	return res
}

func Contains[T comparable](slice []T, val T) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}
