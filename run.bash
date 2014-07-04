#!/usr/bin/env bash
set -ex
cleanup() {
    make clean
}
trap cleanup EXIT
make all
./eatem \
    -listen="127.0.0.1:5000" \
    -templates="./templates" \
    -static="./static" \
    -logtostderr
