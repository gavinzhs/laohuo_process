#!/usr/bin/env bash
hosts=(123.56.135.128)
app=dizhi


dir=/opt/$app
echo "cross compile"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

for host in ${hosts[@]}
do
    echo "deal with $host"
    echo "prepare"
    ssh root@$host "mkdir -p $dir"

    echo "scp"
    scp restart.sh root@$host:$dir/
    scp $app root@$host:$dir/$app.new

    echo "remote operate"
    ssh root@$host "cd $dir && fuser -k $app"
    ssh root@$host "cd $dir && mv $app.new $app && sh restart.sh `</dev/null` >nohup.out 2>&1 &"
done


rm -rf $app
echo "done"