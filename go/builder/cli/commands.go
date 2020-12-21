package cli

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

type App struct {
	*kingpin.Application
	// Debug allows to run the command in debug mode
	Debug *bool
	// BuildCmd builds a cluster image.
	BuildCmd buildCmd
	// GenCmd bootstraps a new project
	GenCmd genCmd
}

type buildCmd struct {
	*kingpin.CmdClause
	// path of bootstrap project
	Path *string
}

type genCmd struct {
	*kingpin.CmdClause
	// name of bootstrap project
	Name *string
	// path of bootstrap project
	Path *string
}

func RegisterCommands(app *kingpin.Application) App {
	builder := App{
		Application: app,
	}

	builder.Debug = app.Flag("debug", "Enable debug mode.").Bool()

	builder.GenCmd.CmdClause = app.Command("gen", "bootstrap a project")
	builder.GenCmd.Name = builder.GenCmd.Flag("name", "name of the project").Default("my-project").String()
	builder.GenCmd.Path = builder.GenCmd.Flag("path", "output path of the project").Default(".").String()

	builder.BuildCmd.CmdClause = app.Command("build", "bootstrap a project")
	builder.BuildCmd.Path = builder.BuildCmd.Flag("path", "output path of the project").Default(".").String()

	//builder.UsageTemplate(kingpin.DefaultUsageTemplate)

	return builder
}
