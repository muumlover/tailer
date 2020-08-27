package main


import (
	// 如果你使用自定义的syso文件则不要引用此包
	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
)

type TMainForm struct {
	*vcl.TForm
	Btn1     *vcl.TButton
}

type TAboutForm struct {
	*vcl.TForm
	Btn1    *vcl.TButton
}

var (
	mainForm *TMainForm
	aboutForm *TAboutForm
)

func main() {
	vcl.RunApp(&mainForm, &aboutForm)
}

// -- TMainForm

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	f.SetCaption("MainForm")

	f.Btn1 = vcl.NewButton(f)
	f.Btn1.SetParent(f)
	f.Btn1.SetBounds(10, 10, 88, 28)
	f.Btn1.SetCaption("Button1")
	f.Btn1.SetOnClick(f.OnBtn1Click)
}

func (f *TMainForm) OnBtn1Click(sender vcl.IObject) {
	aboutForm.Show()
}


// -- TAboutForm

func (f *TAboutForm) OnFormCreate(sender vcl.IObject) {
	f.SetCaption("About")
	f.Btn1 = vcl.NewButton(f)
	//f.Btn1.SetName("Btn1")
	f.Btn1.SetParent(f)
	f.Btn1.SetBounds(10, 10, 88, 28)
	f.Btn1.SetCaption("Button1")
	f.Btn1.SetOnClick(f.OnBtn1Click)
}

func (f *TAboutForm) OnBtn1Click(sender vcl.IObject) {
	vcl.ShowMessage("Hello!")
}