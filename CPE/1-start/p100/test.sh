#!/bin/bash

make

./main.exe << EOF
1 10
100 200
201 210
900 1000
EOF

make clean
