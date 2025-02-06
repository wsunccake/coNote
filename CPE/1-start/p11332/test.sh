#!/bin/bash

make

./main.exe << EOF
2
11
47
1234567892
10
0
EOF

make clean
