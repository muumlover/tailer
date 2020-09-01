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
	protocol      *conf.Protocol
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
	f.CBProtocols.SetOnExit(f.onCBProtocolsExit)

	//debug
	f.Memo1.SetText("{\n  \"gems_log_id\": 123,\n  \"size\": 12,\n  \"flag\": 0,\n  \"cmd\": 0,\n  \"param\": 0,\n  \"log_id\": 0,\n  \"attach\": 0,\n  \"checksum\": 0,\n  \"gems_data\": 0\n}")
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

func (f *TMainForm) onCBProtocolsExit(sender vcl.IObject) {
	selected := f.CBProtocols.ItemIndex()
	fmt.Println("Protocols selected index:", selected)
	if selected >= 0 {
		f.protocol = f.protocols[selected]
	} else {
		f.protocol = nil
	}
}

func (f *TMainForm) onBtn2Click(sender vcl.IObject) {
	if f.protocol == nil {
		vcl.ShowMessage("Please select protocol.")
		return
	}
	data := f.Memo1.Text()
	v := make(map[string]interface{})
	err := json.Unmarshal([]byte(data), &v)
	if err != nil {
		fmt.Println("Json Unmarshal Error:", err)
		vcl.ShowMessage(err.Error())
		return
	}
	err = f.protocol.ToByte(v)
	if err != nil {
		fmt.Println("Protocol ToByte Error:", err)
		vcl.ShowMessage(err.Error())
		return
	}
}
