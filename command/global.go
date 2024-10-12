package command

import (
	"fmt"
	"github.com/sinlov/git-extra-diff-excel/constant"
	"github.com/sinlov/git-extra-diff-excel/internal/d_log"
	"github.com/sinlov/git-extra-diff-excel/internal/pkg_kit"
	"github.com/sinlov/git-extra-diff-excel/internal/urfave_cli/cli_exit_urfave"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

type GlobalConfig struct {
	LogLevel      string
	TimeoutSecond uint
	ExecFullPath  string
	RunRootPath   string
}

type (
	// GlobalCommand
	//	command root
	GlobalCommand struct {
		Name     string
		Version  string
		BuildId  string
		HomePage string

		Args cli.Args

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
	if c.Verbose {
		d_log.Verbosef("cli exec full: %s", c.RootCfg.ExecFullPath)
		d_log.Verbosef("     run path: %s", c.RootCfg.RunRootPath)
		d_log.Verbosef("     args len: %v", c.Args.Len())
		if c.Args.Len() > 0 {
			d_log.Verbosef("     args content: %s", strings.Join(c.Args.Slice(), " | "))
		}
	}

	return nil
}

// withGlobalFlag
//
// bind global flag to globalExec
func withGlobalFlag(c *cli.Context, cliVersion, cliName, homePage string) (*GlobalCommand, error) {
	d_log.Debug("-> withGlobalFlag")

	isVerbose := c.Bool(constant.NameCliVerbose)

	execFullPath, err := os.Executable()
	if err != nil {
		return nil, fmt.Errorf("os.Executable can not found, err: %v", err)
	}

	cliRunRootPath := c.String(constant.NameCliRunPath)
	if len(cliRunRootPath) == 0 {
		rootDir, err := os.Getwd()
		if err != nil {
			d_log.Errorf(err, "get rooted path name corresponding to the current directory path err")
			return nil, cli_exit_urfave.Err(err)
		}
		cliRunRootPath = rootDir
	}

	timeoutSecond := c.Uint(constant.NameCliTimeoutSecond)
	if timeoutSecond < constant.MinimumTimeoutSecond {
		d_log.Warnf("timeout second is too small, will use default value: %d", constant.MinimumTimeoutSecond)
	}

	config := GlobalConfig{
		LogLevel:      c.String(constant.NameLogLevel),
		TimeoutSecond: timeoutSecond,
		ExecFullPath:  execFullPath,
		RunRootPath:   cliRunRootPath,
	}

	p := GlobalCommand{
		Name:     cliName,
		Version:  cliVersion,
		BuildId:  pkg_kit.FetchNowBuildId(),
		HomePage: homePage,

		Args: c.Args(),

		Verbose: isVerbose,
		RootCfg: config,
	}
	return &p, nil
}

// GlobalBeforeAction
// do command Action before flag global.
func GlobalBeforeAction(c *cli.Context) error {
	isVerbose := c.Bool(constant.NameCliVerbose)
	if isVerbose {
		d_log.OpenDebug()
	}

	appName := c.App.Name
	cliVersion := c.App.Version
	if isVerbose {
		d_log.Warnf("-> open verbose, and now command version is: %s", cliVersion)
	}
	homePage := pkg_kit.GetPackageJsonHomepage()

	cmdEntry, err := withGlobalFlag(c, cliVersion, appName, homePage)
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
		if c.Bool(constant.NameCliVerbose) {
			d_log.Infof("-> finish run command: %s, version %s", cmdGlobalEntry.Name, cmdGlobalEntry.Version)
		}
	}
	return nil
}
