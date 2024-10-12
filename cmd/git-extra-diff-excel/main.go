//go:build !test

package main

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/sinlov/git-extra-diff-excel"
	"github.com/sinlov/git-extra-diff-excel/cmd/cli"
	"github.com/sinlov/git-extra-diff-excel/internal/d_log"
	"github.com/sinlov/git-extra-diff-excel/internal/pkg_kit"
	os "os"
)

const (
	exitCodeCmdArgs = 2
)

var buildID string

func init() {
	if buildID == "" {
		buildID = "unknown"
	}
}

func main() {
	d_log.SetLogLineDeep(d_log.DefaultExtLogLineMaxDeep)
	pkg_kit.InitPkgJsonContent(git_extra_diff_excel.PackageJson)

	app := cli.NewCliApp(buildID)

	args := os.Args
	if len(args) < 2 {
		fmt.Printf("%s %s --help\n", color.Yellow.Render("please see help as:"), app.Name)
		os.Exit(exitCodeCmdArgs)
	}
	if err := app.Run(args); nil != err {
		color.Redf("cli err at %v\n", err)
	}
}
