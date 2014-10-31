package characteristic

import (
	"testing"
    "github.com/stretchr/testify/assert"
    "reflect"
)

func TestConvertIntToFloat32(t *testing.T) {
    to := reflect.TypeOf(float32(0))
    assert.Equal(t, ConvertValue(10, to), 10.0)
}

func TestConvertFloat32ToInt64(t *testing.T) {
    to := reflect.TypeOf(int64(0))
    assert.Equal(t, ConvertValue(float32(10.5), to), int64(10))
}
