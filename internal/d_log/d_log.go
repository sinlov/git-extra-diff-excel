package d_log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gookit/color"
	"log"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

const DefaultExtLogLineMaxDeep = 3

const (
	selfCallerFileMark      = "d_log/d_log.go"
	findExtLogLineCallStart = 2
	findExtLogLineCallDeep  = 4
)

func init() {
	changeLogOutput()
}

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(string(bytes))
}

var extLogLineMaxDeep = DefaultExtLogLineMaxDeep

var (
	openDebug  = false
	showLineNo = false
)

func OpenDebug() {
	openDebug = true
	ShowLogLineNo(true)
	changeLogOutput()
}

func ShowLogLineNo(openLine bool) {
	showLineNo = openLine
}

func SetLogLineDeep(deep uint) {
	if deep < 1 {
		deep = 1
	}
	extLogLineMaxDeep = int(deep)
}

func changeLogOutput() {
	if openDebug {
		log.SetFlags(log.LstdFlags)
		log.SetOutput(log.Default().Writer())
	} else {
		log.SetFlags(0)
		log.SetOutput(new(logWriter))
	}
}

func formatLog(msg string) string {
	if showLineNo {
		return extLogLine(msg)
	}
	return msg
}

func extLogLine(logContent string) string {
	srcFile := ""
	for i := findExtLogLineCallStart; i <= findExtLogLineCallDeep; i++ {
		pc, file, line, ok := runtime.Caller(i)

		if strings.HasSuffix(file, selfCallerFileMark) {
			continue
		}
		if ok {
			funcFullName := runtime.FuncForPC(pc).Name()
			funcName := filepath.Base(funcFullName)
			splitFuncName := strings.Split(funcName, ".")
			splitFuncLen := len(splitFuncName)
			funNameTemp := ""
			if splitFuncLen > 1 {
				for s := 1; s < splitFuncLen; s++ {
					funNameTemp = funNameTemp + "." + splitFuncName[s]
				}
			}
			funcName = strings.TrimPrefix(funNameTemp, ".")
			indexFunc := func(file string) string {
				// depth: extLogLineMaxDeep
				backup := "/" + file

				cacheIndex := 0
				findCache := backup
				for i := 0; i < extLogLineMaxDeep; i++ {
					findIndex := strings.LastIndex(findCache, "/")
					if findIndex < 0 {
						return backup[cacheIndex+1 : 0]
					}
					cacheIndex = findIndex
					findCache = findCache[:findIndex]
				}
				return backup[cacheIndex+1:]
			}
			srcFile = fmt.Sprintf("%s:%s %s() [g:%d]", indexFunc(file), strconv.Itoa(line), funcName, getRuntimeGid())
		}
		break
	}
	return fmt.Sprintf("%s %s", srcFile, logContent)
}

func getRuntimeGid() (gid uint64) {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		panic(err)
	}
	return n
}

func Debug(msg string) {
	if !openDebug {
		return
	}
	log.Printf("%s %s", color.Blue.Render("debug:"), formatLog(msg))
}

func Debugf(format string, v ...any) {
	if !openDebug {
		return
	}
	logInfo := fmt.Sprintf(format, v...)
	log.Printf("%s %s", color.Blue.Render("debug:"), formatLog(logInfo))
}

func DebugJson(d any) {
	DebugJsonf(d, "")
}

func DebugJsonf(d any, format string, v ...any) {
	if !openDebug {
		return
	}
	dStr, done := any2JsonBeauty(d)
	if done {
		return
	}
	if format != "" {
		logInfo := fmt.Sprintf(format, v...)
		log.Printf("%s %s\n%s\n", color.Blue.Render("debug:"), formatLog(logInfo), dStr)
	} else {
		log.Printf("%s%s\n", color.Blue.Render("debug:"), formatLog(dStr))
	}
}

func Verbosef(format string, v ...any) {
	logInfo := fmt.Sprintf(format, v...)
	if showLineNo {
		logInfo = extLogLine(logInfo)
	}
	log.Print(logInfo)
}

func VerboseJson(d any) {
	VerboseJsonf(d, "")
}

func VerboseJsonf(d any, format string, v ...any) {
	dStr, done := any2JsonBeauty(d)
	if done {
		return
	}
	if format != "" {
		logInfo := fmt.Sprintf(format, v...)
		log.Printf("%s\n%s\n", formatLog(logInfo), dStr)
	} else {
		log.Printf("%s\n", formatLog(dStr))
	}
}

func Info(msg string) {
	log.Printf("%s %s", color.Green.Render("info:"), formatLog(msg))
}

func Infof(format string, v ...any) {
	logInfo := fmt.Sprintf(format, v...)
	log.Printf("%s %s", color.Green.Render("info:"), formatLog(logInfo))
}

func Warnf(format string, v ...any) {
	logInfo := fmt.Sprintf(format, v...)
	log.Printf("%s %s", color.Yellow.Render("warn:"), formatLog(logInfo))
}

func Error(err error) {
	Errorf(err, "")
}

func Errorf(err error, format string, v ...any) {
	if format != "" {
		logInfo := fmt.Sprintf(format, v...)
		log.Printf("%s %s\nerror content: %v", color.Red.Render("err:"), formatLog(logInfo), err)
	} else {
		log.Printf("%s %s\n", color.Red.Render("err:"), formatLog(err.Error()))
	}
}

func Fatalf(format string, v ...any) {
	logInfo := fmt.Sprintf(format, v...)
	log.Fatalf("fatal: %s", formatLog(logInfo))
}

func Panicf(format string, v ...any) {
	logInfo := fmt.Sprintf(format, v...)
	log.Panicf("panic: %s", formatLog(logInfo))
}

func any2JsonBeauty(d any) (string, bool) {
	data, errJson := json.Marshal(d)
	if errJson != nil {
		log.Printf("%s %s %v", color.Yellow.Render("warn:"), "json marshal err: ", errJson)
		return "", true
	}
	var str bytes.Buffer
	errBeauty := json.Indent(&str, data, "", "  ")
	if errBeauty != nil {
		log.Printf("%s %s %v", color.Yellow.Render("warn:"), "json marshal beauty err: ", errJson)
		return "", true
	}
	dStr := str.String()
	return dStr, false
}
