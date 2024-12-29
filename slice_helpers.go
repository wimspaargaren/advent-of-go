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
