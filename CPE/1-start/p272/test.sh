#!/bin/bash

make

./main.exe << EOF
"To be or not to be," quoth the Bard, "that
is the question".
The programming contestant replied: "I must disagree.
To \`C' or not to \`C', that is The Question!"
EOF

make clean
