// 由res2go自动生成。
// 在这里写你的事件。

package gui

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"github.com/ying32/govcl/vcl"
	"tailer/src/conf"
	"tailer/src/logger"
	"time"
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

	f.configuration.Load()
	var err error
	//读取所有协议
	f.protocols, err = conf.NewProtocols()
	if err != nil {
		logger.Error("Load Protocols Error:", err)
		vcl.ShowMessage(err.Error())
	}
	//加载协议到ComboBox
	for _, v := range f.protocols {
		f.CBProtocols.AddItem(v.Name, f.CBProtocols)
		logger.Trace("Protocol Added:", v.Name)
	}
	f.CBProtocols.SetOnExit(f.onCBProtocolsExit)

	//debug
	f.Memo1.SetText("{\n  \"gems_log_id\": 123,\n  \"size\": 12,\n  \"flag\": 0,\n  \"cmd\": 0,\n  \"param\": 0,\n  \"log_id\": 0,\n  \"attach\": 0,\n  \"checksum\": 0,\n  \"gems_data\": 0\n}")
}

//OnConnectItemClick
func (f *TMainForm) OnConnectItemClick(sender vcl.IObject) {
	ConnectForm.ShowModal()
}

//OnTestItemClick
func (f *TMainForm) OnTestItemClick(sender vcl.IObject) {
	f.configuration.Load()
	f.configuration.Save()
}

//OnBtnFormatClick 内容校验与格式化
func (f *TMainForm) OnBtnFormatClick(sender vcl.IObject) {
	data := f.Memo1.Text()
	var out bytes.Buffer
	err := json.Indent(&out, []byte(data), "", "  ")
	if err != nil {
		logger.Error("Json Convert Error:", err)
		vcl.ShowMessage(err.Error())
	} else {
		f.Memo1.SetText(out.String())
	}
}

func (f *TMainForm) onCBProtocolsExit(sender vcl.IObject) {
	selected := f.CBProtocols.ItemIndex()
	logger.Trace("Protocols selected index:", selected)
	if selected >= 0 {
		f.protocol = f.protocols[selected]
	} else {
		f.protocol = nil
	}
}

func (f *TMainForm) OnBtn2Click(sender vcl.IObject) {
	if f.protocol == nil {
		vcl.ShowMessage("Please select protocol.")
		return
	}
	data := f.Memo1.Text()
	v := make(map[string]interface{})
	err := json.Unmarshal([]byte(data), &v)
	if err != nil {
		logger.Error("Json Unmarshal Error:", err)
		vcl.ShowMessage(err.Error())
		return
	}
	dataBytes, err := f.protocol.ToByte(v)
	if err != nil {
		logger.Error("Protocol ToByte Error:", err)
		vcl.ShowMessage(err.Error())
		return
	}
	logger.Trace("Bytes:", hex.EncodeToString(dataBytes))
	//items := f.ListBox1.Items()
	//items.Insert(0, hex.EncodeToString(dataBytes))
	//f.ListBox1.SetItems(items)
	f.StringGrid1.InsertColRow(false, 0)
	f.StringGrid1.SetCells(0, 0, f.protocol.Name)
	f.StringGrid1.SetCells(1, 0, time.Now().Format("01/02 15:04:05.000"))
	f.StringGrid1.SetCells(2, 0, hex.EncodeToString(dataBytes))
	f.StringGrid1.ClearSelections()
	f.StringGrid1.SetRow(f.StringGrid1.RowCount() - 1)
	f.StringGrid1.SetRow(f.StringGrid1.RowCount())
	f.StringGrid1.Update()
	f.StringGrid1.InsertRowWithValues(0, []string{"a", "b", "c"})
	//f.StringGrid1.SetRowCount(f.StringGrid1.RowCount() + 1)
	//f.StringGrid1.SetCells(0, f.StringGrid1.RowCount()-1, f.protocol.Name)
	//f.StringGrid1.SetCells(1, f.StringGrid1.RowCount()-1, hex.EncodeToString(dataBytes))
}
