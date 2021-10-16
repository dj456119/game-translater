/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-10-17 01:21:43
 * @LastEditors: cm.d
 * @LastEditTime: 2021-10-17 01:22:59
 */
package test

import (
	"github.com/dj456119/game-translater/capture"
	samplecapture "github.com/dj456119/game-translater/capture/sample-capture"
	"github.com/dj456119/game-translater/ocr"
	aliyunocr "github.com/dj456119/game-translater/ocr/aliyun-ocr"
	"github.com/dj456119/game-translater/translater"
	baidutranslater "github.com/dj456119/game-translater/translater/baidu-translater"
	"github.com/sirupsen/logrus"
)

func Test() {
	gTCaptureModel := new(capture.GTCaptureModel)
	s := new(samplecapture.SampleGTCapture)
	s.CaptureFullScreen(gTCaptureModel)
	bytes := gTCaptureModel.Image
	gTOCRModel := new(ocr.GTOCRModel)
	gTOCRModel.Image = bytes
	o := aliyunocr.NewAliyunOCR()
	o.OCR(gTOCRModel)
	bt := baidutranslater.NewBaiduTranslater()
	gTranslaterModel := new(translater.GTranslaterModel)
	gTranslaterModel.Words = gTOCRModel.Words
	bt.Translate(gTranslaterModel)
	logrus.Debug(gTranslaterModel.Translated)
}
