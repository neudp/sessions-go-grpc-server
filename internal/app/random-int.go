package app

import (
	"fmt"
)

var ErrMinGreaterOrEqualMax = fmt.Errorf("minimum value cannot be greater than or equal to maximum value")

type RandomIntGenerator interface {
	GenerateInt(min, max int64) (int64, error)
}
