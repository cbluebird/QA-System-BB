# QA-System-Server
定制化问卷系统的后端

复制yaml，填写配置

注意yaml中的allow_origins是前端的地址，port是运行端口

```cmd
go build -o QA
```

在redis中配置

```redis
set password 密码
set admin 用户名
```



然后运行上面打包文件出来的即可