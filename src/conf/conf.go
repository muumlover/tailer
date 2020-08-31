package conf

import (
	"encoding/json"
	"fmt"
	"os"
)

var configurationPath = "test.json"

type Configuration struct {
	Enabled  bool
	Path     string
	Connects []*Connect
}

//Load 读取设置
func (c *Configuration) Load() {
	var confPath = configPath(configurationPath)
	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		c.Enabled = false
		c.Path = "Path"
		return
	}
	file, err := os.Open(confPath)
	if err != nil {
		fmt.Println("Open Config Error:", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c)
	if err != nil {
		fmt.Println("Config Load Error:", err)
	}
	fmt.Println(c.Path)
}

//Save 保存设置
func (c *Configuration) Save() {
	var confPath = configPath(configurationPath)
	file, err := os.Create(confPath)
	if err != nil {
		fmt.Println("Create Config Error:", err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(c)
	if err != nil {
		fmt.Println("Config Save Error:", err)
	}
}
