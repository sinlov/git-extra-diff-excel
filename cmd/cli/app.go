package cli

import (
	"fmt"
	"github.com/sinlov/git-extra-diff-excel/command"
	"github.com/sinlov/git-extra-diff-excel/command/subcommand_new"
	"github.com/sinlov/git-extra-diff-excel/internal/pkg_kit"
	"github.com/sinlov/git-extra-diff-excel/internal/urfave_cli"
	"github.com/sinlov/git-extra-diff-excel/internal/urfave_cli/cli_exit_urfave"
	"github.com/urfave/cli/v2"
)

const (
	// defaultExitCode SIGINT as 2
	defaultExitCode = 1
)

func NewCliApp(bdInfo pkg_kit.BuildInfo) *cli.App {
	cli_exit_urfave.ChangeDefaultExitCode(defaultExitCode)
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Name = bdInfo.PgkNameString()
	app.Version = bdInfo.VersionString()
	if pkg_kit.GetPackageJsonHomepage() != "" {
		app.Usage = fmt.Sprintf("see: %s", pkg_kit.GetPackageJsonHomepage())
	}
	app.Description = pkg_kit.GetPackageJsonDescription()
	jsonAuthor := pkg_kit.GetPackageJsonAuthor()
	app.Copyright = bdInfo.String()
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
