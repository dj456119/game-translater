package govclview

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"math"

	"github.com/nfnt/resize"

	"github.com/dj456119/game-translater/capture"
	"github.com/dj456119/game-translater/controller"
	"github.com/sirupsen/logrus"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
)

const (
	CaptureFormWidth    = 800
	TranslateFormWidth  = 270
	TranslateFormHeight = 400
)

type TranslaterForm struct {
	*vcl.TForm
	Label           *vcl.TLabel  //当前捕捉信息
	Label1          *vcl.TLabel  //原始文字
	Label2          *vcl.TLabel  //翻译文字
	ButtonTranslate *vcl.TButton //翻译
	ButtonCapture   *vcl.TButton //截图按钮
	EditWords       *vcl.TMemo   //原始文字
	EditTranslated  *vcl.TMemo   //翻译文字
	ButtonAbout     *vcl.TButton //关于按钮
	Con             *controller.GTController
	GTRequest       *controller.GTRequest
}

var MyTranslaterForm *TranslaterForm
var MyCaptureForm *CaptureForm
var IsInit bool = false

func Init() {
	vcl.RunApp(&MyTranslaterForm, &MyCaptureForm)
}

func (f *TranslaterForm) OnFormCreate(sender vcl.IObject) {
	con := controller.NewGTController()
	f.Con = con

	request := new(controller.GTRequest)
	f.GTRequest = request

	f.SetShowInTaskBar(types.StAlways)
	f.SetCaption("游戏翻译机v0.0.1")
	f.EnabledMaximize(false)
	f.SetWidth(TranslateFormWidth)
	f.SetHeight(TranslateFormHeight)
	f.ScreenCenter()

	f.Label = vcl.NewLabel(f)
	f.Label.SetParent(f)
	f.Label.SetCaption(f.GetRealArea())
	f.Label.SetLeft(20)
	f.Label.SetBounds(10, 5, 250, 20)

	//设置Label
	f.Label1 = vcl.NewLabel(f)
	f.Label1.SetParent(f)
	f.Label1.SetCaption("实时捕捉的屏幕文字")
	f.Label1.SetLeft(10)
	f.Label1.SetTop(25)

	//设置实时捕捉的文本
	f.EditWords = vcl.NewMemo(f)
	f.EditWords.SetParent(f)
	f.EditWords.SetReadOnly(true)
	f.EditWords.SetText("在这里显示正从屏幕获取的文本")
	f.EditWords.SetBounds(10, 40, 250, 150)

	//设置Labe2
	f.Label2 = vcl.NewLabel(f)
	f.Label2.SetParent(f)
	f.Label2.SetCaption("实时翻译的文字")
	f.Label2.SetLeft(10)
	f.Label2.SetTop(195)

	//设置实时翻译的文本
	f.EditTranslated = vcl.NewMemo(f)
	f.EditTranslated.SetParent(f)
	f.EditTranslated.SetText("在这里显示被翻译的文本")
	f.EditTranslated.SetBounds(10, 210, 250, 150)

	//设置按钮
	f.ButtonTranslate = vcl.NewButton(f)
	f.ButtonTranslate.SetParent(f)
	f.ButtonTranslate.SetCaption("启动翻译")
	f.ButtonTranslate.SetLeft(98)
	f.ButtonTranslate.SetTop(365)
	f.ButtonTranslate.SetOnClick(f.OnButtonTranslateClick)

	f.ButtonCapture = vcl.NewButton(f)
	f.ButtonCapture.SetParent(f)
	f.ButtonCapture.SetCaption("文本区域")
	f.ButtonCapture.SetLeft(10)
	f.ButtonCapture.SetTop(365)
	f.ButtonCapture.SetOnClick(f.OnButtonCaptureClick)

	f.ButtonAbout = vcl.NewButton(f)
	f.ButtonAbout.SetParent(f)
	f.ButtonAbout.SetCaption("关于")
	f.ButtonAbout.SetLeft(185)
	f.ButtonAbout.SetTop(365)
	f.ButtonAbout.SetOnClick(func(sender vcl.IObject) {
		vcl.ShowMessage("版本v0.0.1，作者cm.d")
	})
}

func (f *TranslaterForm) OnFormClose(sender vcl.IObject, action *types.TCloseAction) {
	vcl.Application.Terminate()
}

func (f *TranslaterForm) OnButtonTranslateClick(sender vcl.IObject) {
	if !IsInit {
		vcl.ShowMessage("您还没设置文字捕捉区域")
		return
	}
	words := f.EditWords.Text()

	logrus.Debug("输入的翻译文字", words)
	logrus.Debug("触发翻译，翻译区间为", f.GTRequest)
	resp := f.Con.CaptureScreenAndTranslate(f.GTRequest)
	if len(resp.Data.Words) != 0 {
		f.EditWords.SetText(resp.Data.Words[0])
	}

	if len(resp.Data.Translated) != 0 {
		f.EditTranslated.SetText(resp.Data.Translated[0])
	}

}

func (f *TranslaterForm) OnButtonCaptureClick(sender vcl.IObject) {
	logrus.Debug("开始重新截图")
	gtcm := new(capture.GTCaptureModel)
	MyTranslaterForm.Con.GTCapture.CaptureFullScreen(gtcm)
	MyCaptureForm.OriImage = vcl.NewPngImage()
	MyCaptureForm.OriImage.LoadFromBytes(gtcm.Image)
	imageBuffer, _ := ScaleImage(CaptureFormWidth, gtcm.Image)
	MyCaptureForm.Image.LoadFromBytes(imageBuffer)
	MyCaptureForm.Canvas().Draw(0, 0, MyCaptureForm.Image)
	MyCaptureForm.Show()
}

