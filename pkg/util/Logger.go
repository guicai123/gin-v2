package util

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path"
	"time"
)

// https://www.cnblogs.com/taoshihan/p/13846256.html
// gorm 的mysql日志进行记录
func Logger() *logrus.Logger {
	var LogInstance = logrus.New()
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/runtime/sql_logs/"
	}
	logFileName := now.Format("2006-01-02") + ".log"
	// 打开文件
	logFileName1 := path.Join(logFilePath, logFileName)
	// 使用滚动压缩方式记录日志
	// 设置输出
	LogInstance.SetOutput(&lumberjack.Logger{
		LocalTime:  true,         //是否使用本地时间来命名备份的日志
		Filename:   logFileName1, //日志文件位置
		MaxSize:    1,            // 单文件最大容量,单位是MB
		MaxBackups: 3,            // 最大保留过期文件个数 默认是0，不会清除文件
		MaxAge:     30,           // 保留过期文件的最大时间间隔,单位是天 默认是0，不会清除文件
		Compress:   true,         // 是否需要压缩滚动日志, 使用的 gzip 压缩
	})
	// 设置日志输出JSON格式
	//LogInstance.SetFormatter(&logrus.TextFormatter{})
	LogInstance.SetFormatter(&logrus.JSONFormatter{})
	// 设置日志记录级别
	LogInstance.SetLevel(logrus.DebugLevel)
	return LogInstance

}
