#!/bin/bash

make

./main.exe << EOF
3
2 2 4
3 2 4 6
4 2 1 999 5
EOF

make clean
