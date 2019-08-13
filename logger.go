package log


type LEVEL int32
type UNIT int64
type _ROLLTYPE int //dailyRolling ,rollingFile

const _DATEFORMAT = "2006-01-02 15"

var logLevel LEVEL = 1

const (
	_       = iota
	KB UNIT = 1 << (iota * 10)
	MB
	GB
	TB
)

const (
	ALL LEVEL = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
	OFF
)

const (
	_DAILY _ROLLTYPE = iota
	_ROLLFILE
)

func Init(servername string) {
	setLevel(ALL)
	SetRollingDaily("../../log/",servername)
	Info(servername," 启动成功")
}

func setConsole(isConsole bool) {
	defaultlog.setConsole(isConsole)
}

func setLevel(_level LEVEL) {
	defaultlog.setLevel(_level)
}

func setFormat(logFormat string) {
	defaultlog.setFormat(logFormat)
}

func SetRollingDaily(fileDir, fileName string) {
	defaultlog.setRollingDaily(fileDir, fileName)
}

func Debug(v...interface{}) {
	defaultlog.debug(v...)
}

func Info(v...interface{}) {
	defaultlog.info(v...)
}
func Warn(v...interface{}) {
	defaultlog.warn(v...)
}
func Error(v...interface{}) {
	
	defaultlog.error(v...)
}
func Fatal(v...interface{}) {

	defaultlog.fatal(v...)
}

func SetLevelFile(level LEVEL, dir, fileName string) {
	defaultlog.setLevelFile(level, dir, fileName)
}
