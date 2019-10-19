current=`date "+%Y-%m-%d %H:%M:%S"`
timeStamp=`date -d "$current" +%s`
currentTimeStamp=$(($timeStamp*1000+`date "+%N"`/1000000))
beforeStartTime=`echo "scale=3;$currentTimeStamp/1000"|bc`
echo "容器启动前时间戳："$beforeStartTime
docker run -v /mnt/xfs:/mnt/xfs empty
