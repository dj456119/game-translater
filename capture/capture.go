/*
 * @Descripttion:截图接口及相关Model
 * @version:
 * @Author: cm.d
 * @Date: 2021-10-14 22:07:07
 * @LastEditors: cm.d
 * @LastEditTime: 2021-10-14 22:18:12
 */

package capture

type GTCaptureModel struct {
	X      int    //图片位于屏幕坐标X
	Y      int    //图片位于屏幕坐标Y
	Length int    //图片宽
	Height int    //图片高
	Image  []byte //图片数据byte数组
}

type GTCapture interface {

	/**
	* @name: CaptureFullScreen
	* @msg: 全局截屏
	* @return 截图结构体
	 */
	CaptureFullScreen() (*GTCaptureModel, error)

	/**
	 * @name: Capture
	 * @msg: 根据给定坐标及宽高截图
	 * @param 给定坐标、宽、高
	 * @return 截图结构体
	 */
	Capture(x, y, length, height int) (*GTCaptureModel, error)
}
