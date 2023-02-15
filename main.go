package main

import (
	"fmt"
	"time"
	"webshell/common"
	"webshell/shell"
)

func main() {
	start := time.Now()
	common.Parse1()
	common.Flag()
	shell.Plugins()
	common.File()
	end := time.Now().Sub(start)
	fmt.Println("[*]:", common.Filename, "\n[*]: 生成耗时:", end)
}
