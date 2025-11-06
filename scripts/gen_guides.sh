#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "$0")/.." && pwd)"
OUT_DIR="$ROOT_DIR/docs/guide"
OUT_FILE="$OUT_DIR/RELEASES.md"

mkdir -p "$OUT_DIR"
{
  echo "# 版本与发布（基于 git tag 自动生成）"
  echo
  echo "以下内容基于仓库中的 tag 生成："
  echo
  echo "| Tag | 日期 | 摘要 |"
  echo "| --- | ---- | ---- |"

  # 使用 commit 日期与提交摘要，兼容轻量和附注 tag
  for tag in $(git tag --list --sort=-creatordate); do
    commit=$(git rev-list -n 1 "$tag")
    date=$(git show -s --format=%as "$commit")
    subject=$(git show -s --format=%s "$commit" | sed 's/|/-/g')
    printf "| %s | %s | %s |\n" "$tag" "$date" "$subject"
  done
} > "$OUT_FILE"

echo "Generated: $OUT_FILE"


