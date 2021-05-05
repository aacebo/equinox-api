package log

import (
	"log"
	"os"

	"github.com/fatih/color"
)

var Info = log.New(os.Stderr, color.CyanString("[INFO] "), log.Ldate|log.Ltime|log.Lshortfile)
var Warn = log.New(os.Stderr, color.YellowString("[WARN] "), log.Ldate|log.Ltime|log.Lshortfile)
var Error = log.New(os.Stderr, color.RedString("[ERROR] "), log.Ldate|log.Ltime|log.Lshortfile)
