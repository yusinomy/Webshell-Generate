package common

import (
	"flag"
	"fmt"
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
		                   vsersion 1.0
		-p 生成密码，默认密码为rural666		
		-s php | jsp | asp | aspx`
	fmt.Println(titles)
}

func Flag() {
	title()
	flag.StringVar(&Webshell, "s", "", "-s php | jsp | asp | aspx")
	flag.StringVar(&Password, "p", "rural666", "-p 生成密码，默认密码为rural666")
	flag.StringVar(&Help, "h", "", "help")
	//flag.StringVar(&encode, "e", "", "指定编码")
	//flag.StringVar(&ccccc, "d", "", "-d 冰蝎 | 哥斯拉 ")
	//flag.Parse()

}
