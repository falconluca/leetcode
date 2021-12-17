#!/bin/bash

# 195. Tenth Line
# https://leetcode.com/problems/tenth-line/

# 提取指定行

# sed
echo Line" "{1..10} | \
    xargs -n 2 | \
    sed -n '10p'

# awk
echo Line" "{1..10} | \
    xargs -n 2 | \
    awk 'FNR==10{print $0}'
