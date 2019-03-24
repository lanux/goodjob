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
```
go get -u -v github.com/kardianos/govendor
```

