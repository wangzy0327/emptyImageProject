# emptyImageProject
emptyImage


### 常用命令
- 静态编译go语言
```
编译为amd64指令集架构,项目名为emptyImageProject(在GOPATH路径src目录下),编译后的二进制文件为empty
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o empty -ldflags '-s' emptyImageProject/

编译为mips64le指令集架构(龙芯),项目名为emptyImageProject(在GOPATH路径src目录下),编译后的二进制文件为empty0
CGO_ENABLED=0 GOOS=linux GOARCH=mips64le go build -a -o empty0 -ldflags '-s' emptyImageProject/

编译为arm64指令集架构,项目名为emptyImageProject(在GOPATH路径src目录下),编译后的二进制文件为empty1
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -o empty1 -ldflags '-s' emptyImageProject/

```

- 查看go语言支持的平台
```
go tool dist list

```

- 基于scratch构建Dockerfile运行
```
docker run -v /mnt/xfs:/mnt/xfs empty /empty --help

Usage of /empty:
  -ipypath string
    	python scripts absolute path (default "/home/jovyan/wzy-Folder/nfs-demo.py")
  -output string
    	monitor output path (default "/mnt/xfs/pipeline_server/OUTPUT/workflowArr.log")
  -pipeline string
    	the pipeline name (lower case with - instead of _) (default "testpipelines")

```