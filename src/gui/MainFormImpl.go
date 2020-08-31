// 由res2go自动生成。
// 在这里写你的事件。

package gui

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ying32/govcl/vcl"
	"tailer/src/conf"
)

//::private::
type TMainFormFields struct {
	configuration *conf.Configuration
	protocols     []*conf.Protocol
}

//OnFormCreate
func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	f.ScreenCenter()
	f.ConnectItem.SetOnClick(f.onConnectItemClick)
	f.TestItem.SetOnClick(f.onTestItemClick)
	f.BtnFormat.SetOnClick(f.onBtnFormatClick)
	f.Btn2.SetOnClick(f.onBtn2Click)

	f.configuration.Load()
	var err error
	//读取所有协议
	f.protocols, err = conf.NewProtocols()
	if err != nil {
		fmt.Println("Protocols Read Error:", err)
		vcl.ShowMessage(err.Error())
	}
	//加载协议到ComboBox
	for _, v := range f.protocols {
		fmt.Println("Protocols Load:", v)
		f.CBProtocols.AddItem(v.Name, f.CBProtocols)
	}
}

//onConnectItemClick
func (f *TMainForm) onConnectItemClick(sender vcl.IObject) {
	ConnectForm.ShowModal()
}

//onTestItemClick
func (f *TMainForm) onTestItemClick(sender vcl.IObject) {
	f.configuration.Load()
	f.configuration.Save()
}

//onBtnFormatClick 内容校验与格式化
func (f *TMainForm) onBtnFormatClick(sender vcl.IObject) {
	data := f.Memo1.Text()
	var out bytes.Buffer
	err := json.Indent(&out, []byte(data), "", "  ")
	if err != nil {
		fmt.Println("Json Convert Error:", err)
		vcl.ShowMessage(err.Error())
	} else {
		f.Memo1.SetText(out.String())
	}
}

func (f *TMainForm) onBtn2Click(sender vcl.IObject) {

}
