package conf

type TConfiguration struct {
	Enabled bool
	Path    string
}

var configurationPath = "test.json"
var Configuration = new(TConfiguration)
