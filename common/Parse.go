package common

import "os"

func Parse1() {
	Paseshell()
}

func Paseshell() {
	if Webshell == " " && Password == " " {
		os.Exit(0)
	}
}
