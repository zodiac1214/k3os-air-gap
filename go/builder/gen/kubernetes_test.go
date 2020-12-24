package gen

import "testing"

func TestKubernetes(t *testing.T) {
	param := GenParameters{
		Name:  "test",
		Path:  "/tmp",
		Force: false,
	}
	Kubernetes(param)
}
