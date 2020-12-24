package gen

import "testing"

func TestManifest(t *testing.T) {
	param := GenParameters{
		Name:  "test",
		Path:  "/tmp",
		Force: false,
	}
	Manifest(param)
}
