/*
 * @Descripttion:控制层
 * @version:
 * @Author: cm.d
 * @Date: 2021-10-15 20:16:57
 * @LastEditors: cm.d
 * @LastEditTime: 2021-10-16 18:37:10
 */
package controller

import (
	"github.com/dj456119/game-translater/capture"
	samplecapture "github.com/dj456119/game-translater/capture/sample-capture"
	"github.com/dj456119/game-translater/ocr"
	aliyunocr "github.com/dj456119/game-translater/ocr/aliyun-ocr"
	"github.com/dj456119/game-translater/translater"
	baidutranslater "github.com/dj456119/game-translater/translater/baidu-translater"
)

type GTResponse struct {
	Status int
	Data   *GTResponseData
	Msg    string
}

type GTResponseData struct {
	Words      []string
	Translated []string
}

const (
	GTResponseStatusSuccess = 200
	GTResponseStatusErr     = -1
)

type GTRequest struct {
	X      int
	Y      int
	Width  int
	Height int
}

type GTController struct {
	GTCapture   capture.GTCapture
	GTOCR       ocr.GTOCR
	GTranslater translater.GTranslater
}

/**
 * @name:CaptureScreenAndTranslate
 * @msg:截屏并且翻译
 * @param {*GTRequest} req 请求参数
 * @return {*GTResponse} 响应体
 */
func (gtc *GTController) CaptureScreenAndTranslate(req *GTRequest) *GTResponse {
	//截图
	gTCaptureModel := new(capture.GTCaptureModel)
	gTCaptureModel.X = req.X
	gTCaptureModel.Y = req.Y
	gTCaptureModel.Width = req.Width
	gTCaptureModel.Height = req.Height
	err := gtc.GTCapture.Capture(gTCaptureModel)
	if err != nil {
		return CreateErrResponse(err)
	}

	//OCR识别文字
	gTOCRModel := new(ocr.GTOCRModel)
	gTOCRModel.Image = gTCaptureModel.Image
	err = gtc.GTOCR.OCR(gTOCRModel)
	if err != nil {
		return CreateErrResponse(err)
	}

	//翻译
	gTranslater := new(translater.GTranslaterModel)
	gTranslater.Words = gTOCRModel.Words
	err = gtc.GTranslater.Translate(gTranslater)
	if err != nil {
		return CreateErrResponse(err)
	}
	return CreateSuccessResponse(gTranslater.Words, gTranslater.Translated)
}

/**
 * @name: CreateErrResponse
 * @msg: 创建失败响应
 * @param {error} err 响应错误对象
 * @return {*GTResponse} 响应错误结构体
 */
func CreateErrResponse(err error) *GTResponse {
	resp := new(GTResponse)
	resp.Status = GTResponseStatusErr
	resp.Msg = err.Error()
	return resp
}

/**
 * @name:CreateSuccessResponse
 * @msg:创建成功响应
 * @param {[]string} words 原始文字
 * @param {[]string} translated 翻译后文字
 * @return {*GTResponse} 响应体
 */
func CreateSuccessResponse(words, translated []string) *GTResponse {
	resp := new(GTResponse)
	respData := new(GTResponseData)
	respData.Words = words
	respData.Translated = translated
	resp.Data = respData
	resp.Status = GTResponseStatusSuccess
	return resp
}

func NewGTController() *GTController {
	controller := new(GTController)
	controller.GTCapture = new(samplecapture.SampleGTCapture)
	controller.GTOCR = aliyunocr.NewAliyunOCR()
	controller.GTranslater = baidutranslater.NewBaiduTranslater()
	return controller
}
