package main

import "go-grpc-server/internal/integration/cli"

func main() {
	if err := cli.Execute(); err != nil {
		panic(err)
	}
}
