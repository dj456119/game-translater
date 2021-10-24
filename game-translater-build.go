//+build mage

/*
 * @Descripttion:mage打包脚本
 * @version:
 * @Author: cm.d
 * @Date: 2021-10-24 01:00:19
 * @LastEditors: cm.d
 * @LastEditTime: 2021-10-24 12:09:00
 */
package main

import (
	"github.com/magefile/mage/sh"
	"github.com/sirupsen/logrus"
)

func Build(os string) error {

	switch os {
	case "windows":
		return buildWindows()
	case "macos":
		return buildMac()
	default:
		logrus.Fatal("不支持的操作系统: ", os)
	}

	return nil
}

func buildMac() error {
	env := make(map[string]string)
	env["GOOS"] = "darwin"
	env["GOARCH"] = "amd64"
	logrus.Info("清理target目录")
	sh.Run("rm", "-rf", "target")
	logrus.Info("创建target目录")
	if err := sh.Run("mkdir", "target"); err != nil {
		return err
	}
	logrus.Info("编译程序中")
	if err := sh.RunWith(env, "go", "build"); err != nil {
		return err
	}
	logrus.Info("拷贝本地库")
	if err := sh.Run("cp", "-rf", "./liblcl.dylib", "./target"); err != nil {
		return err
	}
	logrus.Info("拷贝程序")
	if err := sh.Run("cp", "-rf", "./game-translater", "./target"); err != nil {
		return err
	}
	logrus.Info("拷贝配置文件")
	if err := sh.Run("cp", "-rf", "aliyun-ocr-config.yaml", "./target"); err != nil {
		return err
	}
	if err := sh.Run("cp", "-rf", "config.yaml", "./target"); err != nil {
		return err
	}
	if err := sh.Run("cp", "-rf", "baidu-translater-config.yaml", "./target"); err != nil {
		return err
	}
	logrus.Info("创建日志目录")
	if err := sh.Run("mkdir", "./target/applog"); err != nil {
		return err
	}
	logrus.Info("清理中间文件")
	if err := sh.Run("rm", "-rf", "game-translater"); err != nil {
		return err
	}
	logrus.Info("已完成")
	return nil
}

func buildWindows() error {
	env := make(map[string]string)
	env["GOOS"] = "windows"
	env["GOARCH"] = "amd64"
	logrus.Info("清理target目录")
	sh.Run("rm", "-rf", "target")
	logrus.Info("创建target目录")
	if err := sh.Run("mkdir", "target"); err != nil {
		return err
	}
	logrus.Info("编译程序中")
	if err := sh.RunWith(env, "go", "build"); err != nil {
		return err
	}
	logrus.Info("拷贝本地库")

	if err := sh.Run("cp", "-rf", "./liblcl.dll", "./target"); err != nil {
		return err
	}
	logrus.Info("拷贝程序")
	if err := sh.Run("cp", "-rf", "./game-translater.exe", "./target"); err != nil {
		return err
	}
	logrus.Info("拷贝配置文件")
	if err := sh.Run("cp", "-rf", "aliyun-ocr-config.yaml", "./target"); err != nil {
		return err
	}
	if err := sh.Run("cp", "-rf", "config.yaml", "./target"); err != nil {
		return err
	}
	if err := sh.Run("cp", "-rf", "baidu-translater-config.yaml", "./target"); err != nil {
		return err
	}
	logrus.Info("创建日志目录")
	if err := sh.Run("mkdir", "./target/applog"); err != nil {
		return err
	}
	logrus.Info("清理中间文件")
	if err := sh.Run("rm", "-rf", "game-translater.exe"); err != nil {
		return err
	}
	logrus.Info("已完成")
	return nil
}
