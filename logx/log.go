package logx

import (
    "fmt"
    "log"
    "os"
    "strconv"
    "time"
)

// LogPool 日志池
// 系统运行中所有的日志文件需在此注册
var LogPool = make(map[string]*os.File)

// LogConfig 日志配置
type LogConfig struct {
    Path string `yaml:"path"`
}

type Logger struct {
    config *LogConfig
    log    *log.Logger
    sn     string
}

func NewLogger(sn string) *Logger {
    if len(sn) == 0 {
        sn = strconv.FormatInt(time.Now().UnixNano(), 10)
    }

    return &Logger{
        config: &LogConfig{Path: "./"},
        log:    nil,
        sn:     sn,
    }
}

// GetLogger 从日志池中获取日志对象
// 如果日志池中不存在则自动注册进池
func (l *Logger) GetLogger(fileName string) *Logger {
    filePath := fmt.Sprintf("%s/%s.log", l.config.Path, fileName)

    file, ok := LogPool[filePath]
    if ok {
        l.log = log.New(file, "", log.Ldate|log.Lmicroseconds)
        return l
    } else {
        file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
        if err != nil {
            fmt.Println(fmt.Sprintf("[%s] 文件打开失败:%s", time.Now().Format("2006-01-02 15:04:05.000000"), filePath))
            panic(err)
        }
        LogPool[filePath] = file

        l.log = log.New(file, "", log.Ldate|log.Lmicroseconds)
        return l
    }
}

// Info 标准日志
func (l *Logger) Info(v ...interface{}) {
    l.log.Println(append([]interface{}{l.sn, "INFO"}, v...)...)
}

// InfoF 标准日志
// 参数以fmt方式处理
func (l *Logger) InfoF(format string, v ...interface{}) {
    l.log.Println(l.sn, "INFO", fmt.Sprintf(format, v...))
}

// Error 错误日志
func (l *Logger) Error(v ...interface{}) {
    l.log.Println(append([]interface{}{l.sn, "ERROR"}, v...)...)
}

// ErrorF 错误日志
// 参数以fmt方式处理
func (l *Logger) ErrorF(format string, v ...interface{}) {
    l.log.Println(l.sn, "ERROR", fmt.Sprintf(format, v...))
}

// Debug 调试日志
func (l *Logger) Debug(v ...interface{}) {
    l.log.Println(append([]interface{}{l.sn, "DEBUG"}, v...)...)
}

// DebugF 调试日志
// 参数以fmt方式处理
func (l *Logger) DebugF(format string, v ...interface{}) {
    l.log.Println(l.sn, "DEBUG", fmt.Sprintf(format, v...))
}

// SystemLog 标准日志
func (l *Logger) SystemLog(v ...interface{}) {
    l.log.Println(append([]interface{}{l.sn, "SYSTEM"}, v...)...)
}

// SystemLogF 标准日志
// 参数以fmt方式处理
func (l *Logger) SystemLogF(format string, v ...interface{}) {
    l.log.Println(l.sn, "SYSTEM", fmt.Sprintf(format, v...))
}
