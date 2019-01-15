package main

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/ovh/cds/sdk"
	"github.com/ovh/cds/sdk/grpcplugin/artifactplugin"
)

/*
This plugin have to be used as a deployment platform plugin

Arsenal deployment plugin must configured as following:
	name: plugin-artifact-openstack
	type: platform
	author: "Yvonnick Esnault"
	description: "OVH Openstack Artifact Storage Plugin"

$ cdsctl admin plugins import plugin-artifact-openstack.yml

Build the present binaries and import in CDS:
	os: linux
	arch: amd64
	cmd: <path-to-binary-file>

$ cdsctl admin plugins binary-add plugin-artifact-openstack plugin-artifact-openstack-bin.yml <path-to-binary-file>

Arsenal platform must configured as following
	name: Openstack
	storage: true
	default_config:
	"address":
		type: string
		value: ""
	"region":
		type: string
		value: ""
	"domain":
		type: string
		value: ""
	"tenant":
		type: string
		value: ""
	"user":
		type: string
		value: ""
	"password":
		type: password
		value: xxxxxxxx

*/

type openstackArtifactPlugin struct {
	artifactplugin.Common
}

func (e *openstackArtifactPlugin) Manifest(ctx context.Context, _ *empty.Empty) (*artifactplugin.ArtifactPluginManifest, error) {
	return &artifactplugin.ArtifactPluginManifest{
		Name:        "OVH Openstack Artifact Storage Plugin",
		Author:      "Yvonnick Esnault",
		Description: "OVH Openstack Artifact Storage Plugin",
		Version:     sdk.VERSION,
	}, nil
}

// ArtifactUpload implementation
func (e *openstackArtifactPlugin) ArtifactUpload(ctx context.Context, q *artifactplugin.Options) (*artifactplugin.Result, error) {
	// 	q.GetOptions()["cds.platform.address"]
	//  q.GetOptions()["cds.platform.region"]
	//  q.GetOptions()["cds.platform.domain"]
	//  q.GetOptions()["cds.platform.tenant"]
	//  q.GetOptions()["cds.platform.user"]
	//  q.GetOptions()["cds.platform.password"]
	return &artifactplugin.Result{
		Details: "none",
		Status:  "success",
	}, nil
}

// ArtifactDownload implementation
func (e *openstackArtifactPlugin) ArtifactDownload(ctx context.Context, q *artifactplugin.Options) (*artifactplugin.Result, error) {
	return &artifactplugin.Result{
		Details: "none",
		Status:  "success",
	}, nil
}

// ServeStaticFiles implementation
func (e *openstackArtifactPlugin) ServeStaticFiles(ctx context.Context, q *artifactplugin.Options) (*artifactplugin.Result, error) {
	return &artifactplugin.Result{
		Details: "none",
		Status:  "success",
	}, nil
}

// CachePull implementation
func (e *openstackArtifactPlugin) CachePull(ctx context.Context, q *artifactplugin.Options) (*artifactplugin.Result, error) {
	return &artifactplugin.Result{
		Details: "none",
		Status:  "success",
	}, nil
}

// CachePush implementation
func (e *openstackArtifactPlugin) CachePush(ctx context.Context, q *artifactplugin.Options) (*artifactplugin.Result, error) {
	return &artifactplugin.Result{
		Details: "none",
		Status:  "success",
	}, nil
}

func main() {
	e := openstackArtifactPlugin{}
	if err := artifactplugin.Start(context.Background(), &e); err != nil {
		panic(err)
	}
	return
}

func fail(format string, args ...interface{}) (*artifactplugin.Result, error) {
	msg := fmt.Sprintf(format, args...)
	fmt.Println(msg)
	return &artifactplugin.Result{
		Details: msg,
		Status:  sdk.StatusFail.String(),
	}, nil
}
