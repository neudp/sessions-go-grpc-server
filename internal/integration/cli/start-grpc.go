package cli

import (
	"context"
	grpc2 "go-grpc-server/internal/integration/grpc"
	"go-grpc-server/internal/integration/random"
	random_number_gen "go-grpc-server/internal/proto/random-number-gen/v1"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

const (
	Byte int = 1 << (10 * iota)
	Kilobyte
	Megabyte
	Gigabyte
	Terabyte
	Petabyte
	Exabyte
)

func NewRunGrpcCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "start-grpc",
		Short: "Start gRPC server",
		RunE: func(cmd *cobra.Command, args []string) error {
			grpcServer := grpc.NewServer(
				grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
					MinTime:             time.Second, // How frequently a client can send a ping
					PermitWithoutStream: true,        // Allow pings even when there are no active streams
				}),
				grpc.KeepaliveParams(keepalive.ServerParameters{
					MaxConnectionIdle:     0,                // How many connections can be idle before being closed? 0 = unlimited
					MaxConnectionAge:      0,                // How old connections should be before it becomes idle 0 = in time
					MaxConnectionAgeGrace: 0,                // How long server can wait for client to finish processing 0 = wait until all contexts done
					Time:                  10 * time.Second, // How frequently client should send ping to be alive
					Timeout:               30 * time.Second, // How long server waits for ping ack
				}),
				grpc.MaxRecvMsgSize(2*Megabyte), // Max receive message size
				grpc.MaxSendMsgSize(2*Megabyte), // Max send message size
			)

			// OKM religion
			// service := grpc2.NewRandomNumberGenServer(
			//	 random.NewPureGoRandomIntGenerator(),
			// )
			// service.VisitRegistrar(grpcServer)

			random_number_gen.RegisterRandomNumberGenServer(
				grpcServer,
				grpc2.NewRandomNumberGenServer(
					random.NewPureGoRandomIntGenerator(),
				),
			)
			// Register reflection service on gRPC server.
			reflection.Register(grpcServer)

			// Subscribe to system signals
			signals := make(chan os.Signal, 1)
			signal.Notify(signals, os.Interrupt, os.Kill)

			// Create internal context
			cmdCtx := cmd.Context()
			ctx, cancel := context.WithCancelCause(cmdCtx)

			// Handle system signals and gracefully stop server
			go func() {
				select {
				case <-signals:
					cancel(context.Canceled)
				case <-cmdCtx.Done():
					cancel(ctx.Err()) // theoretically unnecessary
				}
			}()

			// Graceful stop server on context done
			go func() {
				<-ctx.Done()
				println("[GRPCServer] context done, stopping...")

				grpcServer.GracefulStop()
			}()

			// Create tcp listener
			listener, err := (&net.ListenConfig{}).Listen(ctx, "tcp", ":50051")
			if err != nil {
				println("failed to listen at :50051")

				return err
			}

			println("[GRPCServer] start listening")

			//// Start server in goroutine to enable multiprocessing (optionally)
			//wg := &sync.WaitGroup{}
			//wg.Add(1)
			//go func() {
			//	defer wg.Done()
			err = grpcServer.Serve(listener)

			if err != nil {
				println("[GRPCServer] server.Serve failed")
			}
			//}()

			//wg.Wait()

			println("[GRPCServer] stopped")

			return err
		},
	}
}
