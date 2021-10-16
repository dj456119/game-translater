/*
 * @Descripttion: 主函数入口
 * @version:
 * @Author: cm.d
 * @Date: 2021-10-11 22:13:29
 * @LastEditors: cm.d
 * @LastEditTime: 2021-10-17 01:23:42
 */
package main

import (
	"github.com/dj456119/game-translater/config"
	"github.com/dj456119/game-translater/log"
	"github.com/dj456119/game-translater/test"
	"github.com/sirupsen/logrus"
)

func init() {
	config.Init()
	log.Init()
}

func main() {
	logrus.Info("游戏英翻机启动")
	test.Test()
}
