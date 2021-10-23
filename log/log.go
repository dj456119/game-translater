/*
 * @Descripttion:Log文件
 * @version:
 * @Author: cm.d
 * @Date: 2021-10-16 16:13:36
 * @LastEditors: cm.d
 * @LastEditTime: 2021-10-24 01:47:27
 */

package log

import (
	"os"
	"time"

	"github.com/dj456119/game-translater/config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

func Init() {
	logFile := "./applog/game-translater.log"
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	writer, _ := rotatelogs.New(
		logFile+".%Y%m%d",
		rotatelogs.WithLinkName(logFile),
		rotatelogs.WithMaxAge(time.Duration(72)*time.Hour),
	)
	logrus.SetOutput(writer)
	logLevel := config.Config.LogLevel
	switch logLevel {
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}

	logrus.Debug("日志模块启动,加载到的类型为", logLevel)
}
