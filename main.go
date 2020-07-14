package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

//go:generate vergitgo -o genautoversion.go gengo
//go:generate gofmt -w genautoversion.go

// cBINARY is the name of the binary
const cBINARY = "vergitgo"

func main() {
	app := cli.NewApp()
	app.Name = cBINARY
	app.Usage = "get version string from git"
	app.Version = VersionInfo
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "outFile, o",
			Usage: "go file to generate",
			Value: "genautoversion.go",
		},
		cli.StringFlag{
			Name:  "packageName, p",
			Usage: "name of the go package",
			Value: "main",
		},
		cli.StringFlag{
			Name:  "versionVariableName, n",
			Usage: "Name of the Version variable",
			Value: "VersionInfo",
		},
		cli.StringFlag{
			Name:  "buildVariableName, b",
			Usage: "Name of the Build variable",
			Value: "BuildInfo",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "gengo",
			Aliases: []string{"g"},
			Usage:   "generate go file",
			Action: func(c *cli.Context) error {
				outFile := c.GlobalString("outFile")
				packageName := c.GlobalString("packageName")
				versionVariableName := c.GlobalString("versionVariableName")
				buildVariableName := c.GlobalString("buildVariableName")
				return run(outFile, packageName, versionVariableName, buildVariableName)
			},
		},
		{
			Name:    "print",
			Aliases: []string{"p"},
			Usage:   "just print the version string",
			Action: func(c *cli.Context) error {
				v, err := getVersionInfo()
				if err != nil {
					return err
				}
				fmt.Println(v)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("Error executing Run: %+v", err.Error())
	}
}
