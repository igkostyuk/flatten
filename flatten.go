package flatten

import (
	"errors"
)

var ErrInvalidSquareMatrix = errors.New("invalid square matrix")

func Clockwise(matrix [][]int) ([]int, error) {
	l := len(matrix)
	if l == 0 {
		return []int{}, nil
	}
	if !isSquareMatrix(matrix) {
		return nil, ErrInvalidSquareMatrix
	}
	top, left, bottom, right := 0, 0, l-1, l-1
	result := make([]int, 0, l*l)

	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--
		for i := right; i >= left; i-- {
			result = append(result, matrix[bottom][i])
		}
		bottom--
		for i := bottom; i >= top; i-- {
			result = append(result, matrix[i][left])
		}
		left++
	}

	return result, nil
}

func isSquareMatrix(m [][]int) bool {
	l := len(m)
	for _, f := range m {
		if len(f) != l {
			return false
		}
	}

	return true
}
