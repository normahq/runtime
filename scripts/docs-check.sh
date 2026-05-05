#!/bin/sh

set -eu

module="$(go list -m -f '{{.Path}}')"

go list -f '{{.ImportPath}}|{{.Doc}}' ./... | while IFS='|' read -r import doc; do
  if [ -z "$doc" ]; then
    echo "missing package doc for $import" >&2
    exit 1
  fi

  local_pkg="./${import#"$module"/}"
  if [ "$import" = "$module" ]; then
    local_pkg="./"
  fi

  go doc "$local_pkg" > /dev/null 2>/dev/null
done
