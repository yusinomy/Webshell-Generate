package main

import (
	"fmt"
	"time"
	"webshell/common"
	"webshell/shell"
)

func main() {
	start := time.Now()
	shell.Exec()
	end := time.Now().Sub(start)
	fmt.Println("[*]FileName:", common.Filename, "\n[*]Password:", common.Password, "\n[*]生成耗时:", end)
}
