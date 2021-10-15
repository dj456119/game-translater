/*
 * @Descripttion:OCR接口类
 * @version:
 * @Author: cm.d
 * @Date: 2021-10-15 20:40:07
 * @LastEditors: cm.d
 * @LastEditTime: 2021-10-15 20:46:23
 */

package ocr

type GTOCRModel struct {
	Image []byte
	Word  []string
}

type GTOCR interface {
	OCR(*GTOCRModel) error
}
