package main

import (
	"context"
	"fmt"
	"os"
	"runtime/debug"

	"dagger.io/dagger"
)

var (
	mountedDir = "article-manager"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			stackTrace := debug.Stack()
			fmt.Println(string(stackTrace))
		}
	}()
	ctx := context.Background()

	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	container := client.Container().From("golang:1.20")
	container = container.
		WithMountedDirectory(mountedDir, client.Host().Directory(".")).
		WithWorkdir(mountedDir)

	//container = container.WithExec([]string{"go", "run", "main.go"})

	if _, err := container.ExitCode(ctx); err != nil {
		panic(err)
	}
}
