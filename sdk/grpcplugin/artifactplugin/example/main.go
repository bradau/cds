package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/ovh/cds/sdk"
	"github.com/ovh/cds/sdk/grpcplugin/artifactplugin"
)

// ExamplePlugin represents an example plugins
type ExamplePlugin struct {
	artifactplugin.Common
}

// Manifest implementation
func (e *ExamplePlugin) Manifest(ctx context.Context, _ *empty.Empty) (*artifactplugin.ArtifactPluginManifest, error) {
	return &artifactplugin.ArtifactPluginManifest{
		Name:        "Example Plugin",
		Author:      "Yvonnick Esnault",
		Description: "This is an example plugin",
		Version:     sdk.VERSION,
	}, nil
}

// ArtifactUpload implementation
func (e *ExamplePlugin) ArtifactUpload(ctx context.Context, q *artifactplugin.Options) (*artifactplugin.Result, error) {
	return &artifactplugin.Result{
		Details: "none",
		Status:  "success",
	}, nil
}

// ArtifactDownload implementation
func (e *ExamplePlugin) ArtifactDownload(ctx context.Context, q *artifactplugin.Options) (*artifactplugin.Result, error) {
	return &artifactplugin.Result{
		Details: "none",
		Status:  "success",
	}, nil
}

// ServeStaticFiles implementation
func (e *ExamplePlugin) ServeStaticFiles(ctx context.Context, q *artifactplugin.Options) (*artifactplugin.Result, error) {
	return &artifactplugin.Result{
		Details: "none",
		Status:  "success",
	}, nil
}

// CachePull implementation
func (e *ExamplePlugin) CachePull(ctx context.Context, q *artifactplugin.Options) (*artifactplugin.Result, error) {
	return &artifactplugin.Result{
		Details: "none",
		Status:  "success",
	}, nil
}

// CachePush implementation
func (e *ExamplePlugin) CachePush(ctx context.Context, q *artifactplugin.Options) (*artifactplugin.Result, error) {
	return &artifactplugin.Result{
		Details: "none",
		Status:  "success",
	}, nil
}

func main() {
	if os.Args[1:][0] == "serve" {
		e := ExamplePlugin{}
		if err := artifactplugin.Start(context.Background(), &e); err != nil {
			panic(err)
		}
		return
	}

	//Server Part - BEGIN
	var e *ExamplePlugin
	go func() {
		e = &ExamplePlugin{}
		if err := artifactplugin.Start(context.Background(), e); err != nil {
			panic(err)
		}
	}()
	//Server Part - END

	time.Sleep(100 * time.Millisecond)

	//Client Part - BEGIN
	c, err := artifactplugin.Client(context.Background(), e.Socket)
	if err != nil {
		panic(err)
	}

	manifest, err := c.Manifest(context.Background(), new(empty.Empty))
	if err != nil {
		panic(err)
	}

	fmt.Println(manifest)
	//Client part - END
}
