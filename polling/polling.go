package polling

import (
	"fmt"
        "strings"
	"strconv"
	"time"
	"emptyImageProject/file"
)

const TIME_LAYOUT="2006-01-02 15:04:05"

type poll interface{

}

func count() int{
	return 10;
}

func diffMax() int64{
	return 5;
}

func diffMin() int64{
	return -5;
}

func Track(curTimeStamp int64,pipeline string,output string){
for ;;{  
         workflowArr,workflowStatus,startAtTime:=parseInfo(output)
	 var flag = 0
	 //fmt.Println("cur time : "+strconv.FormatInt(curTimeStamp,10))
	 for i:=0;i<len(workflowArr);i++  {
		workflowStartAtTime:=strings.Replace(strings.Replace(startAtTime[i],"T"," ",-1),"Z","",-1)
		loc,_:=time.LoadLocation("Local")
		theTime,_:=time.ParseInLocation(TIME_LAYOUT,workflowStartAtTime,loc)
		workflowTimeStamp:=theTime.Unix()
		if workflowTimeStamp == -62135596800 {
		   fmt.Println(" Time Error !!!")
		   return ;
		}
		//fmt.Println(theTime)
		//fmt.Println(workflowTimeStamp)
		//fmt.Println("workflowArr : "+workflowArr[i])
		//fmt.Println("workflowStatus : "+workflowStatus[i])
		//fmt.Println("startAtTime : "+theTime.Format(TIME_LAYOUT))
		workflowTimeStampDiff:=workflowTimeStamp - curTimeStamp
                //fmt.Println("timestamp : "+strconv.FormatInt(workflowTimeStamp,10))
                //fmt.Println("curstamp : "+strconv.FormatInt(curTimeStamp,10))
		//fmt.Println("diff : "+strconv.FormatInt(workflowTimeStampDiff,10))
                if strings.Index(workflowArr[i],pipeline)!=-1 && workflowTimeStampDiff > diffMin() && workflowTimeStampDiff < diffMax(){
			fmt.Println("workflow name : "+workflowArr[i])
			fmt.Println("workflow status : "+workflowStatus[i])
			fmt.Println("......................over....................")
			flag=1
			break
		}
	}
	if flag == 1{break}
}
}

func datetimeHandler(startAtTime []string){
   for i:=0;i<len(startAtTime)-1;i++{
     workflowStartAtTime:=strings.Replace(strings.Replace(startAtTime[i],"T"," ",-1),"Z","",-1)
     fmt.Println("handler time : "+workflowStartAtTime)
     loc,_:=time.LoadLocation("Local")
     //theTime,_:=time.ParseInLocation(TIME_LAYOUT,"2018-09-10 00:00:00",loc)
     theTime,_:=time.ParseInLocation(TIME_LAYOUT,workflowStartAtTime,loc)
     workflowTimeStamp:=theTime.Unix()
     fmt.Println(theTime.Format(TIME_LAYOUT))
     fmt.Println(workflowTimeStamp) 
   }
}

func parseInfo(output string)([]string,[]string,[]string){
  var strs = file.ReadLastBlock(output,count())
	var workflowArr []string
	var workflowStatus [] string
	var startAtTime []string
	for j:=0;j<len(strs);j++{
		strarray := strings.Fields(strings.TrimSpace(strs[j]))
		for i:=0;i<len(strarray);i++ {
			if(i == 0){
				workflowArr = append(workflowArr, strarray[i])
			}
			if(i == 1){
				workflowStatus = append(workflowStatus,strarray[i])
			}
			if(i == 2){
				startAtTime = append(startAtTime,strarray[i])
			}
		}
	}
	/*
	for i:=0;i<len(workflowArr);i++{
		fmt.Println(workflowArr[i])
	}
	for i:=0;i<len(workflowStatus);i++{
		fmt.Println(workflowStatus[i])
	}
	for i:=0;i<len(startAtTime);i++{
		fmt.Println(startAtTime[i])
	}
	*/
	return workflowArr,workflowStatus,startAtTime 
}

