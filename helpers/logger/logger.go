package logger

import (
	"fmt"
	gcolor "github.com/daviddengcn/go-colortext"
	"log"
	"os"
)

func Info(message string, other ...string) {
	gcolor.ChangeColor(gcolor.Black, false, gcolor.Yellow, false)
	fmt.Print("[INFO]")
	gcolor.ResetColor()
	fmt.Print(":   ", message+"\n")
	if len(other) != 0 {
		for _, elem := range other {
			fmt.Println("	  ", elem)
		}
	}
}

func Error(message string, err interface{}) {
	gcolor.ChangeColor(gcolor.Black, false, gcolor.Red, true)
	log.Println("[ERROR]: ", message+"{", err, "}")
	gcolor.ResetColor()
}

func FatalError(message string, err error) {
	gcolor.ChangeColor(gcolor.Black, false, gcolor.Red, false)
	fmt.Print("[FATAL_ERROR]: ", message+"{", err, "}\n")
	gcolor.ResetColor()
	os.Exit(1)
}

func Debug(message interface{}) {
	gcolor.ChangeColor(gcolor.Black, false, gcolor.Green, false)
	fmt.Print("[DEBUG]")
	gcolor.ResetColor()
	fmt.Print(":   ", message, "\n")
}

func String(key string, message string) string {
	return "{" + key + "} : " + message
}
