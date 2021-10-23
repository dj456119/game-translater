/*
 * @Descripttion: 主函数入口
 * @version:
 * @Author: cm.d
 * @Date: 2021-10-11 22:13:29
 * @LastEditors: cm.d
 * @LastEditTime: 2021-10-24 00:55:16
 */
package main

import (
	"github.com/dj456119/game-translater/config"
	"github.com/dj456119/game-translater/log"
	govclview "github.com/dj456119/game-translater/view/govcl-view"
	"github.com/sirupsen/logrus"
)

func init() {
	config.Init()
	log.Init()
}

func main() {
	logrus.Info("游戏翻译机启动")
	govclview.Init()
}
