package cmd

import (
	"fmt"
	"strings"
	"toyjvm/pkg/classpath"

	"github.com/urfave/cli/v2"
)

func NewApp(version string) *cli.App {
	app := &cli.App{
		Name:    "java-go",
		Usage:   "Toy JVM written in Go",
		Flags:   createFlags(),
		Version: version,
		Action:  run,
	}

	return app
}

func run(ctx *cli.Context) error {
	cp := classpath.Parse(ctx.String("jre-classpath"), ctx.String("classpath"))
	classArg := ctx.Args().Get(0)
	args := ctx.Args().Slice()[1:]
	fmt.Printf("classpath: %s, class: %s, args: %+v\n", cp, classArg, args)

	className := strings.Replace(classArg, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		return fmt.Errorf("Could not find or load main class %s\n", classArg)
	}

	fmt.Printf("class data: %v\n", classData)
	return nil
}

func createFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "classpath",
			Aliases: []string{"cp"},
			Value:   "",
			Usage:   "classpath",
		},
		&cli.StringFlag{
			Name:    "jre-classpath",
			Aliases: []string{"jre"},
			Value:   "",
			Usage:   "classpath of JRE",
		},
	}
}
