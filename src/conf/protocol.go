package conf

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

const ProtocolPath = "protocols"

type Protocol struct {
	Name   string                   `json:"name"`
	Encode string                   `json:"encode"`
	Const  map[string]interface{}   `json:"const"`
	Head   []map[string]interface{} `json:"head"`
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
	decoder.UseNumber()
	err = decoder.Decode(&protocol)
	if err != nil {
		fmt.Println("Config Load Error:", err)
		return nil, err
	}
	fmt.Println(protocol.Name)
	return protocol, nil
}

func (p Protocol) ToByte(data map[string]interface{}) error {
	buf := bytes.NewBuffer([]byte{})
	for index, field := range p.Head {
		if field["name"] == nil {
			return errors.New("protocol head miss name at index[" + strconv.Itoa(index) + "]")
		}
		valueName := field["name"].(string)
		fmt.Println("valueName: ", valueName)

		if field["type"] == nil {
			return errors.New("protocol head miss type at index[" + strconv.Itoa(index) + "]")
		}
		valueType := field["type"].(string)
		fmt.Println("valueType: ", valueType)

		valueSize := 1
		if field["size"] != nil {
			size, _ := field["size"].(json.Number).Int64()
			valueSize = int(size)
		}
		fmt.Println("valueSize: ", valueSize)

		valueDefault := field["default"]
		fmt.Println("valueDefault: ", valueDefault)

		var value interface{}
		if valueDefault != nil {
			value = valueDefault
		} else {
			value = data[valueName]
		}
		fmt.Println("value: ", value)
		if value == nil {
			return errors.New("miss field '" + valueName + "'")
		}

		var valueI64 int64 = 0
		switch valueType {
		case "int32":
		case "uint32":
		case "int64":
		case "uint64":
			switch value.(type) {
			case json.Number:
				valueI64, _ = value.(json.Number).Int64()
			case float64:
				valueI64 = int64(value.(float64))
			}
		case "byte":
			switch value.(type) {
			case json.Number:
				valueI64, _ = value.(json.Number).Int64()
			}
		}

		fmt.Println("valueI64: ", valueI64)

		switch valueType {
		case "int32":
			_ = binary.Write(buf, binary.LittleEndian, int32(valueI64))
		case "uint32":
			_ = binary.Write(buf, binary.LittleEndian, uint32(valueI64))
		case "int64":
			_ = binary.Write(buf, binary.LittleEndian, valueI64)
		case "uint64":
			_ = binary.Write(buf, binary.LittleEndian, uint64(valueI64))
		case "byte":
			_ = binary.Write(buf, binary.LittleEndian, byte(valueI64))
		}
	}
	return nil
}
