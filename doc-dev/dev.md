# dev

## env

- minimum go version: go 1.21
- change `go 1.21`, `^1.21`, `1.21.13` to new go version
- change `golangci-lint@v1.59.1`
  from [golangci-lint version release](https://github.com/golangci/golangci-lint/releases) to new version
    - more info see [golangci-lint local-installation](https://golangci-lint.run/usage/install/#local-installation)

```
          # version: v1.61.0 for go verison 1.22.1+
          # version: v1.59.1 for go verison 1.21+
          # version: v1.55.2 for go verison 1.20+
          # version: v1.53.3 for go verison 1.19+
          # version: v1.45.2 for go verison 1.18+
          # version: v1.42.1 for go verison 1.17+
          # version: v1.41.0 for go verison 1.16+
```

### libs

| lib                                 | version |
|:------------------------------------|:--------|
| https://github.com/stretchr/testify | v1.9.0  |
| https://github.com/sebdah/goldie    | v2.5.5  |
| https://github.com/gookit/color     | v1.5.4  |
| https://github.com/urfave/cli/      | v2.27.4 |

- more libs see [go.mod](https://github.com/sinlov/git-extra-diff-excel/blob/main/go.mod)

## depends

in go mod project

```bash
# warning use private git host must set
# global set for once
# add private git host like github.com to evn GOPRIVATE
$ go env -w GOPRIVATE='github.com'
# use ssh proxy
# set ssh-key to use ssh as http
$ git config --global url."git@github.com:".insteadOf "https://github.com/"
# or use PRIVATE-TOKEN
# set PRIVATE-TOKEN as gitlab or gitea
$ git config --global http.extraheader "PRIVATE-TOKEN: {PRIVATE-TOKEN}"
# set this rep to download ssh as https use PRIVATE-TOKEN
$ git config --global url."ssh://github.com/".insteadOf "https://github.com/"

# before above global settings
# test version info
$ git ls-remote -q https://github.com/sinlov/git-extra-diff-excel.git

# test depends see full version
$ go list -mod readonly -v -m -versions github.com/sinlov/git-extra-diff-excel
# or use last version add go.mod by script
$ echo "go mod edit -require=$(go list -mod=readonly -m -versions github.com/sinlov/git-extra-diff-excel | awk '{print $1 "@" $NF}')"
$ echo "go mod vendor"
```

## dev tasks

```bash
# It needs to be executed after the first use or update of dependencies.
make init dep
```

- test code

```bash
make test
# benchmark and coverage show
make ci.test.benchmark ci.coverage.show
```

- ci to fast check as CI pipeline

```bash
# check style at local
make style

# run ci at local
make ci
```

### docker

```bash
# then test build as test/Dockerfile
$ make dockerTestRestartLatest
# clean test build
$ make dockerTestPruneLatest

# more info see
$ make helpDocker
```

### EngineeringStructure

```
.
├── Dockerfile                     # ci docker build
├── Dockerfile.s6                  # local docker build
├── Makefile                       # make entry
├── README.md
├── build                          # build output
├── cmd
│     └── git-extra-diff-excel     # command line main package install and dev entrance
│         ├── main.go                   # command line entry
│         └── main_test.go              # integrated test entry
├── command                        # command line package
│         ├── TestMain.go             # common entry in unit test package
│         ├── flag.go                 # global flag
│         ├── global.go               # global command
│         ├── global_test.go          # global command unit test
│         ├── golder_data_test.go     # unit test test data case
│         ├── init_test.go            # unit test initialization tool
│         └── subcommand_new          # subcommandPackage new
├── constant                       # constant package 
│         └── env.go                  # constant environment variable
├── doc                            # command line tools documentation
│         ├── cmd.md
│         └── dev.md
├── go.mod
├── go.sum
├── package.json                   # command line profile information
├── resource.go                    # embed resource 
├── utils                          # toolkit package
│         ├── env_kit                 # environment variables toolkit
│         ├── log                     # log toolkit
│         ├── pkgJson                 # package.json toolkit
│         └── urfave_cli              # urfave/cli toolkit
├── vendor
└── z-MakefileUtils                # make toolkit

```

### log

- cli log use [github.com/bar-counter/slog](https://github.com/bar-counter/slog)
    - open debug log by env `CLI_VERBOSE=true` or global flag `--verbose`

```go
package foo

import (
	"github.com/sinlov/git-extra-diff-excel/constant"
	"github.com/sinlov/git-extra-diff-excel/internal/d_log"
)

// GlobalBeforeAction
// do command Action before flag global.
func GlobalBeforeAction(c *cli.Context) error {
	isVerbose := c.Bool(constant.NameKeyCliVerbose)
	if isVerbose {
		d_log.OpenDebug()
	}
	return nil
}

func Biz() {
	d_log.Debug("debug log")
	d_log.Debugf("debug log")
}

```