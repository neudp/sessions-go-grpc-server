package cli

import "github.com/spf13/cobra"

func Execute() error {
	rootCmd := &cobra.Command{
		Use:   "random-number-int",
		Short: "Generate random integer",
	}

	rootCmd.AddCommand(
		NewRunGrpcCommand(),
		NewCallGenerateIntCommand(),
	)

	return rootCmd.Execute()
}
