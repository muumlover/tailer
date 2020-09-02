package util

import (
	"log"
	"os"
)

var Logger = log.New(os.Stderr, "", log.LstdFlags|log.Llongfile)
