package gen

import "testing"

func TestHelm(t *testing.T) {
	param := GenParameters{
		Name:  "test",
		Path:  "/tmp",
		Force: false,
	}
	Helm(param)
}
