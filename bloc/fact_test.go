package bloc

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestFactorial(t *testing.T) {
	assert.Equal(t, Factorial(10), 3628800)
	assert.Equal(t, Factorial(9), 362880)
	assert.Equal(t, Factorial(8), 40320)
	assert.Equal(t, Factorial(1), 1)
	assert.Equal(t, Factorial(0), 1)
}
