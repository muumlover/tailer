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
}

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	f.ScreenCenter()
	f.ConnectItem.SetOnClick(f.onConnectItemClick)
	f.TestItem.SetOnClick(f.onTestItemClick)
	f.ButtonFormat.SetOnClick(f.onButtonFormatClick)
	f.Button2.SetOnClick(f.onButton2Click)

	protocols, err := conf.Protocols.Load()
	if err != nil {
		fmt.Println("Protocols Load Error:", err)
		vcl.ShowMessage(err.Error())
	}
	for _, v := range protocols {
		fmt.Println("Protocols Load Error:", v)
		f.ComboBox1.AddItem(v.Name, f.ComboBox1)
	}
}

//onConnectItemClick
func (f *TMainForm) onConnectItemClick(sender vcl.IObject) {
	ConnectForm.ShowModal()
}

//onTestItemClick
func (f *TMainForm) onTestItemClick(sender vcl.IObject) {
	conf.Configuration.Load()
	conf.Configuration.Save()
}

//onButtonFormatClick 内容校验与格式化
func (f *TMainForm) onButtonFormatClick(sender vcl.IObject) {
	data := f.Memo1.Text()
	var out bytes.Buffer
	err := json.Indent(&out, []byte(data), "", "    ")
	if err != nil {
		fmt.Println("Json Convert Error:", err)
		vcl.ShowMessage(err.Error())
	} else {
		f.Memo1.SetText(out.String())
	}
}

func (f *TMainForm) onButton2Click(sender vcl.IObject) {

}
