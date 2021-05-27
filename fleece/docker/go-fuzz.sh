#! /usr/bin/env bash
set -e

# TODO: replace this shell script
# $1 - fuzzer package
# $2 - name of the fuzzer function (will be used for the `workdir` argument to `go-fuzz`).
# `--build` - (optional) builds fuzzer package before being run.
# Additional args following `--` are passed directly to `go-fuzz`.

if [[ $# -lt 2 ]]; then
  echo "usage: entrypoint.sh <fuzzer package path> <fuzzer func name> [-b|--build] [-- [go-fuzz arg[, ...]]]"
fi

snake_case() {
  echo $1 | sed 's,/,_,g'
}

pkg=$1
shift
name=$1
shift
# TODO: this hack breaks support for global --config flag
workdir=$(cat .fleece.yaml |grep fleece-dir|sed -E 's,.+:\s(.+),\1,')/workdirs/${name}
bin=$(snake_case "${pkg:2}")-fuzz.zip
if [[ pkg == "." || pkg == "./" ]]; then
  bin="./dot-fuzz.zip"
fi

bin_path=${workdir}/${bin}

has_built=false
build() {
  mkdir -p "${workdir}"
  go-fuzz-build -o "${bin_path}" "${pkg}"
  echo "build done"
}

while [[ $# -gt 0 ]]; do
  case $1 in
  -b | --build)
    echo "Building fuzz binary..."
    build
    has_built=true
    shift
    ;;
  --)
    shift
    rest_args=$@
    break
    ;;
  *)
    shift
    ;;
  esac
done

if [[ $has_built == "false" && ! -e $bin_path ]]; then
  echo "Fuzz binary not found; building..."
  build
fi

rest_args=$@

go-fuzz -bin="${bin_path}" -func="${name}" -workdir="${workdir}" $rest_args
