1.2：

```
1：Unicode编码，目前只支持哥斯拉jsp编码
2：当未指定密码时，将生成随机的密码，随机的文件名
```

![img](https://pic.certbug.com/i/2023/02/28/xqtkxd.webp)



1.1：

```

              _         _          _ _
__      _____| |__  ___| |__   ___| | |
\ \ /\ / / _ \ '_ \/ __| '_ \ / _ \ | |
 \ V  V /  __/ |_) \__ \ | | |  __/ | |
  \_/\_/ \___|_.__/|___/_| |_|\___|_|_|
                                   vsersion  1.1
                -p password | Default password  noway
                -s php | jsp | asp | aspx
                -e xor | aes
                -d be | god

  -b string
        something bypass waf | php
  -d string
        -d Behinder(Be) | Godzilla(God)
  -e string
         xor aes 128 only Behinder and Godzilla
  -h    help
  -l string
        -l spring tomacat resin jdk
  -p string
        -p password (default "noway")
  -s string
        -s php | jsp | asp | aspx
       
```

There are some tips you can get :

```
Generate memory shell
Usage of God
tomcat:
.\webshell.exe -d God -e aes -l tomcat
spring:
 .\webshell.exe -d God -e aes -l spring
 resin:
 .\webshell.exe -d God -e aes -l resin
 jdk:
  .\webshell.exe -d God -e aes -l jdk
  
  
  Usage of Be (xor aes)
  tomcat:
   .\webshell.exe -d be -e xor -l tomcat
  spring:
   .\webshell.exe -d be -e xor -l spring
  resin:
   .\webshell.exe -d be -e xor -l resin
  
```

 what is more you need to konw is that  Godzilla only has aes code

About the bypass option, only generation is currently supported And only php>7.0 is supported

```
The following is general usage
 .\webshell.exe -s php -p 666
```

![img](https://pic.certbug.com/i/2023/02/28/xqtbhl.webp)

![img](https://pic.certbug.com/i/2023/02/28/xqtwzh.webp)

End:Update every two months





