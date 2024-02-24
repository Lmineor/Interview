package array

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0{
		return []int{}
	}
	n := len(matrix[0])
	answer := make([]int, n*m)
	left, right := 0, n-1
	top, bottom := 0, m-1
	p := 0
	for p < n*m {
		for j := left; j <= right && p < n*m; j++ {
			answer[p] = matrix[top][j]
			p++
		}
		top++
		for j := top; j <= bottom && p < n*m; j++ {
			answer[p] = matrix[j][right]
			p++
		}
		right--
		for j := right; j >= left && p < n*m; j-- {
			answer[p] = matrix[bottom][j]
			p++
		}
		bottom--
		for j := bottom; j >= top && p < n*m; j-- {
			answer[p] = matrix[j][left]
			p++
		}
		left++
	}
	return answer
}
