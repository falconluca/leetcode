#!/bin/bash

# 194. Transpose File
# https://leetcode.com/problems/transpose-file/

# 行转列
echo -e "name age\nalice 21\nryan 30" | \
awk '
{
    for (i = 1;i <= NF; i++) {
        if (NF == 1) {
            arr[i] = $i;
            continue;
        }
        
        arr[i] = arr[i] " " $i;
    }
}

END {
    for (i = 1; i <= NR; i++) {
        if (arr[i] != "") {
            print arr[i];
        }
    }
}
'
