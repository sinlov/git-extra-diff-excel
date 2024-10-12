package subcommand_csv

import (
	"github.com/sinlov/git-extra-diff-excel/command"
	"github.com/sinlov/git-extra-diff-excel/internal/d_log"
	"github.com/urfave/cli/v2"
	"strings"
)

const commandName = "csv"

var commandEntry *CsvCommand

type CsvCommand struct {
	cliName  string
	version  string
	buildId  string
	homePage string

	Args cli.Args

	isDebug      bool
	execFullPath string
	runRootPath  string

	isIgnoreRead bool
}

func (n *CsvCommand) Exec() error {
	d_log.Debugf("-> Exec cli [ %s ] by subCommand [ %s ], version %s buildID %s", n.cliName, commandName, n.version, n.buildId)
	if n.isDebug {
		d_log.Verbosef("cli full path: %s", n.execFullPath)
		d_log.Verbosef("     run path: %s", n.runRootPath)
		d_log.Verbosef("     args len: %v", n.Args.Len())
		if n.Args.Len() > 0 {
			d_log.Verbosef("     args content: %s", strings.Join(n.Args.Slice(), " | "))
		}
	}

	return nil
}

func flag() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:  "ignore-read",
			Usage: "ignore read excel file error",
			Value: false,
		},
	}
}

func withEntry(c *cli.Context) (*CsvCommand, error) {
	d_log.Debugf("-> withEntry subCommand [ %s ]", commandName)

	if c.Bool("lib") {
		d_log.Info("new lib mode")
	}
	globalEntry := command.CmdGlobalEntry()
	return &CsvCommand{
		cliName:  globalEntry.Name,
		version:  globalEntry.Version,
		buildId:  globalEntry.BuildId,
		homePage: globalEntry.HomePage,

		Args: c.Args(),

		isDebug:      globalEntry.Verbose,
		execFullPath: globalEntry.RootCfg.ExecFullPath,
		runRootPath:  globalEntry.RootCfg.RunRootPath,

		isIgnoreRead: c.Bool("ignore-read"),
	}, nil
}

func action(c *cli.Context) error {
	d_log.Debugf("-> Sub Command action [ %s ] start", commandName)
	entry, err := withEntry(c)
	if err != nil {
		return err
	}
	commandEntry = entry
	return commandEntry.Exec()
}

func Command() []*cli.Command {
	return []*cli.Command{
		{
			Name:   commandName,
			Usage:  "",
			Action: action,
			Flags:  flag(),
		},
	}
}
