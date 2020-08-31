package conf

import (
	"os"
)

func configPath(name string) string {
	pathSep := string(os.PathSeparator)
	home, _ := os.UserConfigDir()
	return home + pathSep + "Tester" + pathSep + name
}
