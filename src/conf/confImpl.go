package conf

import (
	"encoding/json"
	"fmt"
	"os"
)

func (conf *TConfiguration) Load() {
	var confPath = configPath(configurationPath)
	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		conf.Enabled = false
		conf.Path = "Path"
		return
	}
	file, err := os.Open(confPath)
	if err != nil {
		fmt.Println("Open Config Error:", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Config Load Error:", err)
	}
	fmt.Println(conf.Path)
}

func (conf *TConfiguration) Save() {
	var confPath = configPath(configurationPath)
	file, err := os.Create(confPath)
	if err != nil {
		fmt.Println("Create Config Error:", err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(conf)
	if err != nil {
		fmt.Println("Config Save Error:", err)
	}
}
