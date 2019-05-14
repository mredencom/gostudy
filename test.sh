#!/usr/bin/env bash
#!/bin/bash
fl=/etc/passwd
count=`cat $fl | wc -l`
#下面是一个管道，下面循环读文件中的每一行
cat $fl |
while read line
do
    user=`echo $line|awk -F ':' '{print $1}'`
    #代表以 ：分段$1就是取第1段
    uid=`echo $line|awk -F ':' '{print $3}'`
    echo "hello, $user Your UID is $uid"
done
echo "====User_count:$count===="
#前面求得的用户数