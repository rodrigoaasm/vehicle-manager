package utils_test

import (
	"demo/external/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {

	input := "string"

	assert.True(t, utils.IsThisType[string](input), "should be true when input is an object")
	assert.False(t, utils.IsThisType[int](input), "should be false")
	assert.True(t, utils.IsThisType[string](&input), "should be true when input is an pointer")
}
