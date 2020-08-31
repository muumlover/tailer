package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const ProtocolPath = "protocols"

type Protocol struct {
	Name string
	Data string
}

func NewProtocol(path string) (*Protocol, error) {
	protocol := &Protocol{}
	protocol.Name = path
	var confPath = configPath(ProtocolPath, path)
	file, err := os.Open(confPath)
	if err != nil {
		fmt.Println("Open Config Error:", err)
		return nil, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&protocol)
	if err != nil {
		fmt.Println("Config Load Error:", err)
		return nil, err
	}
	fmt.Println(protocol.Name)
	return protocol, nil
}

func NewProtocols() ([]*Protocol, error) {
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
	var protocols = make([]*Protocol, 0)
	for i, v := range files {
		fmt.Println(i, "=", v.Name())
		t, err := NewProtocol(v.Name())
		if err != nil {
			fmt.Println("Read Protocol Error", err)
		} else {
			protocols = append(protocols, t)
		}
	}
	return protocols, nil
}
