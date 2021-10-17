package govclview

import (
	"fmt"
	"math"

	"github.com/dj456119/game-translater/capture"
	"github.com/dj456119/game-translater/controller"
	"github.com/sirupsen/logrus"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
	"github.com/ying32/govcl/vcl/types/keys"
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
	Con             *controller.GTController
	GTRequest       *controller.GTRequest
}

var MyTranslaterForm *TranslaterForm

func Init() {

	vcl.RunApp(&MyTranslaterForm, &MyCaptureForm)

}

func (f *TranslaterForm) OnFormCreate(sender vcl.IObject) {
	con := controller.NewGTController()
	f.Con = con

	request := new(controller.GTRequest)
	request.X = 0
	request.Y = 0
	request.Height = 900
	request.Width = 1600
	f.GTRequest = request

	f.SetShowInTaskBar(types.StAlways)
	f.SetCaption("游戏翻译机")
	f.EnabledMaximize(false)
	f.SetWidth(400)
	f.SetHeight(600)
	f.ScreenCenter()

	f.Label = vcl.NewLabel(f)
	f.Label.SetParent(f)
	f.Label.SetCaption(f.GetRealArea())
	f.Label.SetLeft(20)
	f.Label.SetTop(20)
	f.Label.SetBounds(20, 20, 360, 20)

	//设置Label
	f.Label1 = vcl.NewLabel(f)
	f.Label1.SetParent(f)
	f.Label1.SetCaption("实时捕捉的屏幕文字")
	f.Label1.SetLeft(20)
	f.Label1.SetTop(50)

	//设置实时捕捉的文本
	f.EditWords = vcl.NewMemo(f)
	f.EditWords.SetParent(f)
	f.EditWords.SetText("在这里显示正从屏幕获取的文本")
	f.EditWords.SetBounds(20, 70, 360, 200)

	//设置Labe2
	f.Label2 = vcl.NewLabel(f)
	f.Label2.SetParent(f)
	f.Label2.SetCaption("实时翻译的文字")
	f.Label2.SetLeft(20)
	f.Label2.SetTop(280)

	//设置实时翻译的文本
	f.EditTranslated = vcl.NewMemo(f)
	f.EditTranslated.SetParent(f)
	f.EditTranslated.SetText("在这里显示被翻译的文本")
	f.EditTranslated.SetBounds(20, 300, 360, 200)

	//设置按钮
	f.ButtonTranslate = vcl.NewButton(f)
	f.ButtonTranslate.SetParent(f)
	f.ButtonTranslate.SetCaption("启动翻译")
	f.ButtonTranslate.SetLeft(220)
	f.ButtonTranslate.SetTop(540)
	f.ButtonTranslate.SetOnClick(f.OnButtonTranslateClick)

	f.ButtonCapture = vcl.NewButton(f)
	f.ButtonCapture.SetParent(f)
	f.ButtonCapture.SetCaption("文本区域")
	f.ButtonCapture.SetLeft(20)
	f.ButtonCapture.SetTop(540)
	f.ButtonCapture.SetOnClick(f.OnButtonCaptureClick)
}

func (f *TranslaterForm) OnFormClose(sender vcl.IObject, action *types.TCloseAction) {
	vcl.Application.Terminate()
}

func (f *TranslaterForm) OnButtonTranslateClick(sender vcl.IObject) {
	words := f.EditWords.Text()
	//if words == "" {
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
	gtcm := new(capture.GTCaptureModel)
	MyTranslaterForm.Con.GTCapture.CaptureFullScreen(gtcm)
	MyCaptureForm.Image.LoadFromBytes(gtcm.Image)
	MyCaptureForm.Canvas().Draw(0, 0, MyCaptureForm.Image)
	MyCaptureForm.Show()
}

func (f *TranslaterForm) OnFormKeyPress(sender vcl.IObject, key *types.Char) {
	fmt.Println("key:", *key)
	if *key == keys.VkReturn {
		f.ButtonTranslate.Click()
	}
}

type CaptureForm struct {
	*vcl.TForm
	Button1 *vcl.TButton
	Image   *vcl.TPngImage
}

var MyCaptureForm *CaptureForm

type TPoint struct {
	X, Y int32
	Down bool
}

var (
	isMouseDown bool
	points      = make([]TPoint, 0)
)

func (f *TranslaterForm) GetRealArea() string {
	return fmt.Sprintf("当前捕捉的区域顶点坐标[%d,%d],长：%d，宽%d", f.GTRequest.X, f.GTRequest.Y, f.GTRequest.Width, f.GTRequest.Height)
}

func (f *CaptureForm) OnFormCreate(sender vcl.IObject) {
	gtcm := new(capture.GTCaptureModel)
	MyTranslaterForm.Con.GTCapture.CaptureFullScreen(gtcm)

	f.Image = vcl.NewPngImage()
	f.Image.LoadFromBytes(gtcm.Image)

	var downX int32
	var downY int32
	var upX int32
	var upY int32

	f.SetCaption("拖拽鼠标截图，截图区域即为选定翻译区域")
	f.SetWidth(1600)
	f.SetHeight(900)
	f.ScreenCenter()
	f.SetDoubleBuffered(true)
	f.SetOnPaint(func(vcl.IObject) {

		canvas := f.Canvas()

		canvas.Draw(0, 0, f.Image)

		canvas.Pen().SetColor(colors.ClBlack)
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
			MyTranslaterForm.GTRequest.X = int(downX)
			MyTranslaterForm.GTRequest.Y = int(downY)
			MyTranslaterForm.GTRequest.Width = int(math.Abs(float64(upX) - float64(downX)))
			MyTranslaterForm.GTRequest.Height = int(math.Abs(float64(upY) - float64(downY)))
			MyTranslaterForm.Label.SetCaption(MyTranslaterForm.GetRealArea())
			message := fmt.Sprintf("截取了新的区域，起点坐标为,x:%d,y为%d,width为%d,height为%d", downX, downY, MyTranslaterForm.GTRequest.Width, MyTranslaterForm.GTRequest.Height)
			vcl.ShowMessage(message)

			f.Hide()
		}
	})

}

func (f *CaptureForm) OnFormDestroy(sender vcl.IObject) {

}

func (f *CaptureForm) OnFormCloseQuery(sender vcl.IObject, canClose *bool) {
	*canClose = vcl.MessageDlg("是否退出?", types.MtConfirmation, types.MbYes, types.MbNo) == types.MrYes
}
