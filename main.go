package main

import (
	"emptyImageProject/http"
        "emptyImageProject/polling"
        "fmt"
	"flag"
        "time"
)

const ZONE_DIFF=3600*8-12

var ipypath string
var pipeline string
var output string

func init(){
	flag.StringVar(&ipypath,"ipypath","/home/jovyan/wzy-Folder/nfs-demo.py","python scripts absolute path")
	flag.StringVar(&pipeline,"pipeline","testpipelines","the pipeline name (lower case with - instead of _)")
	flag.StringVar(&output,"output","/mnt/xfs/pipeline_server/output/workflowArr.log","monitor output path")
}

func main(){
	timeUnixNano := time.Now().UnixNano()
	timeUnixMirco := float64(timeUnixNano)/1000000000
	//fmt.Printf("纳秒时间为：%d\n",timeUnixNano)
	fmt.Printf("容器启动后时间戳为：%.3f\n",timeUnixMirco)
	var host string = "http://10.18.127.2:8081/pipeline"
	http.HttpPost(host,ipypath,pipeline)
	//var output = "/mnt/xfs/pipeline_server/output/workflowArr.log"
	fmt.Printf("发送请求后时间戳(调整后)为:%d\n",time.Now().UTC().Unix()-ZONE_DIFF)
	fmt.Println(time.Now().UTC().Format(polling.TIME_LAYOUT))
	fmt.Println(time.Now().UTC().Unix())
	polling.Track(time.Now().UTC().Unix()-ZONE_DIFF, pipeline, output)
}



