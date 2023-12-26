#!/bin/bash

make

./main.exe << EOF
123 456
555 555
123 594
999 1
0 0
EOF

make clean
