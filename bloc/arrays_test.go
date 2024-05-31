package bloc

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestMultiplyBy100(t *testing.T) {
	arrays := [...]int32{1, 2, 3, 4, 5, 6, 7}
	MultiplyBy100(arrays)
	assert.Equal(t, arrays, [...]int32{1, 2, 3, 4, 5, 6, 7})
}

func TestMultiplyBy100_(t *testing.T) {
	arrays := [...]int32{1, 2, 3, 4, 5, 6, 7}
	MultiplyBy100_(&arrays)
	assert.Equal(t, arrays, [...]int32{100, 200, 300, 400, 500, 600, 700})
}

func TestMatrixMultiply(t *testing.T) {
	m1 := [][]int{{1, 2}, {3, 4}}
	m2 := [][]int{{2}, {3}}
	assert.Equal(t, MatrixMultiply(m1, m2), [][]int{
		{8},
		{18},
	})

	m3 := [][]int{{1}, {4}}
	m4 := [][]int{{1, 2, 3}}
	assert.Equal(t, MatrixMultiply(m3, m4), [][]int{
		{1, 2, 3},
		{4, 8, 12},
	})
}

func TestInsertInto(t *testing.T) {
	assert.Equal(t,
		InsertInto([]int{1, 2, 3, 4, 5}, 2, []int{10, 11, 12, 13}),
		[]int{1, 2, 10, 11, 12, 13, 3, 4, 5},
	)
}
