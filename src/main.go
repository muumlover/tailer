package main

import (
	"github.com/ying32/govcl/vcl"
	"tailer/src/gui"
)

func main() {
	vcl.Application.Initialize()
	vcl.Application.CreateForm(&gui.MainForm)
	vcl.Application.CreateForm(&gui.ConnectForm)
	vcl.Application.Run()
}
