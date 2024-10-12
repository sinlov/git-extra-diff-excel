package cli

import (
	"fmt"
	"github.com/sinlov/git-extra-diff-excel/command"
	"github.com/sinlov/git-extra-diff-excel/command/subcommand_new"
	"github.com/sinlov/git-extra-diff-excel/internal/pkg_kit"
	"github.com/sinlov/git-extra-diff-excel/internal/urfave_cli"
	"github.com/sinlov/git-extra-diff-excel/internal/urfave_cli/cli_exit_urfave"
	"github.com/urfave/cli/v2"
	"runtime"
	"time"
)

const (
	copyrightStartYear = "2023"
	defaultExitCode    = 1
)

func NewCliApp(buildId string) *cli.App {
	cli_exit_urfave.ChangeDefaultExitCode(defaultExitCode)
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Version = pkg_kit.GetPackageJsonVersionGoStyle(false)
	app.Name = pkg_kit.GetPackageJsonName()
	if pkg_kit.GetPackageJsonHomepage() != "" {
		app.Usage = fmt.Sprintf("see: %s", pkg_kit.GetPackageJsonHomepage())
	}
	app.Description = pkg_kit.GetPackageJsonDescription()
	year := time.Now().Year()
	jsonAuthor := pkg_kit.GetPackageJsonAuthor()
	app.Copyright = fmt.Sprintf("© %s-%d %s by: %s, build id: %s, run on %s %s",
		copyrightStartYear, year, jsonAuthor.Name, runtime.Version(), buildId, runtime.GOOS, runtime.GOARCH)
	author := &cli.Author{
		Name:  jsonAuthor.Name,
		Email: jsonAuthor.Email,
	}
	app.Authors = []*cli.Author{
		author,
	}

	flags := urfave_cli.UrfaveCliAppendCliFlag(command.GlobalFlag(), command.HideGlobalFlag())

	app.Flags = flags
	app.Before = command.GlobalBeforeAction
	app.Action = command.GlobalAction
	app.After = command.GlobalAfterAction

	var appCommands []*cli.Command
	appCommands = urfave_cli.UrfaveCliAppendCliCommand(appCommands, subcommand_new.Command())

	app.Commands = appCommands

	return app
}
