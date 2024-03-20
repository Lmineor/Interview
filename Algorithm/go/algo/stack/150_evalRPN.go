package stack

const (
	plus  = "+"
	minus = "-"
	cheng = "*"
	chu   = "/"
)

func evalRPN(tokens []string) int {
	i := 0
	for len(tokens) > 1 {
		if !isNum(tokens[i]){
			leftNum := tokens[i-2]
			rightNum := tokens[i-1]
			op = tokens[i]
		}
	}
}

func cal(leftNum, rightNum, op string)int{
	
}

func isNum(c string) bool {
	nums := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	for _, num := range nums {
		if c == num {
			return true
		}
	}
	return false
}
