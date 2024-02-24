package array

//给定一个正整数 n，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。
//
//示例:
//
//输入: 3 输出: [ [ 1, 2, 3 ], [ 8, 9, 4 ], [ 7, 6, 5 ] ]
//
//

func generateMatrix(n int) [][]int {
	// 依次循环到里面
	left, right := 0, n-1
	top, bottom := 0, n-1
	answer := make([][]int, n, n)
	for i := 0; i < n; i++ {
		answer[i] = make([]int, n, n)
	}
	nn := n * n
	num := 1
	for num <= nn {
		for i:= left; i<=right; i++{
			answer[top][i] = num
			num++
		}
		top++
		for i:= top; i<=bottom; i++{
			answer[i][right] = num
			num++
		}
		right--
		for i:=right; i>= left; i--{
			answer[bottom][i]=num
			num++
		}
		bottom--
		for i:= bottom; i>= top; i--{
			answer[i][left]=num
			num++
		}
		left++
	}
	return answer
}
