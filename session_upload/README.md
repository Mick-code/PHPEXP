# SessionUploadExp

```
  -bf string
       用于阻塞进程的大文件
  -f string
       要注入到session中的文件
  -s string
       自定义phpsessid
  -u string
       目标url
  -zip
       是否利用zip协议绕过（会去掉注入文件的前16个字节）
```



```
  -bf string
        big filepath(for Blocking process),default value is  null
  -f string
        inject filepath,default value is  null
  -s string
        phpsessid,default value is bubb1esession (default "bubb1esession")
  -u string
        target url ,default value is  null
  -zip
        to use zip protocal ,default value is  false
```

关于zip协议绕过利用可参照文章