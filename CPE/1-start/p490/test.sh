#!/bin/bash

make

./main.exe << EOF
Rene Decartes once said,
"I think, therefore I am."
EOF

make clean
