package calc

import "errors"

func Sum(num1, num2 int) int {
	return num1 + num2
}

func Sub(num1, num2 int) int {
	return num1 - num2
}

func Divide(num, den int) (int, error) {
	if den == 0 {
		return 0, errors.New("erro: O denominador não pode ser 0")
	}

	return num / den, nil
}

func Order(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}
