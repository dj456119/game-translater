/*
 * @Descripttion:OCR接口类
 * @version:
 * @Author: cm.d
 * @Date: 2021-10-15 20:40:07
 * @LastEditors: cm.d
 * @LastEditTime: 2021-10-15 21:14:08
 */

package ocr

type GTOCRModel struct {
	Image []byte   //图像byte数组
	Word  []string //识别出的文字
}

type GTOCR interface {
	/**
	 * @name:OCR
	 * @msg:识别图片中的文字
	 * @param {*GTOCRModel} 图片结构体
	 * @return {error} 识别失败，返回error，否则返回nil
	 */
	OCR(*GTOCRModel) error
}
