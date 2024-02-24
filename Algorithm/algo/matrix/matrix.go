package matrix

import "fmt"

func printMatrix(m [][]int) {
	for _, row := range m {
		for _, col := range row {
			fmt.Printf("%d,", col)
		}
		fmt.Printf("\n")
	}
}
func InClockRotate(m [][]int) [][]int {
	d := len(m)
	for _, row := range m {
		for i := 0; i < d/2; i++ {
			tmp := row[i]
			row[i] = row[d-i-1]
			row[d-i-1] = tmp
		}
	}
	for row := 0; row < d; row++ {
		for col := 0; col < d-row-1; col++ {
			m[row][col], m[d-col-1][d-row-1] = m[d-col-1][d-row-1], m[row][col]
		}
	}
	printMatrix(m)
	return m
}

func Entry() {
	var m [][]int
	m = append(m, []int{1, 2}, []int{3, 4})
	//m = append(m, []int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9})
	//m = append(m, []int{1, 2, 3, 4}, []int{5, 6, 7, 8}, []int{9, 10, 11, 12}, []int{13, 14, 15, 16})
	InClockRotate(m)
}
