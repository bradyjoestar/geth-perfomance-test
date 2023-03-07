#!/bin/bash
git clone https://github.com/bradyjoestar/optimism.git
cd optimism/l2geth
git checkout wb/performance-test

make geth
mv build/bin/geth /usr/bin/geth_op_linux