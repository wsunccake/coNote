#!/bin/bash

make

./main.exe << EOF
pretty
women
walking
down
the
street
EOF

make clean
