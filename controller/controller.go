/*
 * @Descripttion:控制层
 * @version:
 * @Author: cm.d
 * @Date: 2021-10-15 20:16:57
 * @LastEditors: cm.d
 * @LastEditTime: 2021-10-15 20:54:48
 */
package controller

import "github.com/dj456119/game-translater/capture"

type GTResponse struct {
	Words      []string
	Translated []string
}

type GTRequest struct {
	X      int
	Y      int
	Width  int
	Height int
}

type GTController struct {
	GTCapture capture.GTCapture
	GTOCR     capture.GTOCR
}

func (gtc *GTController) CaptureScreenAndTranslate(req *GTRequest) (*GTResponse, error) {
	gtm, err := gtc.GTCapture.Capture(req.X, req.Y, req.Width, req.Height)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
