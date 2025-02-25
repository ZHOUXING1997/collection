# 默认值
TAG_NAME ?=
TAG_MESSAGE ?= ""

.PHONY: add_tag
add_tag:
	@if [ -z "$(TAG_NAME)" ]; then \
		echo "错误: TAG_NAME为空，请提供有效的标签名称"; \
		exit 1; \
	fi
	@if [ -z "$(TAG_MESSAGE)" ]; then \
		TAG_MESSAGE=$(TAG_NAME)
	fi

	git tag $(TAG_NAME) -m $(TAG_MESSAGE)
	git push origin $(TAG_NAME)

.DEFAULT_GOAL := help

.PHONY: push_pkg
push_pkg:
	 go build -mod=mod -o github.com/ZHOUXING1997/collection
	 go mod tidy
	 go list -m github.com/ZHOUXING1997/collection@$(TAG_NAME)

.PHONY: help
help:
	@echo "可用命令列表:"
	@echo "  wire: 生成代码."
	@echo "请使用 'make wire' 来执行相应命令."
