#! /usr/bin/env bash

# Based on https://github.com/kubernetes/kubernetes/issues/79384#issuecomment-521493597

set -o errexit
set -o pipefail
set -o nounset

readonly VERS=${1:-""}

readonly MODS=($(
    curl -sS -L "https://raw.githubusercontent.com/kubernetes/kubernetes/v${VERS}/go.mod" | \
        sed -n 's|.*k8s.io/\(.*\) => ./staging/src/k8s.io/.*|k8s.io/\1|p'
))

if [ ${#MODS[@]} -eq 0 ]; then
    echo no matching version for $VERS
    exit 1
fi

for m in "${MODS[@]}" ; do
    echo matching version for $m
    v=$(go mod download -json "${m}@kubernetes-${VERS}" | jq -r .Version)
    go mod edit -replace=${m}=${m}@${v}
done

go get "k8s.io/kubernetes@v${VERS}"
