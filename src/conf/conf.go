package conf

import (
	"encoding/json"
	"os"
	"tailer/src/logger"
)

var configurationPath = "test.json"

type Configuration struct {
	Enabled  bool      `json:"enable"`
	Path     string    `json:"path"`
	Connects []Connect `json:"connects"`
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
		logger.Error("Open Config Error:", err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	decoder.UseNumber()
	err = decoder.Decode(&c)
	if err != nil {
		logger.Error("Load Config Error:", err)
		return
	}
	logger.Trace("Config Loaded:", c)
}

//Save 保存设置
func (c *Configuration) Save() {
	var confPath = configPath(configurationPath)
	file, err := os.Create(confPath)
	if err != nil {
		logger.Error("Create Config Error:", err)
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(c)
	if err != nil {
		logger.Error("Save Config Error:", err)
		return
	}
}
