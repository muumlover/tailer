package conf

import (
	"fmt"
	"io/ioutil"
	"os"
)

const ProtocolPath = "protocols"

type TProtocol struct {
	Name string
	Data string
}

var Protocol *TProtocol

func (p *TProtocol) NewProtocol(name string) (protocol *TProtocol) {
	protocol = new(TProtocol)
	protocol.Name = name
	return protocol
}

type TProtocols []TProtocol

var Protocols *TProtocols

func (p *TProtocols) NewProtocols(path string) (protocols *TProtocols) {
	Protocol.NewProtocol("s")
	return
}

func (p *TProtocols) Load() ([]*TProtocol, error) {
	var confPath = configPath(ProtocolPath)
	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		err = os.Mkdir(confPath, os.ModeDir)
		return nil, err
	}
	files, err := ioutil.ReadDir(confPath)
	if err != nil {
		fmt.Println("Read Protocols Error", err)
		return nil, err
	}
	var protocols = make([]*TProtocol, len(files))
	for i, v := range files {
		fmt.Println(i, "=", v.Name())
		t := Protocol.NewProtocol(v.Name())
		protocols[i] = t
	}
	return protocols, nil
}
