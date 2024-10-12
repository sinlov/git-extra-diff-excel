package command

import (
	"github.com/sinlov/git-extra-diff-excel/constant"
	"github.com/urfave/cli/v2"
)

// GlobalFlag
// Other modules also have flags
func GlobalFlag() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:    constant.NameCliVerbose,
			Usage:   "open cli verbose mode",
			Value:   false,
			EnvVars: []string{constant.EnvKeyCliVerbose},
		},
	}
}

func HideGlobalFlag() []cli.Flag {
	return []cli.Flag{
		&cli.UintFlag{
			Name:    constant.NameCliTimeoutSecond,
			Usage:   "command timeout setting second",
			Hidden:  true,
			Value:   10,
			EnvVars: []string{constant.EnvKeyCliTimeoutSecond},
		},
	}
}
