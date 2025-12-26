#!/bin/bash

# llm_client 推送脚本
# 使用方法: ./push.sh

echo "正在推送 llm_client 到 GitHub..."

# 推送代码
git push -u origin main

# 推送标签
git push origin v1.0.0

echo "✅ 推送完成!"
echo "仓库地址: https://github.com/zhimma/llm_client"
