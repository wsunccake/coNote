#!/bin/bash

make

./main.exe << EOF
3
This is a test.
Count me 1 2 3 4 5.
Wow!!!!  Is this question easy?
EOF

make clean
