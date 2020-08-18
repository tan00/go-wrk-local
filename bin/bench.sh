#!/bin/bash

./go-wrk-local.linux  -alg AES-128-ECB -size 1024 -n 200000 -t 2
./go-wrk-local.linux  -alg AES-256-ECB -size 1024 -n 200000 -t 2

./go-wrk-local.linux  -alg SMS4-ECB -size 1280 -n 200000 -t 2
./go-wrk-local.linux  -alg SM3 -size 1024 -n 200000 -t 2
