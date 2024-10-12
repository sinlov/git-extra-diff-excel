package command

import (
	"fmt"
	"github.com/sinlov/git-extra-diff-excel/constant"
	"github.com/sinlov/git-extra-diff-excel/internal/d_log"
	"github.com/sinlov/git-extra-diff-excel/internal/pkg_kit"
	"github.com/sinlov/git-extra-diff-excel/internal/urfave_cli/cli_exit_urfave"
	"github.com/urfave/cli/v2"
)

type GlobalConfig struct {
	TimeoutSecond uint
}

type (
	// GlobalCommand
	//	command root
	GlobalCommand struct {
		Name    string
		Version string
		Verbose bool
		RootCfg GlobalConfig
	}
)

var (
	cmdGlobalEntry *GlobalCommand
)

// CmdGlobalEntry
//
//	return global command entry
func CmdGlobalEntry() *GlobalCommand {
	return cmdGlobalEntry
}

// globalExec
//
//	do global command exec
func (c *GlobalCommand) globalExec() error {

	d_log.Debug("-> start GlobalAction")

	return nil
}

// withGlobalFlag
//
// bind global flag to globalExec
func withGlobalFlag(c *cli.Context, cliVersion, cliName string) (*GlobalCommand, error) {
	d_log.Debug("-> withGlobalFlag")

	isVerbose := c.Bool(constant.NameKeyCliVerbose)
	config := GlobalConfig{
		TimeoutSecond: c.Uint(constant.NamePluginTimeOut),
	}

	p := GlobalCommand{
		Name:    cliName,
		Version: cliVersion,
		Verbose: isVerbose,
		RootCfg: config,
	}
	return &p, nil
}

// GlobalBeforeAction
// do command Action before flag global.
func GlobalBeforeAction(c *cli.Context) error {
	isVerbose := c.Bool(constant.NameKeyCliVerbose)
	if isVerbose {
		d_log.OpenDebug()
	}

	cliVersion := pkg_kit.GetPackageJsonVersionGoStyle(false)
	if isVerbose {
		d_log.Warnf("-> open verbose, and now command version is: %s", cliVersion)
	}
	appName := pkg_kit.GetPackageJsonName()
	cmdEntry, err := withGlobalFlag(c, cliVersion, appName)
	if err != nil {
		return cli_exit_urfave.Err(err)
	}

	cmdGlobalEntry = cmdEntry
	return nil
}

// GlobalAction
// do command Action flag.
func GlobalAction(c *cli.Context) error {
	if cmdGlobalEntry == nil {
		panic(fmt.Errorf("not init GlobalBeforeAction success to new cmdGlobalEntry"))
	}

	err := cmdGlobalEntry.globalExec()
	if err != nil {
		return cli_exit_urfave.Format("run GlobalAction err: %v", err)
	}
	return nil
}

// GlobalAfterAction
//
//	do command Action after flag global.
//
//nolint:golint,unused
func GlobalAfterAction(c *cli.Context) error {
	if cmdGlobalEntry != nil {
		if c.Bool(constant.NameKeyCliVerbose) {
			d_log.Infof("-> finish run command: %s, version %s", cmdGlobalEntry.Name, cmdGlobalEntry.Version)
		}
	}
	return nil
}
