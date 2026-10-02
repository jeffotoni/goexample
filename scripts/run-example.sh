#!/usr/bin/env bash

set -euo pipefail

if [ "$#" -lt 1 ]; then
	printf 'usage: %s <example-directory> [command]\n' "$0" >&2
	printf 'example: %s promise\n' "$0" >&2
	exit 2
fi

script_dir=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)
repo_root=$(CDPATH= cd -- "$script_dir/.." && pwd)
example_path=$1
shift

case "$example_path" in
	/*|../*|*/../*)
		printf 'example directory must be inside the repository\n' >&2
		exit 2
		;;
esac

source_dir=$(CDPATH= cd -- "$repo_root/$example_path" 2>/dev/null && pwd) || {
	printf 'example directory not found: %s\n' "$example_path" >&2
	exit 1
}

case "$source_dir" in
	"$repo_root"/*) ;;
	*)
		printf 'example directory must be inside the repository\n' >&2
		exit 2
		;;
esac

temp_dir=$(mktemp -d "${TMPDIR:-/tmp}/goexample.XXXXXX")
cleanup() {
	/bin/rm -rf "$temp_dir"
}
trap cleanup EXIT

cp -R "$source_dir"/. "$temp_dir"/
find "$temp_dir" -type f \( -name go.mod -o -name go.sum \) -delete

cd "$temp_dir"
go mod init "local.goexample/$(basename "$source_dir")" >/dev/null
go mod tidy

if [ "$#" -gt 0 ]; then
	if "$@"; then
		return_code=0
	else
		return_code=$?
	fi
else
	if go test ./...; then
		return_code=0
	else
		return_code=$?
	fi
fi

exit "$return_code"
