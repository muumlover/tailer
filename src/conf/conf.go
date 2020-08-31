package conf

import (
	"encoding/json"
	"fmt"
	"os"
)

var windows = "%USERPROFILE%"
var linux = "%HOME%"

type TConfiguration struct {
	Enabled bool
	Path    string
}

var Configuration = new(TConfiguration)

func (conf *TConfiguration) configPath(name string) string {
	home, _ := os.UserConfigDir()
	return home + "/Tester/" + name
}

func (conf *TConfiguration) open() (*os.File, error) {
	var confPath = conf.configPath("test.json")
	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		fmt.Println("Config file does not exist, create it.")
		_ = os.Mkdir(conf.configPath(""), os.ModeDir)
		file, err := os.Create(conf.configPath("test.json"))
		if err != nil {
			fmt.Println("Create File Error:", err)
		}
		_, _ = file.WriteString("{}")
		_ = file.Close()
	}
	return os.Open(confPath)
}

func (conf *TConfiguration) Load() {
	var confPath = conf.configPath("test.json")
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
	var confPath = conf.configPath("test.json")
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
