#!/usr/bin/env bash

required_files=(
  README.md
  VERSION
  CHANGELOG.md
  CONTRIBUTING.md
  LICENSE
  base_manifest.yaml
)

for file in "${required_files[@]}"; do
  [[ -f "$file" ]] || {
    printf 'Missing required file: %s\n' "$file" >&2
    exit 1
  }
done

printf 'Repository baseline is present.\n'
