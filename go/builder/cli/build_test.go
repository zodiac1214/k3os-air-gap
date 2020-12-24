package cli

import (
	"github.com/zodiac1214/go/builder/build"
	"testing"
)

func TestRunBuild(t *testing.T) {
	t.Run("Force build", func(t *testing.T) {
		param := build.BuildParameters{
			Path:      "/tmp",
			Force:     true,
			ImageType: "vagrant",
		}
		RunBuild(param)
	})

	t.Run("Rebuild", func(t *testing.T) {
		param := build.BuildParameters{
			Path:      "/tmp",
			Force:     false,
			ImageType: "vagrant",
		}
		RunBuild(param)
	})
}
