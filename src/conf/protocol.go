package conf

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
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
		logger.Println("Read Protocols Error", err)
		return nil, err
	}
	var protocols = make([]*Protocol, 0)
	for i, f := range files {
		logger.Println("Protocol", i, "=", f.Name())
		p, err := NewProtocol(f.Name())
		if err != nil {
			logger.Println("Load Protocol Error", err)
			continue
		}
		protocols = append(protocols, p)
		logger.Println("Protocol Loaded:", p)
	}
	return protocols, nil
}

func NewProtocol(path string) (*Protocol, error) {
	protocol := &Protocol{}
	protocol.Name = path
	var confPath = configPath(ProtocolPath, path)
	file, err := os.Open(confPath)
	if err != nil {
		logger.Println("Open Protocol Error:", err)
		return nil, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	decoder.UseNumber()
	err = decoder.Decode(&protocol)
	if err != nil {
		logger.Println("Read Protocol Error:", err)
		return nil, err
	}
	logger.Println("Protocol", protocol.Name, "Loaded")
	return protocol, nil
}

func (p Protocol) ToByte(data map[string]interface{}) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	logger.Println()
	for index, field := range p.Head {
		if field["name"] == nil {
			return nil, errors.New("protocol head miss name at index[" + strconv.Itoa(index) + "]")
		}
		valueName := field["name"].(string)
		logger.Println("valueName:", valueName)

		if field["type"] == nil {
			return nil, errors.New("protocol head miss type at index[" + strconv.Itoa(index) + "]")
		}
		valueType := field["type"].(string)
		logger.Println("valueType:", valueType)

		valueSize := 1
		if field["size"] != nil {
			size, _ := field["size"].(json.Number).Int64()
			valueSize = int(size)
		}
		logger.Println("valueSize:", valueSize)

		valueDefault := field["default"]
		logger.Println("valueDefault:", valueDefault)

		var value interface{}
		if valueDefault != nil {
			value = valueDefault
		} else {
			value = data[valueName]
		}
		logger.Println("value:", value)
		if value == nil {
			return nil, errors.New("miss field '" + valueName + "'")
		}

		var valueI64 int64 = 0
		switch valueType {
		case "int32", "uint32", "int64", "uint64":
			switch value.(type) {
			case json.Number:
				logger.Println("value is json.Number")
				valueI64, _ = value.(json.Number).Int64()
			case float64:
				logger.Println("value is float64")
				valueI64 = int64(value.(float64))
			default:
				logger.Println("value is error")
			}
		case "byte":
			switch value.(type) {
			case json.Number:
				logger.Println("value is json.Number")
				valueI64, _ = value.(json.Number).Int64()
				logger.Println("json.Number:", valueI64)
			default:
				logger.Println("value is error")
			}
		}

		logger.Println("valueI64:", valueI64)

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

		logger.Println()
	}
	return buf.Bytes(), nil
}
