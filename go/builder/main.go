package main

import (
	"fmt"
	"github.com/zodiac1214/go/builder/cli"
	"github.com/zodiac1214/go/builder/gen"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

func main() {
	app := kingpin.New("builder", "A tool to build VM images for kubernates application")
	builder := cli.RegisterCommands(app)
	cmd, err := builder.Parse(os.Args[1:])
	if err != nil {
		builder.Application.FatalUsage("%s", err.Error())
	}
	switch cmd {
	case builder.GenCmd.FullCommand():
		cli.RunGen(gen.GenParameters{Name: *builder.GenCmd.Name, Path: *builder.GenCmd.Path})
	case builder.BuildCmd.FullCommand():
		fmt.Println("build run")
	}
}