type CaptureForm struct {
	*vcl.TForm
	Button1  *vcl.TButton
	Image    *vcl.TJPEGImage
	OriImage *vcl.TPngImage
}

type TPoint struct {
	X, Y int32
	Down bool
}

var (
	isMouseDown bool
	points      = make([]TPoint, 0)
)

func (f *TranslaterForm) GetRealArea() string {
	return fmt.Sprintf("捕捉区域顶点[%d,%d],长%d,宽%d", f.GTRequest.X, f.GTRequest.Y, f.GTRequest.Width, f.GTRequest.Height)
}

func (f *CaptureForm) OnFormCreate(sender vcl.IObject) {
	gtcm := new(capture.GTCaptureModel)
	MyTranslaterForm.Con.GTCapture.CaptureFullScreen(gtcm)
	f.OriImage = vcl.NewPngImage()
	f.OriImage.LoadFromBytes(gtcm.Image)
	newImgBytes, err := ScaleImage(CaptureFormWidth, gtcm.Image)
	if err != nil {
		log.Panic("屏幕截图失败", err)
	}
	newImg := vcl.NewJPEGImage()
	newImg.LoadFromBytes(newImgBytes)
	logrus.Debug("屏幕截图加载完毕")
	f.Image = newImg

	var downX int32
	var downY int32
	var upX int32
	var upY int32

	zoom := float64(f.Image.Width()) / float64(f.OriImage.Width())
	logrus.Debug("系数", zoom)
	f.SetCaption("拖拽鼠标截图，截图区域即为选定翻译区域")
	f.SetWidth(CaptureFormWidth)
	f.SetHeight(int32(ZoomOut(int(f.OriImage.Height()), zoom)))
	f.ScreenCenter()
	f.SetDoubleBuffered(true)
	f.SetOnPaint(func(vcl.IObject) {

		canvas := f.Canvas()
		canvas.Draw(0, 0, f.Image)

		canvas.Pen().SetColor(colors.ClBlack)
		canvas.Pen().SetWidth(3)
		for _, p := range points {
			if p.Down {
				canvas.MoveTo(p.X, p.Y)
			} else {
				canvas.Draw(0, 0, f.Image)

				canvas.MoveTo(downX, downY)
				canvas.LineTo(downX, p.Y)
				canvas.LineTo(p.X, p.Y)
				canvas.MoveTo(downX, downY)
				canvas.LineTo(p.X, downY)
				canvas.LineTo(p.X, p.Y)

			}
		}
	})

	f.SetOnMouseDown(func(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
		if button == types.MbLeft {
			downX = x
			downY = y
			points = append(points, TPoint{X: x, Y: y, Down: true})
			isMouseDown = true
		}
	})

	f.SetOnMouseMove(func(sender vcl.IObject, shift types.TShiftState, x, y int32) {
		if isMouseDown {
			points = append(points, TPoint{X: x, Y: y, Down: false})
			f.Repaint()
		}
	})

	f.SetOnMouseUp(func(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
		if button == types.MbLeft {
			upX = x
			upY = y
			isMouseDown = false

			logrus.Debug("鼠标抬起，两个点坐标分别为", downX, downY, upX, upY)

			if upX > downX {
				MyTranslaterForm.GTRequest.X = ZoomIn(int(downX), zoom)
				MyTranslaterForm.GTRequest.Y = ZoomIn(int(downY), zoom)
			} else {
				MyTranslaterForm.GTRequest.X = ZoomIn(int(upX), zoom)
				MyTranslaterForm.GTRequest.Y = ZoomIn(int(upY), zoom)
			}

			MyTranslaterForm.GTRequest.Width = ZoomInFloat64(math.Abs(float64(upX)-float64(downX)), zoom)
			MyTranslaterForm.GTRequest.Height = ZoomInFloat64(math.Abs(float64(upY)-float64(downY)), zoom)
			MyTranslaterForm.Label.SetCaption(MyTranslaterForm.GetRealArea())
			message := fmt.Sprintf("截取了新的区域，起点坐标为,x:%d,y为%d,width为%d,height为%d", MyTranslaterForm.GTRequest.X, MyTranslaterForm.GTRequest.Y, MyTranslaterForm.GTRequest.Width, MyTranslaterForm.GTRequest.Height)
			IsInit = true
			vcl.ShowMessage(message)

			f.Hide()
		}
	})

}

func ZoomIn(x int, zoom float64) int {
	return int(float64(x) / zoom)
}

func ZoomOut(x int, zoom float64) int {
	return int(float64(x) * zoom)
}

func ZoomInFloat64(x float64, zoom float64) int {
	return int(x / zoom)
}

func ZoomOutFloat64(x float64, zoom float64) int {
	return int(x * zoom)
}

func ScaleImage(width int, imageBytes []byte) ([]byte, error) {
	img, _, err := image.Decode(bytes.NewBuffer(imageBytes))
	if err != nil {
		return nil, err
	}
	bound := img.Bounds()
	dx := bound.Dx()
	dy := bound.Dy()
	if err != nil {
		return nil, err
	}

	newImage := resize.Resize(uint(width), uint(width*dy/dx), img, resize.Lanczos3)
	logrus.Debug("新图片大小", newImage.Bounds().Dx(), newImage.Bounds().Dy())
	buffer := new(bytes.Buffer)
	err = jpeg.Encode(buffer, newImage, &jpeg.Options{Quality: 100})
	return buffer.Bytes(), err
}

func (f *CaptureForm) OnFormDestroy(sender vcl.IObject) {

}

func (f *CaptureForm) OnFormCloseQuery(sender vcl.IObject, canClose *bool) {
	*canClose = vcl.MessageDlg("是否退出?", types.MtConfirmation, types.MbYes, types.MbNo) == types.MrYes
}
