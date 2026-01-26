package random

import (
	"go-grpc-server/internal/app"
	"math/rand"
)

type PureGoRandomIntGenerator struct{}

func NewPureGoRandomIntGenerator() *PureGoRandomIntGenerator {
	return &PureGoRandomIntGenerator{}
}

func (*PureGoRandomIntGenerator) GenerateInt(min, max int64) (int64, error) {
	if min >= max {
		return 0, app.NewAppError(
			"min cannot be greater or equal to max",
			app.ErrInvalidArgument,
			app.ErrMinGreaterOrEqualMax,
		)
	}

	return min + rand.Int63n(max-min+1), nil
}
