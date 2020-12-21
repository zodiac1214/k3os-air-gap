package cli

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/alecthomas/kingpin.v2"
	"testing"
)

func newTestApp() *kingpin.Application {
	return kingpin.New("test_commands", "lol").Terminate(nil)
}
func TestRegisterCommands(t *testing.T) {
	app := newTestApp()
	RegisterCommands(app)

	_, err := app.Parse([]string{"asdf"})
	assert.Error(t, err)
}
