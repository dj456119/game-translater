/*
 * @Descripttion:截图接口及相关Model
 * @version:
 * @Author: cm.d
 * @Date: 2021-10-14 22:07:07
 * @LastEditors: cm.d
 * @LastEditTime: 2021-10-15 21:18:44
 */

package capture

type GTCaptureModel struct {
	X      int    //图片位于屏幕坐标X
	Y      int    //图片位于屏幕坐标Y
	Width  int    //图片宽
	Height int    //图片高
	Image  []byte //图片数据byte数组
}

type GTCapture interface {

	/**
	 * @name: CaptureFullScreen
	 * @msg: 全屏截图
	 * @param {*GTCaptureModel} 截屏描述结构体
	 * @return {error} 截图结构体，截图失败返回error, 否则返回nil
	 */
	CaptureFullScreen(GTCaptureModel) error

	/**
	 * @name:
	 * @msg:
	 * @param {*GTCaptureModel} 截屏描述结构体
	 * @return {error} 截图结构体，截图失败返回error, 否则返回nil
	 */
	Capture(*GTCaptureModel) error
}
