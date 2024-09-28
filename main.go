package main

import (
	"context"
	"fmt"
	"os"
	"r_slash_place/server"
)


func main() {
	ctx := context.Background()
	if err := server.Run(ctx, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}