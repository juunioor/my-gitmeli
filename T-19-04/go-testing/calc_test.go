package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	num1 := 3
	num2 := 5
	expectedResult := 8

	result := Sum(num1, num2)

	assert.Equal(t, expectedResult, result, "devem ser iguais")
}

func TestSub(t *testing.T) {
	num1 := 3
	num2 := 2
	expectedResult := 1

	result := Sub(num1, num2)

	assert.Equal(t, expectedResult, result, "Devem ser iguais")
}

func TestDivide(t *testing.T) {
	num := 3
	den := 1
	expectedResult := 3

	result, err := Divide(num, den)

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result, "Devem ser iguais.")
}

func TestOrder(t *testing.T) {
	array := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	expectedResult := []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}

	result := Order(array)

	assert.Equal(t, expectedResult, result, "Array desordenado.")
}
