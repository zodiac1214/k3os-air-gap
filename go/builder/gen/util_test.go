package gen

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateFolder(t *testing.T) {
	genParame := GenParameters{
		Name:  "testname",
		Path:  "/tmp",
		Force: false,
	}
	path, err := CreateFolder(genParame, "subFolder")
	assert.Nil(t, err)
	assert.Equal(t, "/tmp/testname/subFolder", path)
}

func TestCreateTextFile(t *testing.T) {
	err := CreateTextFile("/tmp", uuid.New().String(), "thisIsContent")
	assert.Nil(t, err)
}
