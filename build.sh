#!/bin/bash
MYARCH=`go env GOARCH`
if [ "$1" != "" ];then
	MYARCH=$1
fi
export GOARCH=$MYARCH
go build -v -a -o hostnic-plugin.$GOARCH hostnic-plugin
echo "program is hostnic-plugin.$GOARCH"

