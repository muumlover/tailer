// 由res2go自动生成。
// 在这里写你的事件。

package gui

import "github.com/ying32/govcl/vcl"

//::private::
type TMainFormFields struct {
}

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	f.ConnectItem.SetOnClick(f.OnConnectItemClick)
}

func (f *TMainForm) OnConnectItemClick(sender vcl.IObject) {
	ConnectForm.ShowModal()
}
