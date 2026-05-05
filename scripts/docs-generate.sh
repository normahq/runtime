#!/bin/sh

set -eu

mkdir -p .cache/godoc
module="$(go list -m -f '{{.Path}}')"

go list ./... | while IFS= read -r import; do
  local_pkg="./${import#"$module"/}"
  if [ "$import" = "$module" ]; then
    local_pkg="./"
  fi

  out_name="$(printf '%s' "$local_pkg" | tr '/.' '__').txt"
  go doc -all "$local_pkg" > ".cache/godoc/$out_name" 2>/dev/null
done
