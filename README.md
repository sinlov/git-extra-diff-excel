[![ci](https://github.com/sinlov/git-extra-diff-excel/workflows/ci/badge.svg)](https://github.com/sinlov/git-extra-diff-excel/actions/workflows/ci.yml)

[![go mod version](https://img.shields.io/github/go-mod/go-version/sinlov/git-extra-diff-excel?label=go.mod)](https://github.com/sinlov/git-extra-diff-excel)
[![GoDoc](https://godoc.org/github.com/sinlov/git-extra-diff-excel?status.png)](https://godoc.org/github.com/sinlov/git-extra-diff-excel)
[![goreportcard](https://goreportcard.com/badge/github.com/sinlov/git-extra-diff-excel)](https://goreportcard.com/report/github.com/sinlov/git-extra-diff-excel)

[![GitHub license](https://img.shields.io/github/license/sinlov/git-extra-diff-excel)](https://github.com/sinlov/git-extra-diff-excel)
[![codecov](https://codecov.io/gh/sinlov/git-extra-diff-excel/branch/main/graph/badge.svg)](https://codecov.io/gh/sinlov/git-extra-diff-excel)
[![GitHub latest SemVer tag)](https://img.shields.io/github/v/tag/sinlov/git-extra-diff-excel)](https://github.com/sinlov/git-extra-diff-excel/tags)
[![GitHub release)](https://img.shields.io/github/v/release/sinlov/git-extra-diff-excel)](https://github.com/sinlov/git-extra-diff-excel/releases)

## Read Usage

- [usage en-US](#usage)
- [中文使用说明 zh-CN](doc/README-zh-CN.md)

## for what

- git extra diff tool for Excel file
    -  dump excel file metadata to csv file by `go-excelize`

## Features

- [x] support windows macOS linux git diff
- [x] subcommand for diff excel can ignore
    - [x] `--ignore-read` ignore read excel file error
    - [x] `--ignore-parse` ignore excel file to csv error
- [ ] more perfect test case coverage
- [ ] more perfect benchmark case

## usage

- install by release binary from [release page](https://github.com/sinlov/git-extra-diff-excel/releases)

- install binary by cli

```bash
# install at ${GOPATH}/bin
$ go install -v github.com/sinlov/git-extra-diff-excel/cmd/git-extra-diff-excel@latest
# install version v1.0.0
$ go install -v github.com/sinlov/git-extra-diff-excel/cmd/git-extra-diff-excel@v1.0.0
```

remember binary install path, if you use `go install` in `$GOPATH/bin`

- add git diff excel config

```bash
# replace <git-extra-diff-excel> full path for installation
git config --global diff."excel".textconv "<git-extra-diff-excel> csv"
# go install can use as
# linux or macOS
git config --global diff."excel".textconv "$(which git-extra-diff-excel) csv"
# windows powershell
git config --global diff."excel".textconv "$((Get-Command git-extra-diff-excel).Source.Replace(`"\`", `"/`")) csv"

# add binary config
git config --global diff."excel".binary true
```

at file `~/.gitconfig` will add config like `/Users/sinlov/go/bin/git-extra-diff-excel` is install path

```conf
[diff "excel"]
	textconv = /Users/sinlov/go/bin/git-extra-diff-excel csv
	binary = true
```

- Windows config as

```conf
[diff "excel"]
	textconv = C:/Users/sinlov/go/bin/git-extra-diff-excel.exe csv
	binary = true
```

### project excel diff config

- at want to use Excel diff, project root file `.gitattributes` add config

```conf
*.xlsx  diff=excel
*.XLSX  diff=excel
*.xls   diff=excel
*.XLS   diff=excel
```

- then submit `.gitattributes` to take effect

- If you change the Excel file again, you can use the command to see the diff changes.

```bash
$ git diff --word-diff=color --unified=1
─────────────────────────────────────────────────────────────┐
• DataTables/demo/Datas/__tables__.xlsx:7: SheetName: Sheet1 │
─────────────────────────────────────────────────────────────┘
,placeholder.TbItem3,item3,TRUE,item_3.xlsx
,placeholder.TbItem4,item4,TRUE,item_3.xlsx

# or
$ git diff --word-diff=color --cached
─────────────────────────────────────────────────────────────┐
• DataTables/demo/Datas/__tables__.xlsx:7: SheetName: Sheet1 │
─────────────────────────────────────────────────────────────┘
,placeholder.TbItem3,item3,TRUE,item_3.xlsx
,placeholder.TbItem4,item4,TRUE,item_3.xlsx
```

or use [sourceTree](https://www.sourcetreeapp.com/) or [fork](https://git-fork.com/) will show change content

![img.png](https://github.com/sinlov/git-extra-diff-excel/blob/main/doc/img/fork-diff-excel.png?raw=true)

## dev

- see [dev.md](doc-dev/dev.md)

## Contributing

[![Contributor Covenant](https://img.shields.io/badge/contributor%20covenant-v1.4-ff69b4.svg)](.github/CONTRIBUTING_DOC/CODE_OF_CONDUCT.md)
[![GitHub contributors](https://img.shields.io/github/contributors/sinlov/git-extra-diff-excel)](https://github.com/sinlov/git-extra-diff-excel/graphs/contributors)

We welcome community contributions to this project.

Please read [Contributor Guide](.github/CONTRIBUTING_DOC/CONTRIBUTING.md) for more information on how to get started.

请阅读有关 [贡献者指南](.github/CONTRIBUTING_DOC/zh-CN/CONTRIBUTING.md) 以获取更多如何入门的信息