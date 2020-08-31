package conf

import (
	"os"
	"strings"
)

//configPath
func configPath(name ...string) string {
	home, _ := os.UserConfigDir()
	s := []string{home, "Tester"}
	s = append(s, name...)
	return strings.Join(s, string(os.PathSeparator))
}
