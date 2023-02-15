package shell

import (
	"webshell/common"
)

func Plugins() {
	if common.Webshell != " " {
		switch common.Webshell {
		case "php":
			Php()
		case "asp":
			Asp()
		case "jsp":
			Jsp()
		case "aspx":
			Aspx()
		case "jspx":
			Jspx()
		}

	}

}
