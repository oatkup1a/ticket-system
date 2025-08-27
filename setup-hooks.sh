#!/usr/bin/env bash
set -euo pipefail

if [ ! -d ".git" ]; then
  echo "❌ Error: run this from the root of a git repository."
  exit 1
fi

HOOKS_DIR=".git/hooks"

chmod +x hooks/pre-commit hooks/pre-push

ln -sf ../../hooks/pre-commit "${HOOKS_DIR}/pre-commit"
ln -sf ../../hooks/pre-push   "${HOOKS_DIR}/pre-push"

echo "Git hooks linked:"
ls -l "${HOOKS_DIR}/pre-commit" "${HOOKS_DIR}/pre-push"

# ✅ Success message
echo "✅ Hooks setup completed successfully!"
