#!/bin/sh

bazel query 'deps(//shapes:shapes)' --output graph > deps.dot
dot -Tpng deps.dot -o deps.png
