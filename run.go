package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"text/template"
	"time"
)

func getVersionInfo() (string, error) {
	out, err := exec.Command("git", "describe", "--long", "--tags", "--dirty", "--always").CombinedOutput()
	if err != nil {
		return "", err
	}
	//var host, thisVersionInfo, thisBuildInfo string
	thisVersionInfo := strings.TrimSpace(string(out))
	return thisVersionInfo, nil
}

const fileTemplate = `package {{.PackageName}}

// This file is created by hfgit, {{.HfGitVersionInfo}} {{.HfGitBuildInfo}}

// VersionInfo contains the version information
const {{.VersionInfoVariableName}} = "{{.VersionInfo}}"

// BuildInfo contains the information about the go generate run
const {{.BuildInfoVariableName}} = "{{.BuildInfo}}"
`

func run(outFile, packageName, versionVariableName, buildVariableName string) error {
	thisVersionInfo, err := getVersionInfo()
	if err != nil {
		return err
	}
	host, err := os.Hostname()
	if err != nil {
		return err
	}
	user, err := user.Current()
	if err != nil {
		return err
	}
	now := time.Now()
	thisBuildInfo := fmt.Sprintf("%s@%s %s", user.Username, host, now.Format(time.RFC3339))
	//fmt.Printf(string(out))

	// init template// Create a new template and parse the letter into it.
	t := template.Must(template.New("version").Parse(fileTemplate))

	// create anonymous struct
	data := struct {
		HfGitVersionInfo        string
		HfGitBuildInfo          string
		PackageName             string
		VersionInfoVariableName string
		BuildInfoVariableName   string
		VersionInfo             string
		BuildInfo               string
	}{
		VersionInfo,
		BuildInfo,
		packageName,
		versionVariableName,
		buildVariableName,
		thisVersionInfo,
		thisBuildInfo,
	}
	file, err := os.Create(outFile)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(file, data)
	return nil
}
