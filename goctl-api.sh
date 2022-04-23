#!/bin/bash

# 获取脚本所在目录
CURRENT_DIR=$(cd $(dirname $0); pwd)
echo $CURRENT_DIR
cd $CURRENT_DIR/api/doc && goctl api go -api server.api -dir $CURRENT_DIR/api/ -home $CURRENT_DIR/api/doc/template