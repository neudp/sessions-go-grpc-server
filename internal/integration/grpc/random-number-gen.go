package grpc

import (
	"context"
	"errors"
	"fmt"
	"go-grpc-server/internal/app"
	random_number_gen "go-grpc-server/internal/proto/random-number-gen/v1"
	"runtime/debug"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RandomNumberGenServer struct {
	random_number_gen.UnimplementedRandomNumberGenServer
	randomIntGenerator app.RandomIntGenerator
	param              string
}

func NewRandomNumberGenServer(
	randomIntGenerator app.RandomIntGenerator,
) *RandomNumberGenServer {
	return &RandomNumberGenServer{
		randomIntGenerator: randomIntGenerator,
		param:              "_",
	}
}

// VisitRegistrar is a part of OKM religion
func (server *RandomNumberGenServer) VisitRegistrar(registrar grpc.ServiceRegistrar) {
	random_number_gen.RegisterRandomNumberGenServer(registrar, server)
}

func (server *RandomNumberGenServer) GenerateInt(
	_ context.Context,
	request *random_number_gen.GenerateIntRequest,
) (out *random_number_gen.GenerateIntResponse, err error) {
	defer func() {
		if r := recover(); r != nil {
			panicErr, ok := r.(error)
			if !ok {
				panicErr = errors.New("panic: " + fmt.Sprintf("%+v", r))
			}
			println(panicErr.Error())
			debug.PrintStack()

			err = status.Error(codes.Internal, "something went wrong:(")
		}
	}()

	result, err := server.randomIntGenerator.GenerateInt(request.LowBorder, request.HighBorder)

	if err != nil {
		var appErr *app.AppError
		ok := errors.As(err, &appErr)

		if ok {
			switch true {
			case errors.Is(appErr, app.ErrInvalidArgument):
				return nil, status.Error(codes.InvalidArgument, appErr.Message())
			}
		}

		println(err.Error())
		return nil, status.Error(codes.Internal, "something went wrong:(")
	}

	return &random_number_gen.GenerateIntResponse{Result: result}, err
}
