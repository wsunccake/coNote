#!/bin/bash

make

./main.exe << EOF
10 12
14 10
EOF

make clean
