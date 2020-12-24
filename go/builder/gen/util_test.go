package gen

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateFolder(t *testing.T) {
	genParam := GenParameters{
		Name:  "testname",
		Path:  "/tmp",
		Force: false,
	}
	path, err := CreateFolder(genParam, "subFolder")
	assert.Nil(t, err)
	assert.Equal(t, "/tmp/testname/subFolder", path)

	t.Run("No permission to create folder", func(t *testing.T) {
		genParam := GenParameters{
			Name:  "testname",
			Path:  "/",
			Force: false,
		}
		_, err := CreateFolder(genParam, "subFolder")
		assert.NotNil(t, err)
	})
}

func TestCreateTextFile(t *testing.T) {
	err := CreateTextFile("/tmp", uuid.New().String(), "thisIsContent")
	assert.Nil(t, err)

	t.Run("No permission to create Text file", func(t *testing.T) {
		err := CreateTextFile("/", uuid.New().String(), "thisIsContent")
		assert.NotNil(t, err)
	})
}
