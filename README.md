# GoodJob
### 项目管理平台
开发一个任务看板，统计工时，汇总周报






















---

## 开始项目

### GOROOT
golang安装路径,相当于java语言的`JAVA_HOME`

### GOPATH
工作目录，允许设置多个路径，但下载的包只存在第一个路径。和各个系统环境多路径设置一样，windows用`;`，linux（mac）用`:`分隔。
我们可以把每个GOPATH下的bin都加入到PATH中。


##### Linux系统
```
export GOROOT=$HOME/go
export GOPATH=$HOME/gopath
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

##### window系统
```bash
GOROOT=D:\go
GOPATH=D:\GoWorkSpace
PATH=%GOROOT%\bin;%GOPATH%\bin;...  # “...”表示其他路径
```

整个工作空间目录结构如下：
```
GoWorkSpace     // GoWorkSpace为GOPATH目录
  -- bin  // golang编译可执行文件存放路径，可自动生成。
  -- pkg  // golang编译的.a中间文件存放路径，可自动生成。
  -- src  // 源码路径。按照golang默认约定，go run，go install等命令的当前工作路径（即在此路径下执行上述命令）。
     -- common 1
     -- common 2
     -- common utils ...
     -- myApp1     // project1
        -- models
        -- controllers
        -- others
        -- main.go
     -- myApp2     // project2
        -- models
        -- controllers
        -- others
        -- main.go
     -- myApp3     // project3
        -- models
        -- controllers
        -- others
        -- main.go
```

> 注意不要把src目录和java项目的src目录混淆，误认为GoWorkSpace就是具体项目的根目录。
>idea开发工具open项目选择myApp目录，非GoWorkSpace目录


#### 项目运行
golang工程的依赖包经常使用go get命令来获取，例如：`go get github.com/kardianos/govendor` ，会将依赖包下载到GOPATH的路径下。
```
go get -u -v github.com/kardianos/govendor
```



常见的命令如下，格式为 `govendor COMMAND`。

| 命令 | 功能 | 
| :------------ :|:---------------| 
|init	|初始化 vendor 目录|
|list	|列出所有的依赖包|
|add	|添加包到 vendor 目录，如 `govendor add +external` 添加所有外部包`add PKG_PATH`添加指定的依赖包到 vendor 目录|
|update	|从 $GOPATH 更新依赖包到 vendor 目录|
|remove	|从 vendor 管理中删除依赖|
|status	|列出所有缺失、过期和修改过的包|
|fetch	|添加或更新包到本地 vendor 目录|
|sync	|本地存在 vendor.json 时候拉去依赖包，匹配所记录的版本|
|get 	|类似 go get 目录，拉取依赖包到 vendor 目录|

对于 govendor 来说，依赖包主要有以下多种类型:

|状态|缩写状态|含义|
|------------|------------|------------|
|+local|l|本地包，即项目自身的包组织|
|+external|e|外部包，即被 $GOPATH 管理，但不在 vendor 目录下|
|+vendor|v|已被 govendor 管理，即在 vendor 目录下|
|+std|s|标准库中的包|
|+unused|u|未使用的包，即包在 vendor 目录下，但项目并没有用到|
|+missing|m|代码引用了依赖包，但该包并没有找到|
|+program|p|主程序包，意味着可以编译为执行文件|
|+outside| |外部包和缺失的包|
|+all| |所有的包|

常用指令说明

```shell
# 安装govendor
go get -u -v github.com/kardianos/govendor

#将GOPATH中本工程使用到的依赖包自动移动到vendor目录中
#说明：如果本地GOPATH没有依赖包，先go get相应的依赖包
govendor add +external
或使用缩写： 
govendor add +e 

# 查看使用的包列表
govendor list -v fmt

# 从线上远端库添加或更新最新的依赖包
govendor fetch golang.org/x/net/context

# 从线上远端库添加或更新标签或分支等于v1的依赖包
govendor fetch golang.org/x/net/context@=v1
```


```shell
docker run --name mysql-8.0 \
-e MYSQL_ROOT_PASSWORD=root \
-p 3308:3306 \
-v /apps/data/mysql/conf-8.0:/etc/mysql/conf.d \
-v /apps/data/mysql/data-8.0:/var/lib/mysql -d mysql:8.0 \
--character-set-server=utf8mb4 \
--collation-server=utf8mb4_unicode_ci


alter user'root'@'%' IDENTIFIED BY 'root'; 
flush privileges;

ALTER USER 'root'@'%' IDENTIFIED BY 'root' PASSWORD EXPIRE NEVER; 
ALTER USER 'root'@'%' IDENTIFIED WITH mysql_native_password BY 'root';
FLUSH PRIVILEGES;
```


### 设置本地环境

```
set GO111MODULE=on
go mod tidy
```

### 打包可执行程序

（一）Windows 下编译Linux 64位可执行程序：
```
    SET CGO_ENABLED=0  //不设置也可以，原因不明
    SET GOOS=linux
    SET GOARCH=amd64
    通过 go env 查看设置是否成功。
```

（二）Linux 下编译Windows可执行程序：
```
    export CGO_ENABLED=0
    export GOOS=windows
    export GOARCH=amd64
    通过 go env 查看设置是否成功。
    go build hello.go
```
