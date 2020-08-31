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
	f.Button1.SetOnClick(f.onButton1Click)
}

func (f *TMainForm) onConnectItemClick(sender vcl.IObject) {
	ConnectForm.ShowModal()
}

func (f *TMainForm) onTestItemClick(sender vcl.IObject) {
	conf.Configuration.Load()
	conf.Configuration.Save()
}
func (f *TMainForm) onButton1Click(sender vcl.IObject) {
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
