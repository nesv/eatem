#!/usr/bin/env bash
set -ex
cleanup() {
    make clean
}
trap cleanup EXIT
make all
./terf -config="terf.example.conf" -v=1 -logtostderr $@