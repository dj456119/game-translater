/*
 * @Descripttion:通过kbinani/screenshot截图
 * @version:
 * @Author: cm.d
 * @Date: 2021-10-16 15:51:39
 * @LastEditors: cm.d
 * @LastEditTime: 2021-10-16 18:29:26
 */
package capture

import (
	"bytes"
	"image"
	"image/png"

	"github.com/dj456119/game-translater/capture"
	"github.com/kbinani/screenshot"
	"github.com/sirupsen/logrus"
)

const (
	DefaultScreenNumber = 0
)

type SampleGTCapture struct {
}

func (sgtc SampleGTCapture) Capture(gTCaptureModel *capture.GTCaptureModel) error {
	logrus.Debug("局部截图，获取到的截图信息", *gTCaptureModel)
	var image *image.RGBA
	var err error
	image, err = screenshot.Capture(gTCaptureModel.X, gTCaptureModel.Y, gTCaptureModel.Width, gTCaptureModel.Height)
	if err != nil {
		return err
	}
	logrus.Debug("截图成功,图像大小", image.Rect)
	buffer := new(bytes.Buffer)
	err = png.Encode(buffer, image)
	if err != nil {
		return err
	}
	gTCaptureModel.Image = buffer.Bytes()
	return nil
}

func (sgtc SampleGTCapture) CaptureFullScreen(gTCaptureModel *capture.GTCaptureModel) error {
	logrus.Debug("全局截图，获取到的截图信息", *gTCaptureModel)
	var image *image.RGBA
	var err error
	image, err = screenshot.CaptureDisplay(DefaultScreenNumber)
	if err != nil {
		return err
	}

	logrus.Debug("截图成功,图像大小", image.Rect)

	buffer := new(bytes.Buffer)
	err = png.Encode(buffer, image)

	if err != nil {
		return err
	}

	//填充截图数据结构体
	gTCaptureModel.X = 0
	gTCaptureModel.Y = 0
	gTCaptureModel.Width = image.Rect.Dx()
	gTCaptureModel.Height = image.Rect.Dy()
	gTCaptureModel.Image = buffer.Bytes()
	logrus.Debug("图像的长宽为", gTCaptureModel.Width, gTCaptureModel.Height)
	return nil
}
