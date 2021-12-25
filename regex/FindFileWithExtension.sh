#!/bin/bash

echo '
readme.md
document.pdf
pdf.txt
image.png
music.mp4
manual.pdf' | grep -P '^\w+\.pdf$'
