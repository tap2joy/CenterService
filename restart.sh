#!/usr/bin/env bash

pid=$(ps -ef | grep CenterService | grep -v grep | awk '{print $2}')
if [[ -n "$pid" ]]; then
    echo "kill old process, pid: "$pid
    kill -9 $pid
fi
nohup ./CenterService &