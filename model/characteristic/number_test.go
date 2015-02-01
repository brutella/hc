package characteristic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumber(t *testing.T) {
	number := NewNumber(20.2, 0, 100, 0.1, FormatFloat)
	assert.Equal(t, number.Format, FormatFloat)
	assert.Equal(t, number.GetValue(), 20.2)
	assert.Equal(t, number.GetMinValue(), 0)
	assert.Equal(t, number.GetMaxValue(), 100)
	assert.Equal(t, number.GetMinStepValue(), 0.1)

	number.SetValue(17)
	number.SetMinValue(0.3)
	number.SetMaxValue(200)
	number.SetMinStepValue(0.3)

	assert.Equal(t, number.GetValue(), 17)
	assert.Equal(t, number.GetMinValue(), 0.3)
	assert.Equal(t, number.GetMaxValue(), 200)
	assert.Equal(t, number.GetMinStepValue(), 0.3)
}
