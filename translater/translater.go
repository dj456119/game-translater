/*
 * @Descripttion:翻译核心接口
 * @version:
 * @Author: cm.d
 * @Date: 2021-10-15 20:47:21
 * @LastEditors: cm.d
 * @LastEditTime: 2021-10-15 21:15:03
 */

package translater

type GTranslaterModel struct {
	Words      []string //原始文字
	Translated []string //翻译后文字
}

type GTranslater interface {
	/**
	 * @name: Translate
	 * @msg: 翻译函数
	 * @param {*GTranslaterModel} 翻译对象
	 * @return {error} 如果翻译失败，返回错误，否则返回nil
	 */
	Translate(*GTranslaterModel) error
}
