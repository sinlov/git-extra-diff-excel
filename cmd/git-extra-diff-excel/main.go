//go:build !test

package main

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/sinlov/git-extra-diff-excel"
	"github.com/sinlov/git-extra-diff-excel/cmd/cli"
	"github.com/sinlov/git-extra-diff-excel/constant"
	"github.com/sinlov/git-extra-diff-excel/internal/d_log"
	"github.com/sinlov/git-extra-diff-excel/internal/pkg_kit"
	os "os"
)

const (
	// exitCodeCmdArgs SIGINT as 2
	exitCodeCmdArgs = 2
)

//nolint:gochecknoglobals
var (
	// Populated by goreleaser during build
	version    = "unknown"
	rawVersion = "unknown"
	buildID    string
	commit     = "?"
	date       = ""
)

func init() {
	if buildID == "" {
		buildID = "unknown"
	}
}

func main() {
	d_log.SetLogLineDeep(d_log.DefaultExtLogLineMaxDeep)
	pkg_kit.InitPkgJsonContent(git_extra_diff_excel.PackageJson)

	bdInfo := pkg_kit.NewBuildInfo(
		pkg_kit.GetPackageJsonName(),
		version,
		rawVersion,
		buildID,
		commit,
		date,
		pkg_kit.GetPackageJsonAuthor().Name,
		constant.CopyrightStartYear,
	)

	app := cli.NewCliApp(bdInfo)

	args := os.Args
	if len(args) < 2 {
		fmt.Printf("%s %s --help\n", color.Yellow.Render("please see help as:"), app.Name)
		os.Exit(exitCodeCmdArgs)
	}
	if err := app.Run(args); nil != err {
		color.Redf("cli err at %v\n", err)
		os.Exit(exitCodeCmdArgs)
	}
}
