# 游戏翻译机

## 一句话的介绍

一个跨平台支持macos与windows的、基于golang开发的游戏翻译机，对接了云上的OCR与翻译能力。

## 支持列表
|支持列表|是否完成|
|:----|:----:|
|支持英文=>中文|是|
|支持Windows 64|是|
|支持MacOS amd64、M1|是|
|支持阿里云OCR|是|
|支持百度云翻译|是|

## 背景

很多年前的一个春节，回老家后百无聊赖的我突然想玩PS2上的女神异闻录3，由于这个游戏PS2上并没有汉化版，因此英文版就成了唯一的选择（PS.自动忽略日文版）。但是好不容易调配好了模拟器后，望着满眼的英文依然无从下手，上网上搜了个游戏英翻机，结果居然都要收费。思来想去，打算自己动手丰衣足食，自己写一个顺便充实一下我的假期。在经历了3、4天的折腾后，一个基于Tesseract-ocr+google翻译api的家伙就完成了，当时各厂商的saas云还没这么完善，真是自己一点一点拼凑出来的，为了提高ocr的识别率，还做了灰度、滤波等大量的图像处理。最终还是完成了这个程序，只不过我也刚好休假结束了:)。若干年后ocr、翻译的相关云产品已经非常丰富了，并且有很多的免费产品，因此简单的做了一个封装，写了个ui，希望能把之前还缺失的地方补充完善。

## 编译与运行
    注意事项
    1.目前支持在Mac x86(darwin amd64)及m1、windows 64版本运行，暂不支持linux
    2.由于依赖的gui库的原因，目前仅支持在mac x86上编译，在m1版本无法编译通过
    3.win64下编译支持进行中
    4.编译脚本采用了mage，如果想要自行编译，可能需要先安装mage
在mac上执行
```
git clone https://github.com/dj456119/game-translater
cd game-translater
```
需要生成mac的可运行程序:
```
mage build macos
```
需要生成windows的可运行程序
```
mage build windows
```
编译完成相关的程序就已经在target目录中了  
macos下启动
```
cd target
chmod +x game-translater
./game-translater
```
windows下直接进入目录执行game-translater.exe即可  
如果没有其他错误，就可以顺利使用了

## 使用说明

待补充

## Todo List

|待支持列表|是否完成|
|:----|:----:|
|支持日语=>中文|否|
|支持tesseract-ocr|否|
|支持Windows|是|
|支持Linux|否|
|制作dmg安装包|否|