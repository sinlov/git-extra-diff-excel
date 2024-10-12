## 特性

- [x] 支持 Windows macOS linux git diff
- [x] diff excel 子命令 csv 可以忽略
    - [x] `--ignore-read` 忽略读取excel文件错误
    - [x] `--ignore-parse` 忽略excel文件到csv错误
- [ ] more perfect test case coverage
- [ ] more perfect benchmark case

## 用法

- 通过从发布二进制文件进行安装 [release page](https://github.com/sinlov/git-extra-diff-excel/releases)

- 通过 cli 安装二进制文件

```bash
# 安装到 ${GOPATH}/bin
$ go install -v github.com/sinlov/git-extra-diff-excel/cmd/git-extra-diff-excel@latest
# 安装版本 v1.0.0
$ go install -v github.com/sinlov/git-extra-diff-excel/cmd/git-extra-diff-excel@v1.0.0
```

记住二进制安装路径，如果你使用 `go install` in `$GOPATH/bin`

- 添加 git diff excel 配置

```bash
# 更换 <git-extra-diff-excel> 安装的完整路径
git config --global diff."excel".textconv "<git-extra-diff-excel> csv"
# go install 可以这么使用
# linux or macOS
git config --global diff."excel".textconv "$(which git-extra-diff-excel) csv"
# windows powershell
git config --global diff."excel".textconv "$((Get-Command git-extra-diff-excel).Source.Replace(`"\`", `"/`")) csv"

# 添加 binary 配置
git config --global diff."excel".binary true
```

结果是在 文件 `~/.gitconfig` 将添加配置，其中 `/Users/sinlov/go/bin/git-extra-diff-excel` 是安装路径

```conf
[diff "excel"]
	textconv = /Users/sinlov/go/bin/git-extra-diff-excel csv
	binary = true
```

- Windows配置为

```conf
[diff "excel"]
	textconv = C:/Users/sinlov/go/bin/git-extra-diff-excel.exe csv
	binary = true
```

### 项目中 excel diff 配置

- 在想要使用Excel diff，项目根文件 `.gitattributes` 添加配置

```conf
*.xlsx  diff=excel
*.XLSX  diff=excel
*.xls   diff=excel
*.XLS   diff=excel
```

- 然后提交 `.gitattributes` 生效

- 再改动 Excel 文件, 使用命令可以看到 diff 改动内容

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

或使用 [sourceTree](https://www.sourcetreeapp.com/) or [fork](https://git-fork.com/) 将显示更改内容

![img.png](https://github.com/sinlov/git-extra-diff-excel/blob/main/doc/img/fork-diff-excel.png?raw=true)