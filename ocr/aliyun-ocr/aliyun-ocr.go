/*
 * @Descripttion:基于阿里云的OCR
 * @version:
 * @Author: cm.d
 * @Date: 2021-10-16 16:46:21
 * @LastEditors: cm.d
 * @LastEditTime: 2021-10-24 02:06:16
 */

package aliyunocr

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/dj456119/game-translater/ocr"
	"github.com/jinzhu/configor"
	"github.com/sirupsen/logrus"
)

type AliyunOCRConfig struct {
	APPCODE string `required:"true"`
}

type OCRRequestBody struct {
	Image     string                   `json:"image"`
	Configure *OCRRequestBodyConfigure `json:"configure"`
}

type OCRRequestBodyConfigure struct {
	MinSize                    int  `json:"min_size"`
	OutputProb                 bool `json:"output_prob"`
	OutputKeypoints            bool `json:"output_keypoints"`
	SkipDetection              bool `json:"skip_detection"`
	WithoutPredictingDirection bool `json:"without_predicting_direction"`
}

type OCRResponseBody struct {
	RequestId           string                `json:"request_id"`
	OCRResponseBodyRets []*OCRResponseBodyRet `json:"ret"`
}

type OCRResponseBodyRet struct {
	//Prob float32 `json:"prob"`
	Word string `json:"word"`
}

func CreateOCRRequestBody(imageBytes []byte) *OCRRequestBody {
	orb := new(OCRRequestBody)
	orbc := new(OCRRequestBodyConfigure)
	orbc.MinSize = 16
	orbc.OutputProb = true
	orbc.OutputKeypoints = false
	orbc.SkipDetection = false
	orbc.WithoutPredictingDirection = false
	orb.Image = base64.StdEncoding.EncodeToString(imageBytes)
	orb.Configure = orbc
	return orb
}

const (
	ContentType       = "application/json; charset=UTF-8"
	OCRConnectTimeout = 10
	OCRSendTimeout    = 300
	OCRRequestAddress = "https://tysbgpu.market.alicloudapi.com/api/predict/ocr_general"
)

func (aliyunOCR *AliyunOCR) SendToAiliyunOCR(imageBytes []byte) (*OCRResponseBody, error) {
	client := http.Client{Transport: &http.Transport{
		Dial: func(netw, addr string) (net.Conn, error) {
			c, err := net.DialTimeout(netw, addr, time.Second*OCRConnectTimeout)
			if err != nil {
				return nil, err
			}
			return c, nil

		},
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time.Second * OCRSendTimeout,
	}}

	orb := CreateOCRRequestBody(imageBytes)
	byte, err := json.Marshal(orb)

	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", OCRRequestAddress, bytes.NewBuffer(byte))

	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", ContentType)
	req.Header.Set("Authorization", aliyunOCR.Config.APPCODE)
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	orespb := new(OCRResponseBody)

	json.Unmarshal(body, orespb)
	return orespb, err
}

type AliyunOCR struct {
	Config AliyunOCRConfig
}

func NewAliyunOCR() *AliyunOCR {
	aliyunOCR := new(AliyunOCR)
	aliyunOCR.Config = AliyunOCRConfig{}
	err := configor.Load(&aliyunOCR.Config, "aliyun-ocr-config.yaml")
	if err != nil {
		logrus.Fatal("读取阿里云ocr配置失败", err)
	}
	logrus.Info("读取阿里云ocr配置成功", aliyunOCR.Config)
	return aliyunOCR
}

func (aocr *AliyunOCR) OCR(gTOCRModel *ocr.GTOCRModel) error {
	resp, err := aocr.SendToAiliyunOCR(gTOCRModel.Image)
	if err != nil {
		return err
	}
	gTOCRModel.Words = make([]string, len(resp.OCRResponseBodyRets))
	for i, rect := range resp.OCRResponseBodyRets {
		gTOCRModel.Words[i] = rect.Word
	}
	logrus.Debug("OCR识别成功，结果为, ", gTOCRModel.Words)
	return nil
}
