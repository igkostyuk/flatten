package flatten_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/igkostyuk/flatten"
)

// nolint: funlen
func TestClockwise(t *testing.T) {
	t.Run("flatten matrix", func(t *testing.T) {
		tt := []struct {
			input [][]int
			want  []int
		}{
			{
				input: [][]int{},
				want:  []int{},
			},
			{
				input: [][]int{
					{1, 2, 3},
					{8, 9, 4},
					{7, 6, 5},
				},
				want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			{
				input: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				want: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
			},
			{
				input: [][]int{
					{1, 2, 3, 1},
					{4, 5, 6, 4},
					{7, 8, 9, 7},
					{7, 8, 9, 7},
				},
				want: []int{1, 2, 3, 1, 4, 7, 7, 9, 8, 7, 7, 4, 5, 6, 9, 8},
			},
		}
		for _, test := range tt {
			got, err := flatten.Clockwise(test.input)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v want %v", got, test.want)
			}
			if err != nil {
				t.Errorf("didn't expect error")
			}
		}
	})
	t.Run("error cases", func(t *testing.T) {
		tt := []struct {
			input [][]int
			err   error
		}{
			{
				input: [][]int{
					{1, 2, 3},
					{8, 9, 4},
				},
				err: flatten.ErrInvalidSquareMatrix,
			},
			{
				input: [][]int{
					{1, 2},
					{8, 9},
					{7, 6},
				},
				err: flatten.ErrInvalidSquareMatrix,
			},
		}
		for _, test := range tt {
			_, err := flatten.Clockwise(test.input)
			if err == nil {
				t.Error("wanted an error but didn't get one")
			}
			if !errors.Is(err, test.err) {
				t.Error("got unexpected error")
			}
		}
	})
}

func BenchmarkClockwise(b *testing.B) {
	a := [][]int{
		{1, 2, 3},
		{8, 9, 4},
		{7, 6, 5},
	}
	for i := 0; i < b.N; i++ {
		flatten.Clockwise(a) // nolint: errcheck
	}
}
