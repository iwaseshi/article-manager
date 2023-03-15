package main

import (
	"context"
	"os"

	"dagger.io/dagger"
)

var (
	artclesPath = "../"
)

func main() {
	ctx := context.Background()

	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		panic(err)
	}
	defer client.Close()

}
