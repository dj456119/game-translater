/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-10-16 19:30:10
 * @LastEditors: cm.d
 * @LastEditTime: 2021-10-16 21:08:09
 */
package baidutranslater

import (
	"fmt"

	"github.com/dj456119/game-translater/translater"
	"github.com/jinzhu/configor"
	baidutranslate "github.com/shenjinti/baidu_translate_go"
	"github.com/sirupsen/logrus"
)

type BaiduTranslaterConfig struct {
	APPID string `required:"true"`
	Key   string `required:"true"`
}

const (
	Address                  = "https://fanyi-api.baidu.com/api/trans/vip/translate"
	ContentType              = "application/json; charset=UTF-8"
	TranslaterConnectTimeout = 10
	TranslaterSendTimeout    = 300
)

type BaiduTranslater struct {
	Config         BaiduTranslaterConfig
	BaiduTranslate *baidutranslate.BaiduTranslate
}

func (bt BaiduTranslater) Translate(gtModel *translater.GTranslaterModel) error {
	gtModel.Translated = make([]string, 1)
	wordString := ""
	var err error
	for _, word := range gtModel.Words {

		wordString = fmt.Sprintf("%s%s", wordString, word)
	}
	gtModel.Words = []string{wordString}
	gtModel.Translated[0], err = bt.BaiduTranslate.Text("en", "zh", wordString)
	logrus.Debug("翻译成功，结果:", gtModel.Translated[0])
	return err
}

func NewBaiduTranslater() *BaiduTranslater {
	bt := new(BaiduTranslater)
	btConfig := new(BaiduTranslaterConfig)
	configor.Load(&btConfig, "baidu-translater-config.yaml")
	bt.Config = *btConfig
	bt.BaiduTranslate = baidutranslate.NewBaiduTranslate(bt.Config.APPID, bt.Config.Key)
	logrus.Info("读取百度云翻译配置成功，", bt.Config)
	return bt
}
