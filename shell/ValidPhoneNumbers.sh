#!/bin/bash

# 193. Valid Phone Numbers
# https://leetcode.com/problems/valid-phone-numbers/

echo -e "987-123-4567\n123 456 7890\n(123) 456-7890" | \
    grep -P '^(\d{3}-|\(\d{3}\) )\d{3}-\d{4}$'
