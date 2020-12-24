package cli

import (
	"github.com/google/uuid"
	"github.com/zodiac1214/go/builder/gen"
	"testing"
)

func TestRunGen(t *testing.T) {
	t.Run("Force build", func(t *testing.T) {
		param := gen.GenParameters{
			Name:  uuid.New().String(),
			Path:  "/tmp",
			Force: true,
		}
		RunGen(param)
	})

	t.Run("Rebuild", func(t *testing.T) {
		param := gen.GenParameters{
			Name:  uuid.New().String(),
			Path:  "/tmp",
			Force: false,
		}
		RunGen(param)
	})
}
