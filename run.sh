#!/usr/bin/env bash

ulimit -u 10240
# 开启keeplive
./main -k=true -c 16 -n 1 -duration 30 -data world -u 127.0.0.1:8088/public/logic/
