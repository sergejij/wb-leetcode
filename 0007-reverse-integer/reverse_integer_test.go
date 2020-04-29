package reverse_integer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testReverse = "Test Reverse Integer"
)

var (
	numbers       = []int{-1, 0, 1, 100, -100, -2147483648, 2147483647, 7463847412, -8463847412}
	expectResults = []int{-1, 0, 1, 1, -1, 0, 0, 0, 0}
)

func Test_reverse(t *testing.T) {

	t.Run(testReverse, func(t *testing.T) {
		for i := 0; i < len(numbers); i++ {
			realResult := reverse(numbers[i])

			assert.Equal(t, realResult, expectResults[i])
		}
	})

}
