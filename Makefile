# Tea Framework Makefile

# 默认目标
all: help

# 帮助信息
help:
	@echo "Tea Framework 命令列表:"
	@echo "  make run-example   - 运行示例项目"
	@echo "  make stop-example  - 停止示例项目（如果正在运行）"
	@echo "  make show-version  - 显示当前框架版本号"
	@echo "  make update-major  - 自增主版本号 (x)"
	@echo "  make update-minor  - 自增次版本号 (y)"
	@echo "  make update-patch  - 自增修订版本号 (z)"
	@echo "  make build-tf      - 构建tf工具二进制文件"
	@echo "  make git-add       - 添加所有修改的文件到Git暂存区"
	@echo "  make git-commit    - 使用当前版本号作为提交信息"
	@echo "  make git-all       - 组合添加并提交操作"
	@echo "  make help          - 显示帮助信息"

# 显示当前版本号
show-version:
	@echo "当前框架版本号:"
	@grep -o 'version = "[0-9\\.]*"' tea.go | grep -o '[0-9\\.]*'

# 运行示例项目
run-example: stop-example
	@echo "正在启动示例项目..."
	@cd examples && go run manifest/cmd/server/main.go

# 停止示例项目
stop-example:
	@echo "正在停止示例项目..."
	@ps aux | grep "tea/examples/manifest/cmd/server/main.go" | grep -v grep | awk '{print $2}' | xargs -I {} kill {} 2>/dev/null || echo "未找到运行中的示例项目进程"

# 自增主版本号 (x)
update-major:
	@echo "正在自增主版本号..."
	@CURRENT_VERSION=$$(grep -o 'version = "[0-9\\.]*"' tea.go | grep -o '[0-9\\.]*'); \
	MAJOR=$$(echo $$CURRENT_VERSION | cut -d. -f1); \
	MAJOR=$$((MAJOR + 1)); \
	NEW_VERSION="$$MAJOR.0.0"; \
	sed -i '' "s/version = \"[0-9\\.]*\"/version = \"$$NEW_VERSION\"/g" tea.go; \
	echo "主版本号更新成功! 新的版本号: $$NEW_VERSION"

# 自增次版本号 (y)
update-minor:
	@echo "正在自增次版本号..."
	@CURRENT_VERSION=$$(grep -o 'version = "[0-9\\.]*"' tea.go | grep -o '[0-9\\.]*'); \
	MAJOR=$$(echo $$CURRENT_VERSION | cut -d. -f1); \
	MINOR=$$(echo $$CURRENT_VERSION | cut -d. -f2); \
	MINOR=$$((MINOR + 1)); \
	NEW_VERSION="$$MAJOR.$$MINOR.0"; \
	sed -i '' "s/version = \"[0-9\\.]*\"/version = \"$$NEW_VERSION\"/g" tea.go; \
	echo "次版本号更新成功! 新的版本号: $$NEW_VERSION"

# 自增修订版本号 (z)
update-patch:
	@echo "正在自增修订版本号..."
	@CURRENT_VERSION=$$(grep -o 'version = "[0-9\\.]*"' tea.go | grep -o '[0-9\\.]*'); \
	MAJOR=$$(echo $$CURRENT_VERSION | cut -d. -f1); \
	MINOR=$$(echo $$CURRENT_VERSION | cut -d. -f2); \
	PATCH=$$(echo $$CURRENT_VERSION | cut -d. -f3); \
	PATCH=$$((PATCH + 1)); \
	NEW_VERSION="$$MAJOR.$$MINOR.$$PATCH"; \
	sed -i '' "s/version = \"[0-9\\.]*\"/version = \"$$NEW_VERSION\"/g" tea.go; \
	echo "修订版本号更新成功! 新的版本号: $$NEW_VERSION"

# 构建tf工具二进制文件
build-tf:
	@echo "正在构建tf工具二进制文件..."
	@cd cli/cmd/tf && go build -o ../../../bin/tf main.go
	@echo "tf工具构建成功! 二进制文件位于: $(PWD)/bin/tf"

# 添加所有修改的文件到Git暂存区
git-add:
	@echo "正在添加所有修改的文件到Git暂存区..."
	@git add .
	@echo "添加完成！"

# 使用当前版本号作为提交信息
git-commit:
	@echo "正在提交更改..."
	@CURRENT_VERSION=$$(grep -o 'version = "[0-9\\.]*"' tea.go | grep -o '[0-9\\.]*'); \
	if git diff --staged --quiet; then \
	  echo "没有任何更改需要提交！"; \
	else \
	  git commit -m "$$CURRENT_VERSION"; \
	  echo "提交完成！"; \
	fi

# 组合添加并提交操作
git-all:
	@echo "执行添加并提交操作..."
	@$(MAKE) git-add
	@$(MAKE) git-commit
	@echo "所有操作完成！"