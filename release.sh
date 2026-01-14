#!/bin/bash
# Go 包发布脚本 - v1.0.0

set -e  # 遇到错误立即退出

echo "🚀 准备发布 llm_client v1.0.0"
echo ""

# 1. 检查是否有未提交的更改
echo "📋 检查 Git 状态..."
if [[ -n $(git status -s) ]]; then
    echo "⚠️  发现未提交的更改:"
    git status -s
    echo ""
    read -p "是否继续提交这些更改? (y/n) " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        echo "❌ 发布已取消"
        exit 1
    fi
else
    echo "✅ 工作区干净"
fi

# 2. 运行测试
echo ""
echo "🧪 运行测试..."
go test -v ./...
if [ $? -eq 0 ]; then
    echo "✅ 所有测试通过"
else
    echo "❌ 测试失败,发布已取消"
    exit 1
fi

# 3. 提交代码
echo ""
echo "📝 提交代码..."
git add .
git commit -m "feat!: change timeout unit from nanoseconds to seconds

BREAKING CHANGE: Timeout field type changed from time.Duration to int (seconds).
This affects ChatCompletionRequest, EmbeddingRequest, and Config structs.

- types.go: Timeout field now uses int (seconds) instead of time.Duration
- config.go: Default timeout changed from 600*time.Second to 600
- client.go: Added conversion logic from seconds to time.Duration
- Added CHANGELOG.md to track version changes
- Added unit tests for timeout serialization

Third-party users need to update their code to pass timeout as seconds."

echo "✅ 代码已提交"

# 4. 推送到远程
echo ""
echo "⬆️  推送到远程仓库..."
git push origin main
echo "✅ 代码已推送"

# 5. 创建版本标签
echo ""
echo "🏷️  创建版本标签 v0.1.2..."
git tag -a v0.1.2 -m "v0.1.2 - Breaking change: timeout unit changed to seconds

Major version bump due to breaking API changes:
- Timeout field type changed from time.Duration to int (seconds)
- Affects ChatCompletionRequest, EmbeddingRequest, and Config
- JSON serialization now outputs seconds instead of nanoseconds
- Default timeout: 600 seconds (10 minutes)

Migration guide available in README.md"

echo "✅ 标签已创建"

# 6. 推送标签
echo ""
echo "⬆️  推送标签到远程..."
git push origin v0.1.2
echo "✅ 标签已推送"

# 7. 完成
echo ""
echo "🎉 发布完成!"
echo ""
echo "📦 版本: v0.1.2"
echo "📍 仓库: github.com/zhimma/llm_client"
echo ""
echo "下一步:"
echo "1. 访问 https://github.com/zhimma/llm_client/releases"
echo "2. 创建 Release Notes (可参考 release-guide.md)"
echo "3. 通知使用方更新到新版本"
echo ""
echo "第三方更新命令:"
echo "  go get github.com/zhimma/llm_client@v0.1.2"
echo "  go mod tidy"
