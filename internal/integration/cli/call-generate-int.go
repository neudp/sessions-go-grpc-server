package cli

import (
	"errors"
	grpc_client "go-grpc-server/internal/integration/grpc-client"
	random_number_gen "go-grpc-server/internal/proto/random-number-gen/v1"
	"strconv"

	"github.com/spf13/cobra"
)

func NewCallGenerateIntCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "call-generate-int",
		Short: "Make gRPC generate int call",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			lowBorder, highBorder := args[0], args[1]

			if lowBorder == "" || highBorder == "" {
				return errors.New("empty borders")
			}

			lowBorderInt, err := strconv.ParseInt(lowBorder, 10, 64)
			if err != nil {
				return err
			}

			highBorderInt, err := strconv.ParseInt(highBorder, 10, 64)
			if err != nil {
				return err
			}

			conn, err := grpc_client.NewClientConnection(":50051")
			if err != nil {
				return err
			}

			client := random_number_gen.NewRandomNumberGenClient(conn)

			response, err := client.GenerateInt(cmd.Context(), &random_number_gen.GenerateIntRequest{
				LowBorder:  lowBorderInt,
				HighBorder: highBorderInt,
			})
			if err != nil {
				return err
			}

			println(response.Result)

			return nil
		},
	}

	return cmd
}
