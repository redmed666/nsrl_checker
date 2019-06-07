package util

import (
	"os"
)

func Check(err error) {
	if err != nil {
		println("[x] Error: " + err.Error())
		os.Exit(-1)
	}
}
