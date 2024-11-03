package main

import (
	"context"
	"io"
	"os"
	"testing"

	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

func TestMainFunctionExecutesWithoutErrors(t *testing.T) {
	if !checkPrivileges() {
		t.Fatalf("E1: checkPrivileges; want true; got false; run as administrator", os.Geteuid())
	}
	// Set up context
	contextBackground := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		t.Fatalf("E1: err; want nil; got %+v", err)
	}
	defer cli.Close()

	reader, err := cli.ImagePull(contextBackground, "golang", image.PullOptions{})
	if err != nil {
		t.Fatal(err)
	}
	defer reader.Close()

	io.Copy(os.Stdout, reader)
	t.Log("done")
}
