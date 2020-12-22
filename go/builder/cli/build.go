package cli

import (
	"context"
	"github.com/zodiac1214/go/builder/build"
	"os"
)

func RunBuild(param build.BuildParameters) {
	if param.Force {
		os.RemoveAll("./dist")
	}
	// copy dir to dist folder before build
	os.MkdirAll("./dist", 0755)
	build.CopyDirectory(param.Path, "./dist")

	// extract image list
	build.Helm(context.Background(), param)
	build.Kubernetes(context.Background(), param)
	build.Manifest(context.Background(), param)

	// pack docker images
	build.Docker(context.Background(), param)

	// packer VM images
	build.Packer(context.Background(), param)
}
