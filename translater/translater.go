/*
 * @Descripttion:翻译核心接口
 * @version:
 * @Author: cm.d
 * @Date: 2021-10-15 20:47:21
 * @LastEditors: cm.d
 * @LastEditTime: 2021-10-15 20:51:49
 */

package translater

type GTranslaterModel struct {
	Word       []string
	Translated []string
}

type GTranslater interface {
	Translate(*GTranslaterModel) error
}
